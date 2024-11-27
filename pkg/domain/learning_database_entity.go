package domain

import (
	"gorm.io/gorm"
	"time"
)

/*
 - In this file I am creating a model of database table
 - I am using the GORM library
 - In GORM for a field to be nullable, it has to be a pointer or use null types from database/sql package
*/

type MessageEntity struct {
	gorm.Model
	Id       int        `gorm:"primaryKey"`
	Content  string     `gorm:"size:100;not null"`
	DateTime *time.Time `gorm:"null"`
}
