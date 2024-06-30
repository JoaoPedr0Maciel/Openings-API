package config

import (
	"api-openings/schemas"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeMySql() (*gorm.DB, error) {
	logger := NewLogger("mysql")

	connection := "jpdev:12345678@tcp(localhost:3306)/openings?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		logger.Err("Failed to connect to database:", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Err("Failed to migrate database:", err)
	}

	// err = migrations.RemoveCollunmDeletedAt(db)
	// if err != nil {
	// 	logger.Err("Failed to remove column from database:", err)
	// }

	return db, nil
}
