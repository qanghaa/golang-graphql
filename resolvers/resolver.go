package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-pg/pg"
	"github.com/qanghaa/golang-graphql/graph/generated"
	"github.com/qanghaa/golang-graphql/model"
)

type Resolver struct {
	DB *pg.DB
}

func (r *mutationResolver) CreateBook(ctx context.Context, book model.NewBook) (*model.Book, error) {
	rating := 0
	newBook := model.Book{
		Name:        book.Name,
		Description: book.Description,
		URL:         book.URL,
		ReleasedAt:  book.ReleasedAt,
		Rating:      &rating,
	}
	_, err := r.DB.Model(&newBook).Insert()
	if err != nil {
		return nil, fmt.Errorf("error inserting new book: %v", err)
	}
	return &newBook, nil
}

func (r *mutationResolver) UpdateBook(ctx context.Context, id string, changes model.UpdateBook) (*model.Book, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id string) (*model.Book, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateUser(ctx context.Context, user model.NewUser) (*model.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, changes model.UpdateUser) (*model.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string, isActive bool) (*model.User, error) {
	user := &model.User{ID: id}
	if err := r.DB.Select(&user); err != nil {
		return nil, errors.New("user does not exist")
	}
	unActive := false
	user.IsActive = &unActive
	return user, nil
}

func (r *mutationResolver) CreateReview(ctx context.Context, review model.NewReview) (*model.Review, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateReview(ctx context.Context, id string, changes model.UpdateReview) (*model.Review, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteReview(ctx context.Context, id string) (*model.Review, error) {
	panic("not implemented")
}

func (r *queryResolver) Book(ctx context.Context, id int) (*model.Book, error) {
	panic("not implemented")
}

func (r *queryResolver) Books(ctx context.Context, limit *int, offset *int) ([]*model.Book, error) {
	panic("not implemented")
}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	panic("not implemented")
}

func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) ([]*model.User, error) {
	panic("not implemented")
}

func (r *queryResolver) Review(ctx context.Context, id int) (*model.Review, error) {
	panic("not implemented")
}

func (r *queryResolver) Reviews(ctx context.Context, limit *int, offset *int) ([]*model.Review, error) {
	panic("not implemented")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
