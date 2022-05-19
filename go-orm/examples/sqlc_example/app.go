package sqlc_example

import (
	"context"
	"database/sql"
	"go-orm/config"
	db "go-orm/examples/sqlc_example/db/sqlc"
	"go-orm/internal"
	"log"
)

func Run() {
	ctx := context.Background()
	conn, err := sql.Open(config.GetDatabaseURL())
	internal.LogFatal(err)

	query := db.New(conn)

	createdUser, err := query.CreateUser(ctx, db.CreateUserParams{
		FirstName: "john",
		LastName:  "park",
	})
	internal.LogFatal(err)
	log.Println(createdUser)

	createdTodo, err := query.CreateTodo(ctx, db.CreateTodoParams{
		Title:       "buy a thing 1",
		Description: "buy a thing 1",
		Status:      db.StatusTodo,
		UserID:      sql.NullInt64{Int64: createdUser.ID, Valid: true},
	})
	internal.LogFatal(err)
	log.Println(createdTodo)

	gotUserWithTodo, err := query.GetUserWithTodo(ctx, createdUser.ID)
	internal.LogFatal(err)
	log.Println(string(internal.PrettyJson(gotUserWithTodo)))

}
