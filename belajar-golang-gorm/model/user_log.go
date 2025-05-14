package model

// before update column
// type UserLog struct {
// 	ID        int       `gorm:"primary_key;column:id;autoIncrement"`
// 	UserId    string    `gorm:"column:user_id"`
// 	Action    string    `gorm:"column:action"`
// 	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:<-:create"`
// 	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
// }

// after update column
type UserLog struct {
	ID        int    `gorm:"primary_key;column:id;autoIncrement"`
	UserId    string `gorm:"column:user_id"`
	Action    string `gorm:"column:action"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (u *UserLog) TableName() string {
	return "user_logs"
}
