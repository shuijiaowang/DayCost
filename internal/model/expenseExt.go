package model

import (
	"time"

	"gorm.io/gorm"
)

// 这里不需要用户id，那这个的增删改查岂不是可以无限制？用户A发起向用户B修改的请求
type ExpenseExt struct {
	ID               int            `gorm:"primaryKey;autoIncrement;comment:扩展ID"`
	ExpenseID        int            `gorm:"uniqueIndex;not null;comment:关联消费ID"` // 非指针（必填）
	ExpenseType      int8           `gorm:"not null;comment:类型(0:时间型,1:数量型)"`    // 非指针（必填）
	StartDate        JSONDate       `gorm:"type:date;not null;comment:开始使用日期"`   // 非指针（必填）
	EstimatedEndDate *JSONDate      `gorm:"comment:预计使用天数(时间型)"`                 // 指针（可选）
	EndDate          *JSONDate      `gorm:"type:date;comment:实际结束日期"`            // 指针（可选）
	TotalQuantity    *float64       `gorm:"type:decimal(10,4);comment:总数量(数量型)"` // 指针（可选）
	Remaining        *float64       `gorm:"type:decimal(10,4);comment:剩余量(数量型)"` // 指针（可选）
	Status           int8           `gorm:"default:1;comment:状态(0:进行中,1:已结束)"`   // 非指针（有默认值）
	CreatedAt        time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt        time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:更新时间"`
	DeletedAt        gorm.DeletedAt `gorm:"index;type:timestamp;comment:删除时间(软删除标志)"`
}
