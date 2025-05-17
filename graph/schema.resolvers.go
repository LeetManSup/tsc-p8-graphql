package graph

import (
	"context"
	"strconv"
	"tsc-p8-graphql/graph/generated"
	"tsc-p8-graphql/graph/model"
)

var idCounter = 1

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Profiles(ctx context.Context) ([]*model.Profile, error) {
	return r.profiles, nil
}

func (r *queryResolver) Profile(ctx context.Context, id string) (*model.Profile, error) {
	for _, p := range r.profiles {
		if p.ID == id {
			return p, nil
		}
	}
	return nil, nil
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateProfile(ctx context.Context, name string, age int) (*model.Profile, error) {
	profile := &model.Profile{
		ID:   strconv.Itoa(idCounter),
		Name: name,
		Age:  age,
	}
	idCounter++
	r.profiles = append(r.profiles, profile)
	return profile, nil
}

func (r *mutationResolver) UpdateProfile(ctx context.Context, id string, name *string, age *int) (*model.Profile, error) {
	for _, p := range r.profiles {
		if p.ID == id {
			if name != nil {
				p.Name = *name
			}
			if age != nil {
				p.Age = *age
			}
			return p, nil
		}
	}
	return nil, nil
}
