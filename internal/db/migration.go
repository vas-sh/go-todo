package db

import (
	"github.com/vas-sh/todo/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	//	err := db.Exec("CREATE TYPE task_status AS ENUM ('new', 'inProgress', 'done', 'canceled');").Error
	//if err != nil {
	//	return err
	//}
	return db.AutoMigrate(models.Task{})
}
