package keeper

import (
	"fmt"
	"github.com/airchains-network/junction/x/trackgate/types"
	"regexp"
)

// ValidateVersion checks if the provided version string conforms to Semantic Versioning 2.0.0
func ValidateVersion(version string) bool {
	// Regex for Semantic Versioning (https://semver.org/)
	var semverRegex = regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(-((0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(\.(0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(\+([0-9a-zA-Z-]+(\.[0-9a-zA-Z-]+)*))?$`)
	return semverRegex.MatchString(version)
}

// BuildExtTrackSchemaPath constructs a path for storing ext-track-schema objects in a database.
func BuildExtTrackSchemaPath(extTrackStationID string) string {
	dbPath := fmt.Sprintf("%s/%s", types.ExtTrackSchemaStoreKey, extTrackStationID)

	return dbPath
}
