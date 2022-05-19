package ent_example

import (
	"ariga.io/atlas/sql/migrate"
	"context"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq"
	"go-orm/ent_example/ent"
	"go-orm/ent_example/ent/todo"
	"go-orm/ent_example/ent/user"
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
	log.Println(newUser)

	newTodo, err := client.Todo.Create().
		SetUser(newUser).
		SetStatus(todo.StatusTodo).
		SetTitle("buy a thing 1").
		SetDescription("this is description").
		Save(ctx)
	internal.LogFatal(err)
	log.Println(newTodo)

	userTodos, err := newUser.QueryTodos().All(ctx)
	internal.LogFatal(err)
	log.Println(userTodos)

	//Eager loading
	gotUser, err := client.User.Query().
		Where(user.ID(newUser.ID)).
		WithTodos().
		Only(ctx)
	internal.LogFatal(err)
	log.Println(gotUser)
	for _, t := range gotUser.Edges.Todos {
		log.Println(t)
	}
	log.Println(string(internal.PrettyJson(gotUser)))
}

func makeClient() *ent.Client {
	client, err := ent.Open(config.GetDatabaseURL())
	internal.LogFatal(err)

	return client
}
