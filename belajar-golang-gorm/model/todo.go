package model

import (
	"gorm.io/gorm"
)

// before using the gorm model
// type Todo struct {
// 	ID          int64          `gorm:"primary_key;column:id;autoIncrement"`
// 	UserId      string         `gorm:"column:user_id"`
// 	Title       string         `gorm:"column:title"`
// 	Description string         `gorm:"column:description"`
// 	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime"`
// 	UpdatedAt   time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
// 	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
// }

// after using the gorm model
type Todo struct {
	gorm.Model
	UserId      string `gorm:"column:user_id"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
}

func (t *Todo) TableName() string {
	return "todos"
}
