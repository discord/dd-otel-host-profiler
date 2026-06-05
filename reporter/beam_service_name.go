package reporter

import "strings"

// beamServiceName extracts the service name from a BEAM executable path.
// For paths like "/home/deploy/discord_presence/erts-13.2.2.4/bin/beam.smp",
// it returns "discord-presence". Names with a "discord_" prefix have that
// prefix rewritten to "discord-" and any remaining underscores converted to
// dashes; other names are returned unchanged.
// Returns empty string if the path doesn't match the expected pattern.
func beamServiceName(execPath string) string {
	if execPath == "" {
		return ""
	}
	parts := strings.Split(execPath, "/")
	for i, part := range parts {
		if strings.HasPrefix(part, "erts-") && i > 0 {
			name := parts[i-1]
			if rest, ok := strings.CutPrefix(name, "discord_"); ok {
				return "discord-" + strings.ReplaceAll(rest, "_", "-")
			}
			return name
		}
	}
	return ""
}
