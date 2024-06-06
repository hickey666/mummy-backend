package domain

import (
	"mummy-storage/internal/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ProductStorageDomain struct {
	dbEngine *gorm.DB
}

func NewProductStorageDomain(dbEngine *gorm.DB) *ProductStorageDomain {
	return &ProductStorageDomain{dbEngine: dbEngine}
}

func (d *ProductStorageDomain) getDb() *gorm.DB {
	return d.dbEngine.Table(new(model.ProductStorage).TableName())
}

// CreateProductStorage 创建产品存储
func (d *ProductStorageDomain) CreateProductStorage(productStorage *model.ProductStorage) error {
	err := d.getDb().Create(productStorage).Error
	return errors.WithMessage(err, "CreateProductStorage failed")
}

// UpdateProductStorage 更新产品存储
func (d *ProductStorageDomain) UpdateProductStorage(productStorage *model.ProductStorage) error {
	err := d.getDb().Where("id = ?", productStorage.ID).Updates(productStorage).Error
	return errors.WithMessage(err, "UpdateProductStorage failed")
}

// DeleteProductStorage 删除产品存储
func (d *ProductStorageDomain) DeleteProductStorage(id int) error {
	err := d.getDb().Where("id = ?", id).Update("is_deleted", 1).Error
	return errors.WithMessage(err, "DeleteProductStorage failed")
}

// GetProductStorageByID 获取产品存储
func (d *ProductStorageDomain) GetProductStorageByID(id int) (*model.ProductStorage, error) {
	var productStorage model.ProductStorage
	err := d.getDb().Where("id = ?", id).First(&productStorage).Error
	if err != nil {
		return nil, errors.WithMessage(err, "GetProductStorageByID failed")
	}
	return &productStorage, nil
}

// ListProductStorage 获取产品存储列表
func (d *ProductStorageDomain) ListProductStorage() ([]model.ProductStorage, error) {
	var productStorages []model.ProductStorage
	err := d.getDb().Scopes(NotDeleted).Find(&productStorages).Error
	if err != nil {
		return nil, errors.WithMessage(err, "ListProductStorage failed")
	}
	return productStorages, nil
}
