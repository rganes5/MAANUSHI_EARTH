package interfaces

import (
	"context"

	domain "github.com/rganes5/maanushi_earth_e-commerce/pkg/domain"
)

type AdminUseCase interface {
	// FindAll(ctx context.Context) ([]domain.Users, error)
	FindByEmail(ctx context.Context, Email string) (domain.Admin, error)
	SignUpAdmin(ctx context.Context, admin domain.Admin) error
	// ListUsers(ctx context.Context) ([]utils.ResponseUsers, error)
	// FindByID(ctx context.Context, id uint) (domain.Users, error)
	// Save(ctx context.Context, user domain.Users) (domain.Users, error)
	// Delete(ctx context.Context, user domain.Users) error
}
