package reporter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBeamServiceName(t *testing.T) {
	tests := []struct {
		name     string
		execPath string
		expected string
	}{
		{
			name:     "standard Elixir release path",
			execPath: "/home/deploy/discord_presence/erts-13.2.2.4/bin/beam.smp",
			expected: "discord-presence",
		},
		{
			name:     "discord prefix with multiple underscores",
			execPath: "/home/deploy/discord_chat_service/erts-13.2.2.4/bin/beam.smp",
			expected: "discord-chat-service",
		},
		{
			name:     "OTP release path",
			execPath: "/opt/my_app/erts-14.0/bin/beam.smp",
			expected: "my_app",
		},
		{
			name:     "nested release path",
			execPath: "/srv/releases/production/my_service/erts-15.1.2/bin/beam.smp",
			expected: "my_service",
		},
		{
			name:     "no erts directory",
			execPath: "/usr/local/bin/beam.smp",
			expected: "",
		},
		{
			name:     "not a BEAM process",
			execPath: "/usr/bin/python3",
			expected: "",
		},
		{
			name:     "empty path",
			execPath: "",
			expected: "",
		},
		{
			name:     "beam.smp but erts at root",
			execPath: "/erts-13.0/bin/beam.smp",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := beamServiceName(tt.execPath)
			assert.Equal(t, tt.expected, result)
		})
	}
}
