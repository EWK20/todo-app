package spec

import "gorm.io/gorm"

func HasItemLike(item string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("item LIKE '%' || ? || '%'", item)
	}

}

func IsCompleted(completed bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("completed=?", completed)
	}

}
