package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestEnvKeyReplacerForNestedKeys(t *testing.T) {
	// Reset viper state for the test
	viper.Reset()

	configureViper()
	setDefaults(false)

	// Verify the default data directory
	if got := viper.GetString("data.directory"); got != defaultDataDirectory {
		t.Errorf("expected default data.directory=%q, got %q", defaultDataDirectory, got)
	}

	// Set OPENCODE_DATA_DIRECTORY env var to override the nested config key
	localDir := "/tmp/opencode-local-test"
	os.Setenv("OPENCODE_DATA_DIRECTORY", localDir)
	defer os.Unsetenv("OPENCODE_DATA_DIRECTORY")

	got := viper.GetString("data.directory")
	if got != localDir {
		t.Errorf("expected data.directory=%q from env var OPENCODE_DATA_DIRECTORY, got %q", localDir, got)
	}
}
