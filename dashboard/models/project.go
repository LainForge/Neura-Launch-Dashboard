package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
    ID     uint   `gorm:"primaryKey"`
    UserID uint   `gorm:"index"`
    Name   string `gorm:"unique;not null"`
    Token  string `gorm:"unique;not null"`
}