package models

import "gorm.io/gorm"

type Record struct {
	gorm.Model
	Genotype string
	Age      string
	Address  string
}

type RecordPayload struct {
	Genotype string `json:"genotype" validate:"required"`
	Age      string `json:"age" validate:"required"`
	Address  string `json:"address" validate:"required"`
}
