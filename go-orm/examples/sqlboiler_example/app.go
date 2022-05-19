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

	newTodo := &models.Todo{
		Title:       "buy a thing 1",
		Description: "buy a thing desc",
		Status:      models.StatusTodo,
		UserID:      null.Int64{Int64: newUser.ID, Valid: true},
	}
	err = newTodo.Insert(ctx, conn, boil.Infer())
	internal.LogFatal(err)
	internal.PrintJSONLog(newTodo)

	gotUser, err := models.Users(qm.Load("Todos")).One(ctx, conn)
	internal.LogFatal(err)
	internal.PrintJSONLog(gotUser)
	internal.PrintJSONLog(gotUser.R.Todos)
}
