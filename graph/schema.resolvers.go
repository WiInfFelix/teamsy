package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
	"log"
	"strconv"
	"teamsy/graph/generated"
	"teamsy/graph/model"
	"teamsy/internal/pkg/jwt"
	"teamsy/internal/pkg/organisations"
	"teamsy/internal/pkg/users"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user users.User
	var err error

	user.Username = input.Username
	user.Email = input.Email
	user.Password = input.Password

	user.Birthday, err = users.ParseBirthdayDate(input.Birthday)

	_, err = user.Save()
	if err != nil {
		return "", err
	}

	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) CreateOrganisation(ctx context.Context, input model.NewOrganisation) (string, error) {
	orgDB := organisations.Organisation{
		Model:            gorm.Model{},
		OrganisationName: input.Name,
		Email:            input.Email,
	}

	res, err := orgDB.Save()
	if err != nil {
		return "", err
	}

	orgID := strconv.FormatInt(int64(res), 10)

	return orgID, err
}

func (r *mutationResolver) CreateDivision(ctx context.Context, input model.NewDivision) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTeam(ctx context.Context, input model.NewTeam) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateJoinRequest(ctx context.Context, input model.NewJoinRequest) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetOrganisations(ctx context.Context) ([]*model.Organisation, error) {
	var orgs []*model.Organisation
	var orgsDB []*organisations.Organisation

	orgsDB, err := organisations.GetAll()
	if err != nil {
		errStr := "There was en error fetching the orgs... %s"
		log.Fatalf(errStr, err)
		return orgs, gqlerror.Errorf(errStr, err.Error())
	}

	for _, org := range orgsDB {
		orgRes := model.Organisation{
			ID:               string(org.ID),
			OrganisationName: org.OrganisationName,
			Email:            org.Email,
			AdminMembers:     nil,
		}
		orgs = append(orgs, &orgRes)

	}

	return orgs, nil
}

func (r *queryResolver) GetOrganisationDivisions(ctx context.Context) ([]*model.Division, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetDivisionTeams(ctx context.Context) ([]*model.Team, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetTeamMembers(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
