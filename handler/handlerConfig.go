package handler

import (
	"api-openings/config"

	"gorm.io/gorm"
)


var (
	logger 	*config.Logger
	db 			*gorm.DB
)

func InitializeHandler() {
	logger = config.NewLogger("handler")
  db = config.GetMysql()
}