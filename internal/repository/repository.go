package repository

import (
	"context"

	"github.com/adhyayana108/movies-crud-api/internal/domain"
)

type MovieRespository interface {

	// GetAll

	GetAll(ctx context.Context) ([]domain.Movie, error)

	//GetByID

	GetByID(ctx context.Context, id string) (domain.Movie, error)

	// Create

	Create(ctx context.Context, movie domain.Movie) (domain.Movie, error)

	// Update

	Update(ctx context.Context, id string, movie domain.Movie) (domain.Movie, error)

	// Delete

	Delete(ctx context.Context, id string) error
}
