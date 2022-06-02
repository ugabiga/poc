package ent_example

import (
	"ariga.io/atlas/sql/migrate"
	"context"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/bxcodec/faker/v3"
	_ "github.com/lib/pq"
	"go-orm/examples/ent_example/ent"
	"go-orm/examples/ent_example/ent/project"
	"go-orm/examples/ent_example/ent/task"
	"go-orm/examples/ent_example/ent/user"
	"go-orm/internal"
	"math/rand"
	"time"

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

	var v []struct {
		UserTasks int `json:"user_tasks"`
		Count     int `json:"count"`
	}
	err = client.Task.Query().
		Where(
			task.HasUserWith(
				user.IDGTE(5),
			),
		).
		Order(func(s *sql.Selector) {
			s.OrderBy(sql.Desc("count"))
		}).
		GroupBy(task.UserColumn).
		Aggregate(ent.Count()).
		Scan(ctx, &v)
	internal.LogFatal(err)
	internal.PrintJSONLog(v)
}

func Seed(userCount int, taskCount int) {
	ctx := context.Background()
	client := makeClient()
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(client)

	users := seedUsers(ctx, client, userCount)
	tasks := seedTasks(ctx, client, users, taskCount)
	internal.PrintJSONLog(users)
	internal.PrintJSONLog(tasks)
}

func seedUsers(ctx context.Context, client *ent.Client, count int) []*ent.User {
	var bulkUsers []*ent.UserCreate
	for i := 0; i < count; i++ {
		bulkUsers = append(bulkUsers,
			client.User.Create().
				SetFirstName(faker.FirstName()).
				SetLastName(faker.LastName()).
				SetBirthday(time.Now()),
		)
	}
	users, err := client.User.CreateBulk(bulkUsers...).Save(ctx)
	internal.LogFatal(err)
	return users
}

func seedTasks(ctx context.Context, client *ent.Client, users []*ent.User, count int) []*ent.Task {
	rand.Seed(time.Now().UnixNano())

	var bulkTasks []*ent.TaskCreate
	for i := 0; i < count; i++ {
		bulkTasks = append(bulkTasks,
			client.Task.Create().
				SetTitle(faker.Word()).
				SetNote(faker.Sentence()).
				SetStatus(task.StatusTodo).
				SetUser(users[rand.Intn(len(users))]),
		)
	}
	tasks, err := client.Task.CreateBulk(bulkTasks...).Save(ctx)
	internal.LogFatal(err)
	return tasks
}

func makeClient() *ent.Client {
	client, err := ent.Open(config.GetDBInfo())
	internal.LogFatal(err)

	return client
}
