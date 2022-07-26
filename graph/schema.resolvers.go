package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-uuid"
	"github.com/leofigy/events/graph/generated"
	"github.com/leofigy/events/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateEvent is the resolver for the createEvent field.
func (r *mutationResolver) CreateEvent(ctx context.Context, input *model.NewEvent) (*model.Event, error) {

	txn := r.memDb.Txn(
		true,
	)

	id, err := uuid.GenerateUUID()
	if err != nil {
		return nil, err
	}

	evt := &model.Event{
		ID:        id,
		CreatedAt: time.Now().String(),
		Source:    input.Source,
		Category:  input.Category,
		Message:   input.Message,
		Severity:  input.Severity,
	}

	if err = txn.Insert("events", evt); err != nil {
		return nil, err
	}

	txn.Commit()

	return evt, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Events is the resolver for the events field.
func (r *queryResolver) Events(ctx context.Context) ([]*model.Event, error) {
	// dummy placeholder returns all events no filter

	results := []*model.Event{}

	txn := r.memDb.Txn(false)
	it, err := txn.Get("events", "id")
	if err != nil {
		return nil, err
	}

	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*model.Event)
		results = append(results, p)
	}

	return results, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
