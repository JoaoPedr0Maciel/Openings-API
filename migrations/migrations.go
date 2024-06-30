package migrations

import (
	"api-openings/schemas"

	"gorm.io/gorm"
)

func RemoveCollunmDeletedAt(db *gorm.DB) error {
	err := db.Migrator().DropColumn(&schemas.Opening{}, "deleted_at")
	return err
}