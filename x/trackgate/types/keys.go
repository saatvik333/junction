package types

const (
	// ModuleName defines the module name
	ModuleName = "trackgate"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_trackgate"

	ExtTrackStationsDataStoreKey = "ext_track_stations_data"
)

var (
	ParamsKey = []byte("p_trackgate")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}