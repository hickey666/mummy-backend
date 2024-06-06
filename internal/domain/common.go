package domain

import "gorm.io/gorm"

// 未删除
func NotDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("is_deleted = 0")
}
