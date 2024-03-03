package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model

	Mail     string    `gorm:"type:varchar(40);unique" json:"mail,omitempty"`
	Password string    `gorm:"size:255" json:"password,omitempty"`
	Name     string    `gorm:"size:100" json:"name"`
	Surname  string    `gorm:"size:100" json:"surname"`
	Birthday time.Time `gorm:"type:time" json:"birthday"`
}
