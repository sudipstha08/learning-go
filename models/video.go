package models

import "time"

type Person struct {
	ID        uint64 `gorm:"primary_key;auto-increment" json:"id"`
	FirstName string `json:"first_name" binding:"required" gorm:"type:varchar(32)"`
	LastName  string `json:"last_name" binding:"required" gorm:"type:varchar(32)"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"` // greater than or equals to 1
	Email     string `json:"email" validate:"required,email" gorm:"type:varchar(256)"`
}

type Video struct {
	ID uint64 `gorm:"primary_key;auto_increment" json:"id"`
	// Title       string `json:"title" binding:"min=2,max=200" validate:"is-cool" gorm:type:varchar(100)`
	Title       string    `json:"title" binding:"min=2,max=100" gorm:"type:varchar(100)"`
	Description string    `json:"description" binding:"max=200" gorm:"type:varchar(200)"`
	URL         string    `json:"url" binding:"required,url" gorm:"type:varchar(256);UNIQUE"`
	Author      Person    `json:"author" binding:"required" gorm:"association_autocreate:true;association_autoupdate:false;foreignkey:PersonID"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"-" gorm:"default: CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"default: CURRENT_TIMESTAMP" json:"updated_at"`
}
