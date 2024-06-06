package _model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CategoryMgr struct {
	*_BaseMgr
}

// CategoryMgr open func
func CategoryMgr(db *gorm.DB) *_CategoryMgr {
	if db == nil {
		panic(fmt.Errorf("CategoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CategoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("category"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CategoryMgr) GetTableName() string {
	return "category"
}

// Reset 重置gorm会话
func (obj *_CategoryMgr) Reset() *_CategoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CategoryMgr) Get() (result Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Category{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CategoryMgr) Gets() (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Category{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CategoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Category{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 Primary Key
func (obj *_CategoryMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 分类名称
func (obj *_CategoryMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDescription description获取 描述
func (obj *_CategoryMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithIsDeleted is_deleted获取 是否删除 0:未删除 1:已删除
func (obj *_CategoryMgr) WithIsDeleted(isDeleted int8) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithCreateTime create_time获取 创建时间
func (obj *_CategoryMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_CategoryMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_CategoryMgr) GetByOption(opts ...Option) (result Category, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Category{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CategoryMgr) GetByOptions(opts ...Option) (results []*Category, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Category{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 Primary Key
func (obj *_CategoryMgr) GetFromID(id int) (result Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Category{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 Primary Key
func (obj *_CategoryMgr) GetBatchFromID(ids []int) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Category{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 分类名称
func (obj *_CategoryMgr) GetFromName(name string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Category{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 分类名称
func (obj *_CategoryMgr) GetBatchFromName(names []string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Category{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容 描述
func (obj *_CategoryMgr) GetFromDescription(description string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Category{}).Where("`description` = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量查找 描述
func (obj *_CategoryMgr) GetBatchFromDescription(descriptions []string) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Category{}).Where("`description` IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容 是否删除 0:未删除 1:已删除
func (obj *_CategoryMgr) GetFromIsDeleted(isDeleted int8) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Category{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找 是否删除 0:未删除 1:已删除
func (obj *_CategoryMgr) GetBatchFromIsDeleted(isDeleteds []int8) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Category{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_CategoryMgr) GetFromCreateTime(createTime time.Time) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Category{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_CategoryMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Category{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_CategoryMgr) GetFromUpdateTime(updateTime time.Time) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Category{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_CategoryMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Category{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_CategoryMgr) FetchByPrimaryKey(id int) (result Category, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Category{}).Where("`id` = ?", id).First(&result).Error

	return
}
