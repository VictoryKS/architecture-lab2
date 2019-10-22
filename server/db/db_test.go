package db

import (
	"testing"
	"log"
)

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:     "balancers",
		User:       "root",
		Password:   "Afynfcvfujhbz",
		Host:       "localhost",
		DisableSSL: true,
	}
	if conn.ConnectionURL() != "postgres://root:Afynfcvfujhbz@localhost/balancers?sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}
