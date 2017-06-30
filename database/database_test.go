package database

import (
	"testing"
)

func TestMysql(t *testing.T) {
	DB := Connect()
	err := DB.Ping()
	if err != nil {
		t.Errorf("DB connection %d", "ping error")
	}
}
