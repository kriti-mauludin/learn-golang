package entity

import "time"

type Role struct {
	Role        string    `gorm:"column:role"`
	UserID      int64     `gorm:"column:user_id"`
	Description string    `gorm:"column:description"`
	DoingAt     time.Time `gorm:"column:doing_at"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	ID          int64     `gorm:"column:id"`
}

func (Role) TableName() string {
	return "role"
}
