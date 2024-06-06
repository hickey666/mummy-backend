package _model

import (
	"time"
)

// Category 分类表
type Category struct {
	ID          int       `gorm:"primaryKey;column:id" json:"-"`         // Primary Key
	Name        string    `gorm:"column:name" json:"name"`               // 分类名称
	Description string    `gorm:"column:description" json:"description"` // 描述
	IsDeleted   int8      `gorm:"column:is_deleted" json:"isDeleted"`    // 是否删除 0:未删除 1:已删除
	CreateTime  time.Time `gorm:"column:create_time" json:"createTime"`  // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time" json:"updateTime"`  // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *Category) TableName() string {
	return "category"
}

// CategoryColumns get sql column name.获取数据库列名
var CategoryColumns = struct {
	ID          string
	Name        string
	Description string
	IsDeleted   string
	CreateTime  string
	UpdateTime  string
}{
	ID:          "id",
	Name:        "name",
	Description: "description",
	IsDeleted:   "is_deleted",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

// Product 产品表
type Product struct {
	ID          int       `gorm:"primaryKey;column:id" json:"-"`         // Primary Key
	Name        string    `gorm:"column:name" json:"name"`               // 产品名称
	Price       float64   `gorm:"column:price" json:"price"`             // 价格
	Description string    `gorm:"column:description" json:"description"` // 描述
	IsDeleted   int8      `gorm:"column:is_deleted" json:"isDeleted"`    // 是否删除 0:未删除 1:已删除
	CreateTime  time.Time `gorm:"column:create_time" json:"createTime"`  // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time" json:"updateTime"`  // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *Product) TableName() string {
	return "product"
}

// ProductColumns get sql column name.获取数据库列名
var ProductColumns = struct {
	ID          string
	Name        string
	Price       string
	Description string
	IsDeleted   string
	CreateTime  string
	UpdateTime  string
}{
	ID:          "id",
	Name:        "name",
	Price:       "price",
	Description: "description",
	IsDeleted:   "is_deleted",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

// ProductCategory 产品分类关联表
type ProductCategory struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`        // Primary Key
	Pid        int       `gorm:"column:pid" json:"pid"`                // 产品ID
	Cid        int       `gorm:"column:cid" json:"cid"`                // 分类ID
	IsDeleted  int8      `gorm:"column:is_deleted" json:"isDeleted"`   // 是否删除 0:未删除 1:已删除
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *ProductCategory) TableName() string {
	return "product_category"
}

// ProductCategoryColumns get sql column name.获取数据库列名
var ProductCategoryColumns = struct {
	ID         string
	Pid        string
	Cid        string
	IsDeleted  string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	Pid:        "pid",
	Cid:        "cid",
	IsDeleted:  "is_deleted",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// ProductStorage 产品库存表
type ProductStorage struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`        // Primary Key
	Pid        int       `gorm:"column:pid" json:"pid"`                // 产品ID
	Total      int       `gorm:"column:total" json:"total"`            // 总库存
	Used       int       `gorm:"column:used" json:"used"`              // 已用库存
	Residue    int       `gorm:"column:residue" json:"residue"`        // 剩余库存
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *ProductStorage) TableName() string {
	return "product_storage"
}

// ProductStorageColumns get sql column name.获取数据库列名
var ProductStorageColumns = struct {
	ID         string
	Pid        string
	Total      string
	Used       string
	Residue    string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	Pid:        "pid",
	Total:      "total",
	Used:       "used",
	Residue:    "residue",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// ProductStorageLog 产品库存日志表
type ProductStorageLog struct {
	ID          int       `gorm:"primaryKey;column:id" json:"-"`          // Primary Key
	Pid         int       `gorm:"column:pid" json:"pid"`                  // 产品ID
	BeforeTotal int       `gorm:"column:before_total" json:"beforeTotal"` // 总库存 before
	PutIn       int       `gorm:"column:put_in" json:"putIn"`             // 新增库存
	PutOut      int       `gorm:"column:put_out" json:"putOut"`           // 减少库存
	AfterTotal  int       `gorm:"column:after_total" json:"afterTotal"`   // 总库存 after
	IsDeleted   int8      `gorm:"column:is_deleted" json:"isDeleted"`     // 是否删除 0:未删除 1:已删除
	CreateTime  time.Time `gorm:"column:create_time" json:"createTime"`   // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time" json:"updateTime"`   // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *ProductStorageLog) TableName() string {
	return "product_storage_log"
}

// ProductStorageLogColumns get sql column name.获取数据库列名
var ProductStorageLogColumns = struct {
	ID          string
	Pid         string
	BeforeTotal string
	PutIn       string
	PutOut      string
	AfterTotal  string
	IsDeleted   string
	CreateTime  string
	UpdateTime  string
}{
	ID:          "id",
	Pid:         "pid",
	BeforeTotal: "before_total",
	PutIn:       "put_in",
	PutOut:      "put_out",
	AfterTotal:  "after_total",
	IsDeleted:   "is_deleted",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}
