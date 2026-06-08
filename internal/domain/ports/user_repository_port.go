package ports

import (
	"context"
	"local-notice-hex-go/internal/domain/model"
)

type UserRepositoryPort interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	FindByID(ctx context.Context, id int64) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	FindAll(ctx context.Context) ([]*model.User, error)
	Update(ctx context.Context, id int64, user *model.User) (*model.User, error)
	Delete(ctx context.Context, id int64) error
}
