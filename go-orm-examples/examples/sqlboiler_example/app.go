package sqlboiler_example

import (
	"context"
	"database/sql"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go-orm/config"
	"go-orm/examples/sqlboiler_example/models"
	"go-orm/internal"
)

func Run() {
	ctx := context.Background()

	conn, err := sql.Open(config.GetDBInfo())
	internal.LogFatal(err)

	newUser := &models.User{
		FirstName: "john",
		LastName:  "park",
	}
	err = newUser.Insert(ctx, conn, boil.Infer())
	internal.LogFatal(err)
	internal.PrintJSONLog(newUser)

	newTask := &models.Task{
		Title:  "but a thing 1",
		Note:   "",
		Status: models.TaskStatusTodo,
		UserID: null.Int64{Int64: newUser.ID, Valid: true},
	}
	err = newTask.Insert(ctx, conn, boil.Infer())

	internal.LogFatal(err)
	internal.PrintJSONLog(newTask)

	gotUser, err := models.Users(qm.Load("Tasks")).One(ctx, conn)
	internal.LogFatal(err)
	internal.PrintJSONLog(gotUser)
	internal.PrintJSONLog(gotUser.R.Tasks)

	newProject := &models.Project{
		Title:       "project 1",
		Description: "project 1 description",
		Status:      models.ProjectStatusTodo,
	}
	err = newProject.Insert(ctx, conn, boil.Infer())
	internal.LogFatal(err)
	internal.PrintJSONLog(newProject)

	newProjectTask := &models.ProjectTask{
		ProjectID: newProject.ID,
		TaskID:    newTask.ID,
	}
	err = newProjectTask.Insert(ctx, conn, boil.Infer())
	internal.LogFatal(err)
	internal.PrintJSONLog(newProjectTask)

	gotProject, err := models.Projects(
		qm.Where("id = ?", newProject.ID),
		qm.Load("ProjectTasks"),
		qm.Load("ProjectTasks.Task"),
		qm.Load("ProjectTasks.Task.User"),
	).One(ctx, conn)
	internal.LogFatal(err)
	internal.PrintJSONLog(gotProject)
	internal.PrintJSONLog(gotProject.R.ProjectTasks)
	internal.PrintJSONLog(gotProject.R.ProjectTasks[0])
	internal.PrintJSONLog(gotProject.R.ProjectTasks[0].R.Task)
	internal.PrintJSONLog(gotProject.R.ProjectTasks[0].R.Task.R.User)
}
