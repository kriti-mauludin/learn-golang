package entity

import "time"

type TodoListCategory struct {
	ID          int64     `gorm:"column:id"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}

func (TodoListCategory) TableName() string {
	return "todo_list_categories"
}
