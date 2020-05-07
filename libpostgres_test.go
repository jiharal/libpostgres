package libpostgres

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLibpostgres(t *testing.T) {
	opt := Options{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "",
		DBName:   "hello_db",
		SSLMode:  "disable",
	}
	db, err := ConnectWithAMP(opt)
	if err != nil {
		require.NoError(t, err)
	}
	defer db.Close()
	var name string
	err = db.QueryRow("select first_name from customer limit 1").Scan(&name)
	require.NoError(t, err)
	t.Log(name)

	db1, err := Connect(opt)
	if err != nil {
		require.NoError(t, err)
	}
	defer db1.Close()
	var name1 string
	err = db1.QueryRow("select first_name from customer limit 1").Scan(&name1)
	require.NoError(t, err)
	t.Log(name1)

	errM := Migrate(db, "file://./migrations")
	require.NoError(t, errM)

}
