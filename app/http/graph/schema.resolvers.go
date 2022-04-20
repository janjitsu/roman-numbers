package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"studiosol.com.br/janjitsu/roman-numbers/app/http/graph/generated"
	"studiosol.com.br/janjitsu/roman-numbers/app/roman"
)

func (r *mutationResolver) Search(ctx context.Context, text string) (*roman.RomanNumber, error) {

	romanNumber := r.Resolver.Search.FindTopValueInString(text)
	return romanNumber, nil
}

func (r *queryResolver) Introspection(ctx context.Context) (string, error) {
	return "This query is only for introspection", nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
