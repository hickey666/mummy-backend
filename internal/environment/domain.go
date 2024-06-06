package environment

import (
	"mummy-storage/internal/domain"

	"gorm.io/gorm"
)

type DomainEnv struct {
	dbEngin        *gorm.DB
	CategoryDomain *domain.CategoryDomain
}

func NewDomainEnv(dbEngin *gorm.DB) *DomainEnv {
	return &DomainEnv{
		dbEngin:        dbEngin,
		CategoryDomain: domain.NewCategoryDomain(dbEngin),
	}
}
