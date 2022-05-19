package ent_example

import (
	"ariga.io/atlas/sql/migrate"
	"context"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq"
	"go-orm/examples/ent_example/ent"
	"go-orm/examples/ent_example/ent/todo"
	"go-orm/examples/ent_example/ent/user"
	"go-orm/internal"

	"go-orm/config"
	"log"
)

func GenerateMigration() {
	ctx := context.Background()
	client := makeClient()
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(client)

	dir, err := migrate.NewLocalDir("ent_example/migrations")
	internal.LogFatal(err)

	name := "init"
	err = client.Schema.NamedDiff(ctx, name, schema.WithDir(dir))
	internal.LogFatal(err)
}

func Run() {
	ctx := context.Background()
	client := makeClient()
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(client)

	newUser, err := client.User.
		Create().
		SetFirstName("john").
		SetLastName("park").
		Save(ctx)
	internal.LogFatal(err)
	internal.PrintJSONLog(newUser)

	newTodo, err := client.Todo.Create().
		SetUser(newUser).
		SetStatus(todo.StatusTodo).
		SetTitle("buy a thing 1").
		SetDescription("buy a thing desc").
		Save(ctx)
	internal.LogFatal(err)
	internal.PrintJSONLog(newTodo)

	userTodos, err := newUser.QueryTodos().All(ctx)
	internal.LogFatal(err)
	internal.PrintJSONLog(userTodos)

	//Eager loading
	gotUser, err := client.User.Query().
		Where(user.ID(newUser.ID)).
		WithTodos().
		Only(ctx)
	internal.LogFatal(err)
	internal.PrintJSONLog(gotUser)
	internal.PrintJSONLog(gotUser.Edges.Todos)
}

func makeClient() *ent.Client {
	client, err := ent.Open(config.GetDBInfo())
	internal.LogFatal(err)

	return client
}
