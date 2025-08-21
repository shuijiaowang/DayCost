package model

import (
	"time"

	"gorm.io/gorm"
)

type ExpenseUsage struct {
	ID         int            `gorm:"primaryKey;autoIncrement;comment:使用记录ID"`
	ExtendedID int            `gorm:"not null;index;comment:关联扩展ID"`
	UseDate    JSONDate       `gorm:"type:date;not null;comment:使用日期"`
	UsedValue  float64        `gorm:"type:decimal(10,4);not null;comment:消耗值(数量/比例)"`
	Notes      string         `gorm:"type:varchar(100);comment:使用备注"`
	CreatedAt  time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt  time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:更新时间"`
	DeletedAt  gorm.DeletedAt `gorm:"index;type:timestamp;comment:删除时间(软删除标志)"`
}
