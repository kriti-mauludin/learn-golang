package entity

import "time"

type TodoListCategory struct {
	ID          int64     `gorm:"column:id"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	CreatedBy   int64     `gorm:"column:created_by"`
}

func (TodoListCategory) TableName() string {
	return "todo_list_categories"
}
