package domain

import (
	"mummy-storage/internal/model"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ProductStorageLogDomain struct {
	dbEngine *gorm.DB
}

func NewProductStorageLogDomain(dbEngine *gorm.DB) *ProductStorageLogDomain {
	return &ProductStorageLogDomain{dbEngine: dbEngine}
}

func (d *ProductStorageLogDomain) getDb() *gorm.DB {
	return d.dbEngine.Table(new(model.ProductStorageLog).TableName())
}

// CreateProductStorageLog 创建产品存储日志
func (d *ProductStorageLogDomain) CreateProductStorageLog(productStorageLog *model.ProductStorageLog) error {
	err := d.getDb().Create(productStorageLog).Error
	return errors.WithMessage(err, "CreateProductStorageLog failed")
}

// UpdateProductStorageLog 更新产品存储日志
func (d *ProductStorageLogDomain) UpdateProductStorageLog(productStorageLog *model.ProductStorageLog) error {
	err := d.getDb().Where("id = ?", productStorageLog.ID).Updates(productStorageLog).Error
	return errors.WithMessage(err, "UpdateProductStorageLog failed")
}

// DeleteProductStorageLog 删除产品存储日志
func (d *ProductStorageLogDomain) DeleteProductStorageLog(id int) error {
	err := d.getDb().Where("id = ?", id).Update("is_deleted", 1).Error
	return errors.WithMessage(err, "DeleteProductStorageLog failed")
}

// GetProductStorageLogByID 获取产品存储日志
func (d *ProductStorageLogDomain) GetProductStorageLogByID(id int) (*model.ProductStorageLog, error) {
	var productStorageLog model.ProductStorageLog
	err := d.getDb().Where("id = ?", id).First(&productStorageLog).Error
	if err != nil {
		return nil, errors.WithMessage(err, "GetProductStorageLogByID failed")
	}
	return &productStorageLog, nil
}

// ListProductStorageLog 获取产品存储日志列表
func (d *ProductStorageLogDomain) ListProductStorageLog() ([]model.ProductStorageLog, error) {
	var productStorageLogs []model.ProductStorageLog
	err := d.getDb().Scopes(NotDeleted).Find(&productStorageLogs).Error
	if err != nil {
		return nil, errors.WithMessage(err, "ListProductStorageLog failed")
	}
	return productStorageLogs, nil
}
