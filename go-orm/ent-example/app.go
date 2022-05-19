package ent_example

import (
	"ariga.io/atlas/sql/migrate"
	"context"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq"

	"go-orm/config"
	"go-orm/ent-example/ent/todo"
	"log"

	"go-orm/ent-example/ent"
)

func MakeClient() *ent.Client {
	client, err := ent.Open(
		"postgres",
		config.GetDBString(),
	)
	panicErr(err)

	return client
}

func GenerateMigration() {
	ctx := context.Background()
	client := MakeClient()
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(client)

	dir, err := migrate.NewLocalDir("ent-example/migrations")
	panicErr(err)

	name := "init"
	err = client.Schema.NamedDiff(ctx, name, schema.WithDir(dir))
	panicErr(err)
}

func Run() {
	ctx := context.Background()
	client := MakeClient()

	newUser, err := client.User.
		Create().
		SetFirstName("john").
		SetLastName("park").
		Save(ctx)
	panicErr(err)
	log.Println(newUser)

	newTodo, err := client.Todo.Create().
		SetUser(newUser).
		SetStatus(todo.StatusTodo).
		SetTitle("buy a thing 1").
		SetDescription("this is description").
		Save(ctx)
	panicErr(err)
	log.Println(newTodo)

	userTodos, err := newUser.QueryTodos().All(ctx)
	panicErr(err)
	log.Println(userTodos)
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
