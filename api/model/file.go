package model

import "time"

type File struct {
	ID 		  int 		`gorm:"primary_key;not_null;autoIncrement:true"`
	Name      string    `gorm:"type:text;not null"`
	Path	  string 	`gorm:"type:text;not null"`
	Timestamp time.Time `sql:"DEFAULT:current_timestamp"`
}