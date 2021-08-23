package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"mse-gql/graph/api"
	"mse-gql/graph/generated"
	"mse-gql/graph/model"
)

func (r *mutationResolver) CreateRecipe(ctx context.Context, input model.NewRecipe) (*model.InsertResult, error) {
	result, err := api.InsertRecipe(input)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *mutationResolver) UpdateRecipe(ctx context.Context, input model.UpdateRecipe) (*model.UpdateDeleteResult, error) {
	result, err := api.UpdateRecipe(input)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *mutationResolver) DeleteRecipe(ctx context.Context, input int) (*model.UpdateDeleteResult, error) {
	result, err := api.DeleteRecipe(input)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *queryResolver) Recipe(ctx context.Context, input int) (*model.Recipe, error) {
	recipe, err := api.FetchRecipe(input)

	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (r *queryResolver) Search(ctx context.Context, input model.SearchParam) ([]*model.Recipe, error) {
	recipes, err := api.SearchRecipe(input.Keyword, input.IsHalal, input.IsVegetarian, input.Page, input.Take)

	if err != nil {
		return nil, err
	}

	return recipes, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
