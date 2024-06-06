package _model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ProductStorageMgr struct {
	*_BaseMgr
}

// ProductStorageMgr open func
func ProductStorageMgr(db *gorm.DB) *_ProductStorageMgr {
	if db == nil {
		panic(fmt.Errorf("ProductStorageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductStorageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("product_storage"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductStorageMgr) GetTableName() string {
	return "product_storage"
}

// Reset 重置gorm会话
func (obj *_ProductStorageMgr) Reset() *_ProductStorageMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ProductStorageMgr) Get() (result ProductStorage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorage{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductStorageMgr) Gets() (results []*ProductStorage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorage{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ProductStorageMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ProductStorage{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 Primary Key
func (obj *_ProductStorageMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPid pid获取 产品ID
func (obj *_ProductStorageMgr) WithPid(pid int) Option {
	return optionFunc(func(o *options) { o.query["pid"] = pid })
}

// WithTotal total获取 总库存
func (obj *_ProductStorageMgr) WithTotal(total int) Option {
	return optionFunc(func(o *options) { o.query["total"] = total })
}

// WithUsed used获取 已用库存
func (obj *_ProductStorageMgr) WithUsed(used int) Option {
	return optionFunc(func(o *options) { o.query["used"] = used })
}

// WithResidue residue获取 剩余库存
func (obj *_ProductStorageMgr) WithResidue(residue int) Option {
	return optionFunc(func(o *options) { o.query["residue"] = residue })
}

// WithCreateTime create_time获取 创建时间
func (obj *_ProductStorageMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_ProductStorageMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_ProductStorageMgr) GetByOption(opts ...Option) (result ProductStorage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ProductStorage{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ProductStorageMgr) GetByOptions(opts ...Option) (results []*ProductStorage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ProductStorage{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 Primary Key
func (obj *_ProductStorageMgr) GetFromID(id int) (result ProductStorage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorage{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 Primary Key
func (obj *_ProductStorageMgr) GetBatchFromID(ids []int) (results []*ProductStorage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorage{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromPid 通过pid获取内容 产品ID
func (obj *_ProductStorageMgr) GetFromPid(pid int) (results []*ProductStorage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorage{}).Where("`pid` = ?", pid).Find(&results).Error

	return
}

// GetBatchFromPid 批量查找 产品ID
func (obj *_ProductStorageMgr) GetBatchFromPid(pids []int) (results []*ProductStorage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorage{}).Where("`pid` IN (?)", pids).Find(&results).Error

	return
}

// GetFromTotal 通过total获取内容 总库存
func (obj *_ProductStorageMgr) GetFromTotal(total int) (results []*ProductStorage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorage{}).Where("`total` = ?", total).Find(&results).Error

	return
}

// GetBatchFromTotal 批量查找 总库存
func (obj *_ProductStorageMgr) GetBatchFromTotal(totals []int) (results []*ProductStorage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorage{}).Where("`total` IN (?)", totals).Find(&results).Error

	return
}

// GetFromUsed 通过used获取内容 已用库存
func (obj *_ProductStorageMgr) GetFromUsed(used int) (results []*ProductStorage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorage{}).Where("`used` = ?", used).Find(&results).Error

	return
}

// GetBatchFromUsed 批量查找 已用库存
func (obj *_ProductStorageMgr) GetBatchFromUsed(useds []int) (results []*ProductStorage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorage{}).Where("`used` IN (?)", useds).Find(&results).Error

	return
}

// GetFromResidue 通过residue获取内容 剩余库存
func (obj *_ProductStorageMgr) GetFromResidue(residue int) (results []*ProductStorage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorage{}).Where("`residue` = ?", residue).Find(&results).Error

	return
}

// GetBatchFromResidue 批量查找 剩余库存
func (obj *_ProductStorageMgr) GetBatchFromResidue(residues []int) (results []*ProductStorage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorage{}).Where("`residue` IN (?)", residues).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_ProductStorageMgr) GetFromCreateTime(createTime time.Time) (results []*ProductStorage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorage{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_ProductStorageMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*ProductStorage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorage{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_ProductStorageMgr) GetFromUpdateTime(updateTime time.Time) (results []*ProductStorage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorage{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_ProductStorageMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*ProductStorage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorage{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ProductStorageMgr) FetchByPrimaryKey(id int) (result ProductStorage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorage{}).Where("`id` = ?", id).First(&result).Error

	return
}
