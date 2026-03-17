package reporter

import "strings"

// beamServiceName extracts the service name from a BEAM executable path.
// For paths like "/home/deploy/discord_presence/erts-13.2.2.4/bin/beam.smp",
// it returns "discord_presence" (the directory before "erts-*").
// Returns empty string if the path doesn't match the expected pattern.
func beamServiceName(execPath string) string {
	if execPath == "" {
		return ""
	}
	parts := strings.Split(execPath, "/")
	for i, part := range parts {
		if strings.HasPrefix(part, "erts-") && i > 0 {
			return parts[i-1]
		}
	}
	return ""
}
