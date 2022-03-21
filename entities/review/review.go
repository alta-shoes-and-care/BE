package review

import (
	"gorm.io/gorm"
)

type Reviews struct {
	gorm.Model
	Rating      int    `gorm:"type:int(1)"`
	Review      string `gorm:"type:text"`
	HasReviewed bool   `gorm:"type:boolean;default:true"`
	ServiceID   uint
	UserID      uint
	OrderID     uint
}
