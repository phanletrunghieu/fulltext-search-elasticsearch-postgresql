package domain

import (
	"github.com/jinzhu/gorm"
)

// Post .
type Post struct {
	gorm.Model
	Title   string
	Content string
}
