package model

import (
	"time"
)

type BaseModel struct {
	ID        int64     `gorm:"primarykey;comment:主键;"`
	CreatedAt time.Time `gorm:"type:timestamp;comment:创建时间;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;comment:更新时间;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
