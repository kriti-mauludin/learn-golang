package entity

import "time"

type Job struct {
	Name        string    `gorm:"column:name"`
	UserID      int64     `gorm:"column:user_id"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	ID          int64     `gorm:"column:id"`
}

func (Job) TableName() string {
	return "job"
}
