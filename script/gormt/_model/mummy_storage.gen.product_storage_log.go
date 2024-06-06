package _model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ProductStorageLogMgr struct {
	*_BaseMgr
}

// ProductStorageLogMgr open func
func ProductStorageLogMgr(db *gorm.DB) *_ProductStorageLogMgr {
	if db == nil {
		panic(fmt.Errorf("ProductStorageLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductStorageLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("product_storage_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductStorageLogMgr) GetTableName() string {
	return "product_storage_log"
}

// Reset 重置gorm会话
func (obj *_ProductStorageLogMgr) Reset() *_ProductStorageLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ProductStorageLogMgr) Get() (result ProductStorageLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductStorageLogMgr) Gets() (results []*ProductStorageLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ProductStorageLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 Primary Key
func (obj *_ProductStorageLogMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPid pid获取 产品ID
func (obj *_ProductStorageLogMgr) WithPid(pid int) Option {
	return optionFunc(func(o *options) { o.query["pid"] = pid })
}

// WithBeforeTotal before_total获取 总库存 before
func (obj *_ProductStorageLogMgr) WithBeforeTotal(beforeTotal int) Option {
	return optionFunc(func(o *options) { o.query["before_total"] = beforeTotal })
}

// WithPutIn put_in获取 新增库存
func (obj *_ProductStorageLogMgr) WithPutIn(putIn int) Option {
	return optionFunc(func(o *options) { o.query["put_in"] = putIn })
}

// WithPutOut put_out获取 减少库存
func (obj *_ProductStorageLogMgr) WithPutOut(putOut int) Option {
	return optionFunc(func(o *options) { o.query["put_out"] = putOut })
}

// WithAfterTotal after_total获取 总库存 after
func (obj *_ProductStorageLogMgr) WithAfterTotal(afterTotal int) Option {
	return optionFunc(func(o *options) { o.query["after_total"] = afterTotal })
}

// WithIsDeleted is_deleted获取 是否删除 0:未删除 1:已删除
func (obj *_ProductStorageLogMgr) WithIsDeleted(isDeleted int8) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithCreateTime create_time获取 创建时间
func (obj *_ProductStorageLogMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_ProductStorageLogMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_ProductStorageLogMgr) GetByOption(opts ...Option) (result ProductStorageLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ProductStorageLogMgr) GetByOptions(opts ...Option) (results []*ProductStorageLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 Primary Key
func (obj *_ProductStorageLogMgr) GetFromID(id int) (result ProductStorageLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 Primary Key
func (obj *_ProductStorageLogMgr) GetBatchFromID(ids []int) (results []*ProductStorageLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromPid 通过pid获取内容 产品ID
func (obj *_ProductStorageLogMgr) GetFromPid(pid int) (results []*ProductStorageLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Where("`pid` = ?", pid).Find(&results).Error

	return
}

// GetBatchFromPid 批量查找 产品ID
func (obj *_ProductStorageLogMgr) GetBatchFromPid(pids []int) (results []*ProductStorageLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Where("`pid` IN (?)", pids).Find(&results).Error

	return
}

// GetFromBeforeTotal 通过before_total获取内容 总库存 before
func (obj *_ProductStorageLogMgr) GetFromBeforeTotal(beforeTotal int) (results []*ProductStorageLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Where("`before_total` = ?", beforeTotal).Find(&results).Error

	return
}

// GetBatchFromBeforeTotal 批量查找 总库存 before
func (obj *_ProductStorageLogMgr) GetBatchFromBeforeTotal(beforeTotals []int) (results []*ProductStorageLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Where("`before_total` IN (?)", beforeTotals).Find(&results).Error

	return
}

// GetFromPutIn 通过put_in获取内容 新增库存
func (obj *_ProductStorageLogMgr) GetFromPutIn(putIn int) (results []*ProductStorageLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Where("`put_in` = ?", putIn).Find(&results).Error

	return
}

// GetBatchFromPutIn 批量查找 新增库存
func (obj *_ProductStorageLogMgr) GetBatchFromPutIn(putIns []int) (results []*ProductStorageLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Where("`put_in` IN (?)", putIns).Find(&results).Error

	return
}

// GetFromPutOut 通过put_out获取内容 减少库存
func (obj *_ProductStorageLogMgr) GetFromPutOut(putOut int) (results []*ProductStorageLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Where("`put_out` = ?", putOut).Find(&results).Error

	return
}

// GetBatchFromPutOut 批量查找 减少库存
func (obj *_ProductStorageLogMgr) GetBatchFromPutOut(putOuts []int) (results []*ProductStorageLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Where("`put_out` IN (?)", putOuts).Find(&results).Error

	return
}

// GetFromAfterTotal 通过after_total获取内容 总库存 after
func (obj *_ProductStorageLogMgr) GetFromAfterTotal(afterTotal int) (results []*ProductStorageLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Where("`after_total` = ?", afterTotal).Find(&results).Error

	return
}

// GetBatchFromAfterTotal 批量查找 总库存 after
func (obj *_ProductStorageLogMgr) GetBatchFromAfterTotal(afterTotals []int) (results []*ProductStorageLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Where("`after_total` IN (?)", afterTotals).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容 是否删除 0:未删除 1:已删除
func (obj *_ProductStorageLogMgr) GetFromIsDeleted(isDeleted int8) (results []*ProductStorageLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找 是否删除 0:未删除 1:已删除
func (obj *_ProductStorageLogMgr) GetBatchFromIsDeleted(isDeleteds []int8) (results []*ProductStorageLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_ProductStorageLogMgr) GetFromCreateTime(createTime time.Time) (results []*ProductStorageLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_ProductStorageLogMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*ProductStorageLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_ProductStorageLogMgr) GetFromUpdateTime(updateTime time.Time) (results []*ProductStorageLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_ProductStorageLogMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*ProductStorageLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ProductStorageLogMgr) FetchByPrimaryKey(id int) (result ProductStorageLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ProductStorageLog{}).Where("`id` = ?", id).First(&result).Error

	return
}
