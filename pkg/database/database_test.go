package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanConnectToPostgres(t *testing.T) {

	t.Run("test_can_connect_to_postgres_and_can_ping", func(t *testing.T) {
		db, err := GetDb(false, "")
		assert.Nil(t, err)
		assert.NotNil(t, db)

		sqlDb, err := db.DB()
		assert.Nil(t, err)
		assert.NotNil(t, sqlDb)
		assert.NotNil(t, sqlDb.Ping())
	})
}
