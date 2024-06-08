package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	Balance uint
	Status  bool
	UserId  string
	User    User
}

type WalletStatusPayload struct {
	Status bool
}
