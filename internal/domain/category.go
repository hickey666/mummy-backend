package domain

import (
	"mummy-storage/internal/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Category struct {
	ID          int    `gorm:"primaryKey;column:id" json:"-"`         // Primary Key
	Name        string `gorm:"column:name" json:"name"`               // 分类名称
	Description string `gorm:"column:description" json:"description"` // 描述
}

type CategoryDomain struct {
	dbEngine *gorm.DB
}

func NewCategoryDomain(dbEngine *gorm.DB) *CategoryDomain {
	return &CategoryDomain{dbEngine: dbEngine}
}

func (d *CategoryDomain) getDb() *gorm.DB {
	return d.dbEngine.Table(new(model.Category).TableName())
}

// CreateCategory 创建分类
func (d *CategoryDomain) CreateCategory(category *Category) error {
	err := d.getDb().Create(category).Error
	return errors.WithMessage(err, "CreateCategory failed")
}

// UpdateCategory 更新分类
func (d *CategoryDomain) UpdateCategory(category *Category) error {
	err := d.getDb().Where("id = ?", category.ID).Updates(category).Error
	return errors.WithMessage(err, "UpdateCategory failed")
}

// DeleteCategory 删除分类
func (d *CategoryDomain) DeleteCategory(id int) error {
	err := d.getDb().Where("id = ?", id).Update("is_deleted", 1).Error
	return errors.WithMessage(err, "DeleteCategory failed")
}

// GetCategoryByID 获取分类
func (d *CategoryDomain) GetCategoryByID(id int) (*Category, error) {
	var category Category
	err := d.getDb().Where("id = ?", id).First(&category).Error
	if err != nil {
		return nil, errors.WithMessage(err, "GetCategoryByID failed")
	}
	return &category, nil
}

// ListCategory 获取分类列表
func (d *CategoryDomain) ListCategory() ([]Category, error) {
	var categories []Category
	err := d.getDb().Where("is_deleted = 0").Find(&categories).Error
	if err != nil {
		return nil, errors.WithMessage(err, "ListCategory failed")
	}
	return categories, nil
}
