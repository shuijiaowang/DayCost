package model

import (
	"time"

	"gorm.io/gorm"
)

type ExpenseExt struct {
	ID            int            `gorm:"primaryKey;autoIncrement;comment:扩展ID"`
	ExpenseID     int            `gorm:"uniqueIndex;not null;comment:关联消费ID"`
	ExpenseType   int8           `gorm:"not null;comment:类型(0:时间型,1:数量型)"`
	StartDate     time.Time      `gorm:"type:date;not null;comment:开始使用日期"`
	EstimatedDays int16          `gorm:"comment:预计使用天数(时间型)"`
	EndDate       *time.Time     `gorm:"type:date;comment:实际结束日期"`
	TotalQuantity float64        `gorm:"type:decimal(10,4);comment:总数量(数量型)"`
	Remaining     float64        `gorm:"type:decimal(10,4);comment:剩余量(数量型)"`
	Status        int8           `gorm:"default:1;comment:状态(0:进行中,1:已结束)"`
	CreatedAt     time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt     time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:更新时间"`
	DeletedAt     gorm.DeletedAt `gorm:"index;type:timestamp;comment:删除时间(软删除标志)"`
	ExpenseUsages []ExpenseUsage `gorm:"foreignKey:ExtendedID"` // 一对多关系
}
