package _model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ProductCategoryMgr struct {
	*_BaseMgr
}

// ProductCategoryMgr open func
func ProductCategoryMgr(db *gorm.DB) *_ProductCategoryMgr {
	if db == nil {
		panic(fmt.Errorf("ProductCategoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductCategoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("product_category"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductCategoryMgr) GetTableName() string {
	return "product_category"
}

// Reset 重置gorm会话
func (obj *_ProductCategoryMgr) Reset() *_ProductCategoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ProductCategoryMgr) Get() (result ProductCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductCategory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductCategoryMgr) Gets() (results []*ProductCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductCategory{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ProductCategoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ProductCategory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 Primary Key
func (obj *_ProductCategoryMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPid pid获取 产品ID
func (obj *_ProductCategoryMgr) WithPid(pid int) Option {
	return optionFunc(func(o *options) { o.query["pid"] = pid })
}

// WithCid cid获取 分类ID
func (obj *_ProductCategoryMgr) WithCid(cid int) Option {
	return optionFunc(func(o *options) { o.query["cid"] = cid })
}

// WithIsDeleted is_deleted获取 是否删除 0:未删除 1:已删除
func (obj *_ProductCategoryMgr) WithIsDeleted(isDeleted int8) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithCreateTime create_time获取 创建时间
func (obj *_ProductCategoryMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_ProductCategoryMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_ProductCategoryMgr) GetByOption(opts ...Option) (result ProductCategory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ProductCategory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ProductCategoryMgr) GetByOptions(opts ...Option) (results []*ProductCategory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ProductCategory{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 Primary Key
func (obj *_ProductCategoryMgr) GetFromID(id int) (result ProductCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductCategory{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 Primary Key
func (obj *_ProductCategoryMgr) GetBatchFromID(ids []int) (results []*ProductCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductCategory{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromPid 通过pid获取内容 产品ID
func (obj *_ProductCategoryMgr) GetFromPid(pid int) (results []*ProductCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductCategory{}).Where("`pid` = ?", pid).Find(&results).Error

	return
}

// GetBatchFromPid 批量查找 产品ID
func (obj *_ProductCategoryMgr) GetBatchFromPid(pids []int) (results []*ProductCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductCategory{}).Where("`pid` IN (?)", pids).Find(&results).Error

	return
}

// GetFromCid 通过cid获取内容 分类ID
func (obj *_ProductCategoryMgr) GetFromCid(cid int) (results []*ProductCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductCategory{}).Where("`cid` = ?", cid).Find(&results).Error

	return
}

// GetBatchFromCid 批量查找 分类ID
func (obj *_ProductCategoryMgr) GetBatchFromCid(cids []int) (results []*ProductCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductCategory{}).Where("`cid` IN (?)", cids).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容 是否删除 0:未删除 1:已删除
func (obj *_ProductCategoryMgr) GetFromIsDeleted(isDeleted int8) (results []*ProductCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductCategory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找 是否删除 0:未删除 1:已删除
func (obj *_ProductCategoryMgr) GetBatchFromIsDeleted(isDeleteds []int8) (results []*ProductCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductCategory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_ProductCategoryMgr) GetFromCreateTime(createTime time.Time) (results []*ProductCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductCategory{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_ProductCategoryMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*ProductCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductCategory{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_ProductCategoryMgr) GetFromUpdateTime(updateTime time.Time) (results []*ProductCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductCategory{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_ProductCategoryMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*ProductCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductCategory{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ProductCategoryMgr) FetchByPrimaryKey(id int) (result ProductCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductCategory{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByProductIDCategoryID primary or index 获取唯一内容
func (obj *_ProductCategoryMgr) FetchUniqueIndexByProductIDCategoryID(pid int, cid int) (result ProductCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductCategory{}).Where("`pid` = ? AND `cid` = ?", pid, cid).First(&result).Error

	return
}
