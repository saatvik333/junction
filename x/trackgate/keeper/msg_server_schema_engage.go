package keeper

import (
	"context"
	"encoding/json"
	"strconv"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/airchains-network/junction/x/trackgate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SchemaEngage(goCtx context.Context, msg *types.MsgSchemaEngage) (*types.MsgSchemaEngageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	operator := msg.Operator
	extTrackStationId := msg.ExtTrackStationId
	schemaKey := msg.SchemaKey
	schemaObject := msg.SchemaObject
	stateRoot := msg.StateRoot
	podNumber := msg.PodNumber

	extTrackStationsDataDB := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ExtTrackStationsDataStoreKey))
	extTrackStationIdByte := []byte(extTrackStationId)
	stationDataByte := extTrackStationsDataDB.Get(extTrackStationIdByte)
	if stationDataByte == nil {
		return nil, status.Errorf(codes.NotFound, "ext track station %s not found", extTrackStationId)
	}

	var stationData types.ExtTrackStations
	k.cdc.MustUnmarshal(stationDataByte, &stationData)

	// check operator
	operators := stationData.Operators
	if !ContainsOperator(operators, operator) {
		return nil, status.Errorf(codes.PermissionDenied, "operator %s not found in ext track station %s", operator, extTrackStationId)
	}

	// check pod number
	if stationData.LatestPod+1 != podNumber {
		return nil, status.Errorf(codes.PermissionDenied, "pod number %d not match with ext track station %s", podNumber, extTrackStationId)
	}

	versionFinderDbPath := BuildExtTrackVersionFinderStoreKey(extTrackStationId)
	versionFinderStore := prefix.NewStore(storeAdapter, types.KeyPrefix(versionFinderDbPath))
	versionFinderKey := []byte(schemaKey)
	versionByte := versionFinderStore.Get(versionFinderKey)
	if versionByte == nil {
		return nil, status.Errorf(codes.NotFound, "schema %s not found in ext track station %s", schemaKey, extTrackStationId)
	}

	dbPath := BuildExtTrackSchemaPath(extTrackStationId)
	trackSchemaStore := prefix.NewStore(storeAdapter, types.KeyPrefix(dbPath))

	var schema types.ExtTrackSchema
	schemaDataByte := trackSchemaStore.Get(versionByte)

	if schemaDataByte == nil {
		return nil, status.Errorf(codes.NotFound, "schema %s not found in ext track station %s", schemaKey, extTrackStationId)
	}

	k.cdc.MustUnmarshal(schemaDataByte, &schema)

	var schemaDef types.SchemaDef
	if err := json.Unmarshal(schema.Schema, &schemaDef); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to unmarshal schema %s", schemaKey)
	}

	schemaObjectString := string(schemaObject)

	_, err := DynamicUnmarshal(schemaDef, schemaObjectString)
	if err != nil {
		return nil, err
	}

	newSchemaEngagement := types.ExtTrackSchemaEngagement{
		ExtTrackStationId: extTrackStationId,
		EngageBy:          operator,
		StateRoot:         stateRoot,
		PodNumber:         podNumber,
		StationName:       stationData.Name,
		SchemaKey:         schemaKey,
		SchemaObject:      schemaObject,
	}

	newSchemaEngagementDbPath := BuildExtTrackSchemaEngagementsStoreKey(extTrackStationId)
	newSchemaEngagementStore := prefix.NewStore(storeAdapter, types.KeyPrefix(newSchemaEngagementDbPath))
	podNumberString := strconv.FormatUint(podNumber, 10)
	newEngagementKey := []byte(podNumberString)

	check := newSchemaEngagementStore.Get(newEngagementKey)
	if check != nil {
		return nil, status.Errorf(codes.AlreadyExists, "schema engagement already exists for pod number %d", podNumber)
	}

	newEngagementValue := k.cdc.MustMarshal(&newSchemaEngagement)
	newSchemaEngagementStore.Set(newEngagementKey, newEngagementValue)
	// update codes
	updateStationDetails := types.ExtTrackStations{
		Operators:            stationData.Operators,
		LatestPod:            podNumber,
		LatestMerkleRootHash: stateRoot,
		Name:                 stationData.Name,
		Id:                   stationData.Id,
		StationType:          stationData.StationType,
		FheEnabled:           stationData.FheEnabled,
		SequencerDetails:     stationData.SequencerDetails,
		DaDetails:            stationData.DaDetails,
		ProverDetails:        stationData.ProverDetails,
	}

	byteStationId := []byte(stationData.Id)
	byteUpdateStationDetails := k.cdc.MustMarshal(&updateStationDetails)

	extTrackStationsDataDB.Set(byteStationId, byteUpdateStationDetails)

	return &types.MsgSchemaEngageResponse{
		Status: true,
	}, nil
}