package common

import (
	"context"
	"database/sql"
	"log"
	"modernc.org/sqlite"
	"starter/internal/ent"
)

func init() {
	sql.Register("sqlite3", sqliteDriver{Driver: &sqlite.Driver{}})
}

func NewEntClient(driverName, dataSourceName string) *ent.Client {
	client, err := ent.Open(driverName, dataSourceName)

	if err != nil {
		log.Fatalf("failed opening connection to: %v", err)
	}

	if err := client.Schema.Create(
		context.Background(),
	); err != nil {
		log.Fatalf("failed creating schema resource: %v", err)
	}

	return client
}
