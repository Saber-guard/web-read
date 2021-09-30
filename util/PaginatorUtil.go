package util

import "gorm.io/gorm"

type PaginatorUtil struct {
	Page uint
	Size uint
}

func (p PaginatorUtil) Exec(db *gorm.DB) *gorm.DB {
	if p.Page == 0 {
		p.Page = 1
	}
	if p.Size == 0 || p.Size > 100 {
		p.Size = 10
	}

	offset := (int)((p.Page - 1) * p.Size)
	size := (int)(p.Size)
	return db.Offset(offset).Limit(size)
}
