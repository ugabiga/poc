package common

import (
	"entgo.io/ent/dialect"
	"starter/internal/ent"
)

type Application struct {
	entClient *ent.Client
}

func NewApplication() *Application {
	entClient := NewEntClient(DBInfo())

	return &Application{
		entClient: entClient,
	}
}

func DBInfo() (string, string) {
	return dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1"
}
