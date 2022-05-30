-- name: CreateTodo :one
insert into todos (title, description, status, user_id)
values ($1, $2, $3, $4)
returning *;