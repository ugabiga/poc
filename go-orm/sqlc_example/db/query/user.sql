-- name: CreateUser :one
insert into users (first_name, last_name)
VALUES ($1, $2)
returning *
;

-- name: GetUserWithTodo :many
select *
from users as u
         left join todos t on u.id = t.user_id
where u.id = $1;