package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `gorm:"primaryKey;autoIncrement;comment:用户ID"`
	Username  string         `gorm:"type:varchar(50);uniqueIndex;not null;comment:用户名"`
	Password  string         `gorm:"type:varchar(100);not null;comment:加密密码"`
	CreatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;type:timestamp;comment:删除时间(软删除标志)"`
}
