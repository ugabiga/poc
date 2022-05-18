package ent_example

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"go-orm/ent-example/ent"
	"go-orm/ent-example/ent/todo"
	"log"
)

func Run() {
	ctx := context.Background()
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	panicErr(err)
	defer func(client *ent.Client) {
		_ = client.Close()
	}(client)

	err = client.Schema.Create(ctx)
	if err != nil {
		panicErr(err)
	}

	newUser, err := client.User.
		Create().
		SetFirstName("john").
		SetLastName("park").
		Save(ctx)
	panicErr(err)
	log.Println(newUser)

	newTodo, err := client.Todo.Create().
		SetUser(newUser).
		SetStatus(todo.StatusTodo).
		SetTitle("buy a thing 1").
		SetDescription("this is description").
		Save(ctx)
	panicErr(err)
	log.Println(newTodo)

	userTodos, err := newUser.QueryTodos().All(ctx)
	panicErr(err)
	log.Println(userTodos)
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
