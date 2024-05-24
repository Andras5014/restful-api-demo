package conf

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestLoadConfigFormToml(t *testing.T) {
	os.Setenv("MYSQL_DATABASE", "unit_test")
	config, err := LoadConfigFormToml("../etc/demo.toml")
	require.NoError(t, err)
	if err != nil {
		return
	}
	fmt.Println(config.App)
}

func TestLoadConfigFormEnv(t *testing.T) {
	os.Setenv("MYSQL_DATABASE", "unit_test")
	config, err := LoadConfigFormEnv()
	require.NoError(t, err)
	if err != nil {
		return
	}
	fmt.Println(config.MySQL.Database)
}
