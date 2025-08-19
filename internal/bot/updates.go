package bot

import (
	"context"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vas-sh/todo/internal/models"
)

func (s *srv) Updates(ctx context.Context) {
	s.createMenu(ctx)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := s.bot.GetUpdatesChan(u)
	for update := range updates {
		//nolint:contextcheck // separate context create for each request
		go s.manageUpdate(update)
	}
}

func (s *srv) createMenu(ctx context.Context) {
	commands := []tgbotapi.BotCommand{
		{
			Command:     "tasks",
			Description: "Show task list",
		},
		{
			Command:     "create",
			Description: "Create new task",
		},
	}
	cfg := tgbotapi.NewSetMyCommands(commands...)
	_, err := s.bot.Request(cfg)
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error())
	}
}

func (s *srv) manageUpdate(update tgbotapi.Update) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if update.Message != nil {
		s.handleTextMessage(ctx, update.Message)
		return
	}
	if update.CallbackQuery != nil {
		s.handleCallback(ctx, update.CallbackQuery)
	}
}

func (s *srv) handleTextMessage(ctx context.Context, message *tgbotapi.Message) {
	if strings.HasPrefix(message.Text, string(models.StartCommand)) {
		err := s.authorize(ctx, message.From.ID, message.Text)
		if err != nil {
			s.logger.ErrorContext(ctx, err.Error())
		}
		return
	}
	if message.Text == string(models.TaskListCommand) {
		err := s.sendTaskMessage(ctx, message.From.ID)
		if err != nil {
			s.logger.ErrorContext(ctx, err.Error())
		}
		return
	}
}

func (s *srv) handleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery) {
	_, err := s.handleTaskCallback(ctx, callback)
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error())
		return
	}
}

func (s *srv) handleTaskCallback(ctx context.Context, callback *tgbotapi.CallbackQuery) (bool, error) {
	if !strings.HasPrefix(callback.Data, string(models.NextButtonType)) &&
		!strings.HasPrefix(callback.Data, string(models.PrevButtonType)) {
		return false, nil
	}
	page := s.getPage(ctx, callback.Data)
	return true, s.refreshTaskMessage(ctx, callback.Message.Chat.ID, callback.Message.MessageID, page)
}

func (s *srv) getPage(ctx context.Context, data string) int {
	parts := strings.Split(data, ":")
	if len(parts) != 2 {
		s.logger.ErrorContext(ctx, "failed get page")
		return 0
	}
	pageStr := parts[1]
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error())
		return 0
	}
	return max(page, 0)
}
