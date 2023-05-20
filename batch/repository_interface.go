package batch

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Migrate(ctx context.Context) error
	Create(ctx context.Context, batch Batch) (*Batch, error)
	List(ctx context.Context) ([]Batch, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Batch, error)
	GetByName(ctx context.Context, name string) (*Batch, error)
	Update(ctx context.Context, batch Batch) (*Batch, error)
	DeleteByID(ctx context.Context, id uuid.UUID) error
}
