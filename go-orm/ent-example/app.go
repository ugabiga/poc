package ent_example

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"go-orm/ent-example/ent"
	"log"
)

func Run() {
	ctx := context.Background()
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	panicErr(err)
	defer func(client *ent.Client) {
		_ = client.Close()
	}(client)

	err = client.Schema.Create(ctx)
	if err != nil {
		panicErr(err)
	}

	user, err := client.User.
		Create().
		SetFirstName("john").
		SetLastName("park").
		Save(ctx)
	panicErr(err)
	log.Println(user)

}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
