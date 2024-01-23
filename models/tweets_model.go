package models

import "gorm.io/gorm"

type Tweet struct {
	gorm.Model
	Author  string
	Content string
}
