package ent_example

import (
	"ariga.io/atlas/sql/migrate"
	"context"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq"
	"go-orm/examples/ent_example/ent"
	"go-orm/examples/ent_example/ent/project"
	"go-orm/examples/ent_example/ent/task"
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

	dir, err := migrate.NewLocalDir("examples/ent_example/migrations")
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

	newTask, err := client.Task.Create().
		SetTitle("buy a ting 1").
		SetNote("").
		SetStatus(task.StatusTodo).
		SetUser(newUser).
		Save(ctx)
	internal.LogFatal(err)
	internal.PrintJSONLog(newTask)

	gotUser, err := client.User.Query().
		Where(user.IDEQ(newUser.ID)).
		Only(ctx)
	internal.LogFatal(err)
	internal.PrintJSONLog(gotUser)

	newProject, err := client.Project.Create().
		SetTitle("project 1").
		SetDescription("project 1").
		SetStatus(project.StatusTodo).
		AddTasks(newTask).
		Save(ctx)
	internal.LogFatal(err)
	internal.PrintJSONLog(newProject)

	gotProject, err := client.Project.Query().
		Where(project.ID(newProject.ID)).
		WithTasks(func(query *ent.TaskQuery) {
			query.WithUser()
		}).
		Only(ctx)
	internal.LogFatal(err)
	internal.PrintJSONLog(gotProject)
}

func makeClient() *ent.Client {
	client, err := ent.Open(config.GetDBInfo())
	internal.LogFatal(err)

	return client
}
