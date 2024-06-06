package _model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ProductMgr struct {
	*_BaseMgr
}

// ProductMgr open func
func ProductMgr(db *gorm.DB) *_ProductMgr {
	if db == nil {
		panic(fmt.Errorf("ProductMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductMgr{_BaseMgr: &_BaseMgr{DB: db.Table("product"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductMgr) GetTableName() string {
	return "product"
}

// Reset 重置gorm会话
func (obj *_ProductMgr) Reset() *_ProductMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ProductMgr) Get() (result Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductMgr) Gets() (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ProductMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Product{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 Primary Key
func (obj *_ProductMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 产品名称
func (obj *_ProductMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPrice price获取 价格
func (obj *_ProductMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithDescription description获取 描述
func (obj *_ProductMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithIsDeleted is_deleted获取 是否删除 0:未删除 1:已删除
func (obj *_ProductMgr) WithIsDeleted(isDeleted int8) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithCreateTime create_time获取 创建时间
func (obj *_ProductMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_ProductMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_ProductMgr) GetByOption(opts ...Option) (result Product, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ProductMgr) GetByOptions(opts ...Option) (results []*Product, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 Primary Key
func (obj *_ProductMgr) GetFromID(id int) (result Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 Primary Key
func (obj *_ProductMgr) GetBatchFromID(ids []int) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 产品名称
func (obj *_ProductMgr) GetFromName(name string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 产品名称
func (obj *_ProductMgr) GetBatchFromName(names []string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 价格
func (obj *_ProductMgr) GetFromPrice(price float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 价格
func (obj *_ProductMgr) GetBatchFromPrice(prices []float64) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容 描述
func (obj *_ProductMgr) GetFromDescription(description string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`description` = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量查找 描述
func (obj *_ProductMgr) GetBatchFromDescription(descriptions []string) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`description` IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容 是否删除 0:未删除 1:已删除
func (obj *_ProductMgr) GetFromIsDeleted(isDeleted int8) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找 是否删除 0:未删除 1:已删除
func (obj *_ProductMgr) GetBatchFromIsDeleted(isDeleteds []int8) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_ProductMgr) GetFromCreateTime(createTime time.Time) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_ProductMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_ProductMgr) GetFromUpdateTime(updateTime time.Time) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_ProductMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ProductMgr) FetchByPrimaryKey(id int) (result Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Where("`id` = ?", id).First(&result).Error

	return
}
