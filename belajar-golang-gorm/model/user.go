package model

import "time"

/*
 1. Default convention di GORM
    - Saat membuat struct, secara default GORM akan melakukan mapping secara otomatis, dimana nama tabel akan dipilih dari nama Struct menggunakan lower_case jamak(banyak).
    ex: struct User => table users, struct OrderDetail => table order_details

    - sedangkan nama kolom akan dipilih menggunakan lower_case.

    - Selain itu, secara otomatis GORM akan memilih field ID sebagai primary key
    ex: struct User mempunyai field ID (uppercase) otomatis jadi primary key

2. Field Permission
  - <- : Write permission, <-:create untuk create only, <-:update untuk update only, <- untuk create dan update
  - -> : Read permission, ->:false untuk no read permission
  - -  : ignore field ini, tidak ada read/write permission
*/
type User struct {
	ID          string    `gorm:"primary_key;column:id;<-:create"`
	Password    string    `gorm:"column:password"`
	Name        Name      `gorm:"embedded"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime:<-:create"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Information string    `gorm:"-"`
	Wallet      Wallet    `gorm:"foreignKey:user_id;references:id"` // -> relasi one to one (hasOne)
}

// method dibawah ini dipakai klo kita mau menggunakan nama table nya secara manaual, maka harus menggunakan method TableName
func (u *User) TableName() string {
	return "users"
}

type Name struct {
	FirstName  string `gorm:"column:first_name"`
	MiddleName string `gorm:"column:middle_name"`
	LastName   string `gorm:"column:last_name"`
}
