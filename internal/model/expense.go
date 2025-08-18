package model

import "time"

type Expense struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	UserId      uint      `gorm:"not null"`
	Note        string    `gorm:"not null"`
	Amount      float64   `gorm:"not null"`
	Category    int       `gorm:"not null"` // 1-餐饮, 2-交通, 3-娱乐, 4-购物, 5-住房, 6-其他
	ExpenseDate time.Time `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
