package bot

import (
	"context"
	"errors"
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vas-sh/todo/internal/models"
)

type keyboardOps struct {
	exist bool
}

func (s *srv) sendTaskMessage(ctx context.Context, chatID int64) error {
	task, nextExist, err := s.getTask(ctx, chatID, 0)
	if errors.Is(err, models.ErrNotFound) {
		return s.sendTextMessage(chatID, "You don't have any tasks yet")
	}
	if err != nil {
		return err
	}
	keyboard := s.keyboard(0, keyboardOps{exist: nextExist})
	taskBody := s.taskFormat(task)
	msg := tgbotapi.NewMessage(chatID, taskBody)
	msg.ReplyMarkup = keyboard
	msg.ParseMode = tgbotapi.ModeHTML
	_, err = s.bot.Send(msg)
	return err
}

func (s *srv) refreshTaskMessage(ctx context.Context, chatID int64, messageID, page int) error {
	task, nextExist, err := s.getTask(ctx, chatID, int64(page))
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error())
		_, err = s.bot.Send(tgbotapi.NewEditMessageText(chatID, messageID, "You don't have the next task"))
		return err
	}
	keyboard := s.keyboard(page, keyboardOps{exist: nextExist})
	taskBody := s.taskFormat(task)
	msg := tgbotapi.NewEditMessageText(chatID, messageID, taskBody)
	msg.ParseMode = tgbotapi.ModeHTML
	msg.ReplyMarkup = &keyboard
	_, err = s.bot.Send(msg)
	return err
}

func (s *srv) getTask(ctx context.Context, chatID, page int64) (models.Task, bool, error) {
	userID, err := s.userSrv.GetUserID(ctx, chatID)
	if err != nil {
		return models.Task{}, false, err
	}
	return s.taskSrv.GetTask(ctx, userID, page)
}

func (s *srv) taskFormat(task models.Task) string {
	taskBody := fmt.Sprintf("<b>📝 %s</b>\n", task.Title)
	if task.Description != "" {
		taskBody = fmt.Sprintf(taskBody+"📎 <i>%s</i>\n", task.Description)
	}
	taskBody += s.statusFormat(task)
	if task.EstimateTime == nil {
		return taskBody
	}
	taskBody += s.timeFormat(task)
	return taskBody
}

func (*srv) statusFormat(task models.Task) string {
	switch task.Status {
	case models.NewStatus:
		return "🟡 New"
	case models.InProgressStatus:
		return "🟡 In progress"
	case models.DoneStatus:
		return "🟢 Done"
	case models.CanceledStatus:
		return "🔴 Canceled"
	}
	return ""
}

func (*srv) timeFormat(task models.Task) string {
	now := time.Now()
	duration := task.EstimateTime.Sub(now)
	days := duration.Hours() / 24
	return fmt.Sprintf("\nDeadline: %s (%.0f days)", task.EstimateTime.Format("02/01/2006 15:04"), days)
}

func (*srv) keyboard(page int, ops keyboardOps) tgbotapi.InlineKeyboardMarkup {
	var buttons []tgbotapi.InlineKeyboardButton
	if page > 0 {
		buttons = append(buttons, tgbotapi.
			NewInlineKeyboardButtonData("⬅️ Prev", fmt.Sprintf("%s:%d", models.PrevButtonType, page-1)))
	}
	if ops.exist {
		buttons = append(buttons, tgbotapi.
			NewInlineKeyboardButtonData("➡️ Next", fmt.Sprintf("%s:%d", models.NextButtonType, page+1)))
	}
	if len(buttons) == 0 {
		return tgbotapi.InlineKeyboardMarkup{InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{}}
	}
	return tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(buttons...))
}
