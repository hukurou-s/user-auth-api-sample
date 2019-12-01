package database

import "github.com/jinzhu/gorm"

type SqlHandler interface {
	Where(interface{}, ...interface{}) *gorm.DB
	First(interface{}, ...interface{}) *gorm.DB
}
