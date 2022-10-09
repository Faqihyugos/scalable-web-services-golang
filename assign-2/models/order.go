package models

import "time"

type Order struct {
	ID           uint      `gorm:"primaryKey;column:order_id"`
	CustomerName string    `gorm:"column:customer_name" json:"customerName"`
	OrderedAt    time.Time `gorm:"column:ordered_at" json:"orderedAt"`
	Items        []Item    `gorm:"foreignKey:order_id" json:"items"`
}
