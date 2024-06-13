package entity

import "time"

type OrderDetails struct {
	ID              int       `gorm:"column:id;primaryKey;autoIncrement"`
	OrderPID        string    `gorm:"column:order_pid;not null;"`
	CustomerDetails string    `gorm:"column:customer_details"`
	ItemDetails     string    `gorm:"column:item_details"`
	OrderStatus     string    `gorm:"column:order_status"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}
