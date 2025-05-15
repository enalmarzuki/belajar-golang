package model

import "time"

type Wallet struct {
	ID        string    `gorm:"primary_key;column:id"`
	UserId    string    `gorm:"column:user_id"`
	Balance   int64     `gorm:"column:balance"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	// User      User      `gorm:"foreignKey:user_id;references:id"` // -> ada error "invalid recursive type User InvalidDeclCycle"
	// solusinya menggunakan pointer di type nya, seperti dibawah
	User *User `gorm:"foreignKey:user_id;references:id"`

	/*
		* NOTES :
		- keluar error InvalidDeclCycle karena struct Wallet butuh Struct User dan di Struct User butuh struct Wallet, jadi
		- Wallet -> User -> Wallet ... dst. sehingga tidak ada selasainya
	*/
}

func (w *Wallet) TableName() string {
	return "wallets"
}
