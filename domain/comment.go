package domain

import "github.com/jinzhu/gorm"

// Comment .
type Comment struct {
	gorm.Model
	Content string
	PostID  uint
}
