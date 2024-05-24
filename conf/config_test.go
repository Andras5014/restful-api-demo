package conf

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestLoadConfigFormToml(t *testing.T) {
	err := LoadConfigFormToml("../etc/demo.toml")
	require.NoError(t, err)
	if err != nil {
		return
	}
	fmt.Println(config.App)
}

func TestLoadConfigFormEnv(t *testing.T) {
	os.Setenv("MYSQL_DATABASE", "unit_test")
	err := LoadConfigFormEnv()
	require.NoError(t, err)
	if err != nil {
		return
	}
	fmt.Println(GetConfig().MySQL.Database)
}

func TestMySQL_GetDB(t *testing.T) {
	err := LoadConfigFormToml("../etc/demo.toml")
	require.NoError(t, err)
	if err != nil {
		return
	}
	query, err := db.Query("select * from `test`")
	if err != nil {
		return
	}
	var tx test
	for query.Next() {
		err := query.Scan(&tx.id, &tx.name)
		if err != nil {
			return
		}
		fmt.Println(tx)
	}
}

type test struct {
	id   int
	name string
}
