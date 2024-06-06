package domain

import (
	"mummy-storage/internal/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ProductCategoryDomain struct {
	dbEngine *gorm.DB
}

func NewProductCategoryDomain(dbEngine *gorm.DB) *ProductCategoryDomain {
	return &ProductCategoryDomain{dbEngine: dbEngine}
}

func (d *ProductCategoryDomain) getDb() *gorm.DB {
	return d.dbEngine.Table(new(model.ProductCategory).TableName())
}

// CreateProductCategory 创建产品分类
func (d *ProductCategoryDomain) CreateProductCategory(productCategory *model.ProductCategory) error {
	err := d.getDb().Create(productCategory).Error
	return errors.WithMessage(err, "CreateProductCategory failed")
}

// UpdateProductCategory 更新产品分类
func (d *ProductCategoryDomain) UpdateProductCategory(productCategory *model.ProductCategory) error {
	err := d.getDb().Where("id = ?", productCategory.ID).Updates(productCategory).Error
	return errors.WithMessage(err, "UpdateProductCategory failed")
}

// DeleteProductCategory 删除产品分类
func (d *ProductCategoryDomain) DeleteProductCategory(id int) error {
	err := d.getDb().Where("id = ?", id).Update("is_deleted", 1).Error
	return errors.WithMessage(err, "DeleteProductCategory failed")
}

// GetProductCategoryByID 获取产品分类
func (d *ProductCategoryDomain) GetProductCategoryByID(id int) (*model.ProductCategory, error) {
	var productCategory model.ProductCategory
	err := d.getDb().Where("id = ?", id).First(&productCategory).Error
	if err != nil {
		return nil, errors.WithMessage(err, "GetProductCategoryByID failed")
	}
	return &productCategory, nil
}

// ListProductCategory 获取产品分类列表
func (d *ProductCategoryDomain) ListProductCategory() ([]model.ProductCategory, error) {
	var productCategories []model.ProductCategory
	err := d.getDb().Scopes(NotDeleted).Find(&productCategories).Error
	if err != nil {
		return nil, errors.WithMessage(err, "ListProductCategory failed")
	}
	return productCategories, nil
}
