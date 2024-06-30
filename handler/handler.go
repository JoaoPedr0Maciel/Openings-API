package handler

import (
	"api-openings/config"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
)

func InitializeHandler() {
	db = config.GetMysql()
}
