package review

import (
	"gorm.io/gorm"
)

type Reviews struct {
	gorm.Model
	Rating    int    `gorm:"type:int(1)"`
	Review    string `gorm:"type:text"`
	ServiceID uint
	UserID    uint
	OrderID   uint
}
