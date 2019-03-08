package gographi

import (
	"context"
	"fmt"
	"math/rand"
)

// Resolver struct
type Resolver struct {
	todos []Todo
}

// Mutation comment
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

// Query comment
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

// Todo comment
func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*Todo, error) {
	todo := &Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, *todo)
	return todo, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]Todo, error) {
	return r.todos, nil
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *Todo) (*User, error) {
	return &User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}
