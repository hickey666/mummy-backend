package categoryctrl

import (
	"context"
	"mummy-storage/internal/domain"
	"mummy-storage/internal/environment"

	"github.com/zeromicro/go-zero/core/logc"
)

type CategoryCtrl struct {
	env *environment.ServiceContext
}

func NewCategoryCtrl(env *environment.ServiceContext) *CategoryCtrl {
	return &CategoryCtrl{
		env: env,
	}
}

func (c *CategoryCtrl) Create(ctx context.Context, name, desc string) error {
	params := &domain.Category{
		Name:        name,
		Description: desc,
	}
	err := c.env.Domain.CategoryDomain.CreateCategory(params)
	if err != nil {
		logc.Errorf(ctx, "CreateCategory failed: %v", err)
		return err
	}
	return nil
}
