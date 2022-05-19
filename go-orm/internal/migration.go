package internal

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go-orm/config"
	"log"
)

func makeMigrate() *migrate.Migrate {
	db, err := sql.Open("postgres", config.GetDBString())
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://ent-example/migrations",
		"postgres", driver)

	return m
}

func UpMigration() {
	m := makeMigrate()
	err := m.Up()
	PanicErr(err)
}

func DownMigration() {
	m := makeMigrate()
	err := m.Down()
	PanicErr(err)
}
