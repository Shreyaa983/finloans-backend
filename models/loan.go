package models

import (
	"gorm.io/gorm"
)

type LoanApplication struct {
	gorm.Model
	UserID     uint    `json:"user_id" gorm:"not null"`
	Amount     float64 `json:"amount" gorm:"not null"`
	Status     string  `json:"status" gorm:"not null;default:'pending'"`
	Tenure     int     `json:"tenure" gorm:"not null"` // in months
	IsEligible bool    `json:"is_eligible" gorm:"not null;default:false"`
	Reason     string  `json:"reason,omitempty" gorm:"type:text"`
	PdfPath    string  `json:"pdf_path,omitempty" gorm:"type:varchar(255)"`
}
