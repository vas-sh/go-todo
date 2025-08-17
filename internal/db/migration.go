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
	err = db.AutoMigrate(models.User{}, models.Task{}, models.UserActivation{}, models.TaskStatus{}, models.BotUser{})
	if err != nil {
		return err
	}
	err = db.Exec(`
	CREATE OR REPLACE FUNCTION public.log_create_task_status()
		RETURNS trigger
		LANGUAGE 'plpgsql'
		VOLATILE NOT LEAKPROOF
	AS $BODY$
		BEGIN
			INSERT INTO task_statuses(created_at, task_id, status)
			SELECT now(), NEW.id, NEW.status; 
			RETURN NEW;
		END;
	$BODY$;

	CREATE OR REPLACE FUNCTION public.log_update_task_status()
		RETURNS trigger
		LANGUAGE 'plpgsql'
		VOLATILE NOT LEAKPROOF
	AS $BODY$
		BEGIN
			IF NEW.status != OLD.status THEN 
				INSERT INTO task_statuses(created_at, task_id, status)
				SELECT now(), NEW.id, NEW.status; 
			END IF;
			RETURN NEW;
		END;
	$BODY$;

	CREATE OR REPLACE TRIGGER create_task
		AFTER INSERT
		ON public.tasks
		FOR EACH ROW
		EXECUTE FUNCTION public.log_create_task_status();

	CREATE OR REPLACE TRIGGER update_task
		BEFORE UPDATE
		ON public.tasks
		FOR EACH ROW
		EXECUTE FUNCTION public.log_update_task_status();
	`).Error
	return err
}
