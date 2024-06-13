package entity

import "time"

type CustomerDetails struct {
	ID              int       `gorm:"column:id;primaryKey;autoIncrement"`
	CustomerPID     string    `gorm:"column:customer_pid;not null;"`
	CustomerDetails string    `gorm:"column:customer_details"`
	UserName        string    `gorm:"column:user_name"`
	Password        string    `gorm:"column:password"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
	DeletedAt       time.Time `gorm:"column:deleted_at"`
}
