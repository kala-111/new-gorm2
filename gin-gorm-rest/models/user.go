package models

import "gorm.io/gorm"

type Mobile struct {
	gorm.Model
	Id      int    `json:"id" gorm:"primary key"`
	Name    int    `json:"name"`
	Version string `json:"version"`
	Cost    string `json:"cost"`
}
