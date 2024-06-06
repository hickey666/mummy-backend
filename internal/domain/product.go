package domain

import (
	"mummy-storage/internal/consts"
	"mummy-storage/internal/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ProductDomain struct {
	dbEngine *gorm.DB
}

func NewProductDomain(dbEngine *gorm.DB) *ProductDomain {
	return &ProductDomain{dbEngine: dbEngine}
}

func (d *ProductDomain) getDb() *gorm.DB {
	return d.dbEngine.Table(new(model.Product).TableName())
}

// CreateProduct 创建产品
func (d *ProductDomain) CreateProduct(product *model.Product) error {
	err := d.getDb().Create(product).Error
	return errors.WithMessage(err, "CreateProduct failed")
}

// UpdateProduct 更新产品
func (d *ProductDomain) UpdateProduct(product *model.Product) error {
	err := d.getDb().Where("id = ?", product.ID).Updates(product).Error
	return errors.WithMessage(err, "UpdateProduct failed")
}

// DeleteProduct 删除产品
func (d *ProductDomain) DeleteProduct(id int) error {
	err := d.getDb().Where("id = ?", id).Update("is_deleted", consts.IsDeleted).Error
	return errors.WithMessage(err, "DeleteProduct failed")
}

// GetProductByID 获取产品
func (d *ProductDomain) GetProductByID(id int) (*model.Product, error) {
	var product model.Product
	err := d.getDb().Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, errors.WithMessage(err, "GetProductByID failed")
	}
	return &product, nil
}

// ListProduct 获取产品列表
func (d *ProductDomain) ListProduct() ([]model.Product, error) {
	var products []model.Product
	err := d.getDb().Where("is_deleted = ?", consts.NotDeleted).Find(&products).Error
	if err != nil {
		return nil, errors.WithMessage(err, "ListProduct failed")
	}
	return products, nil
}
