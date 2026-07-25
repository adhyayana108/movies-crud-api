package service

import (
	"context"

	"github.com/adhyayana108/movies-crud-api/internal/domain"

)

type MovieService interface {

	GetAll(ctx context.Context) ([]domain.Movie, error)

	Create(ctx context.Context , movie domain.Movie)(domain.Movie, error)

	Update(ctx context.Context , id string , movie domain.Movie) (domain.Movie, error)

	Delete(ctx context.Context , id string) error

}