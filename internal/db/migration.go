package db

import (
	"github.com/vas-sh/todo/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.Exec(`
	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'task_status') THEN
			CREATE TYPE task_status AS ENUM ('new', 'inProgress', 'done', 'canceled');
		END IF;
	END$$;`).Error
	if err != nil {
		return err
	}
	return db.AutoMigrate(models.Task{})
}
