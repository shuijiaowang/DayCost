package model

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	ID          int            `gorm:"primaryKey;autoIncrement;comment:消费ID"`
	UserID      int            `gorm:"not null;index;comment:关联用户ID"`
	Note        string         `gorm:"type:varchar(100);not null;comment:物品名称/消费摘要"`
	Amount      float64        `gorm:"type:decimal(10,2);not null;comment:消费金额"`
	Remarks     string         `gorm:"type:text;comment:详细备注(支持扩展标签)"`
	ExpenseDate time.Time      `gorm:"type:date;not null;comment:消费日期"`
	Category    int8           `gorm:"not null;comment:消费分类(0:餐饮,1:日用,2:交通...)"`
	IsExtended  bool           `gorm:"default:false;comment:是否为扩展消费"`
	CreatedAt   time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt   time.Time      `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"index;type:timestamp;comment:删除时间(软删除标志)"`
}

//type Expense struct {
//	ID          uint      `gorm:"primaryKey;autoIncrement"`
//	UserId      uint      `gorm:"not null"`
//	Note        string    `gorm:"not null"`
//	Amount      float64   `gorm:"not null"`
//	Category    int       `gorm:"not null"` // 1-餐饮, 2-交通, 3-娱乐, 4-购物, 5-住房, 6-其他
//	ExpenseDate time.Time `gorm:"not null"`
//	CreatedAt   time.Time `gorm:"autoCreateTime"`
//	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
//}
