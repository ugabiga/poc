// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: todo.sql

package db

import (
	"context"
	"database/sql"
)

const createTodo = `-- name: CreateTodo :one
insert into todos (title, description, status, user_id)
values ($1, $2, $3, $4)
returning id, title, description, status, updated_at, created_at, user_id
`

type CreateTodoParams struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Status      Status        `json:"status"`
	UserID      sql.NullInt64 `json:"user_id"`
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, createTodo,
		arg.Title,
		arg.Description,
		arg.Status,
		arg.UserID,
	)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Status,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.UserID,
	)
	return i, err
}
