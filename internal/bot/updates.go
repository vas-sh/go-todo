package bot

import (
	"context"
	"log"
	"log/slog"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vas-sh/todo/internal/models"
)

func (s *srv) Updates() {
	s.createMenu()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := s.bot.GetUpdatesChan(u)
	for update := range updates {
		go s.manageUpdate(update)
	}
}

func (s *srv) createMenu() {
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
		log.Println("failed create menu: ", err)
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
		s.authorize(ctx, message.From.ID, message.Text)
		return
	}
	if message.Text == string(models.TaskListCommand) {
		s.sendTaskMessage(ctx, message.From.ID)
		return
	}
}

func (s *srv) handleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery) {
	if s.handleTaskCallback(ctx, callback) {
		return
	}
}

func (s *srv) handleTaskCallback(ctx context.Context, callback *tgbotapi.CallbackQuery) bool {
	if !strings.HasPrefix(callback.Data, string(models.NextButtonType)) &&
		!strings.HasPrefix(callback.Data, string(models.PrevButtonType)) {
		return false
	}
	page := s.getPage(ctx, callback.Data)
	s.refreshTaskMessage(ctx, callback.Message.Chat.ID, callback.Message.MessageID, page)
	return true
}

func (s *srv) getPage(ctx context.Context, data string) int {
	parts := strings.Split(data, ":")
	if len(parts) != 2 {
		s.logger.WarnContext(ctx, "failed get page")
		return 0
	}
	pageStr := parts[1]
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		s.logger.WarnContext(ctx, "failed convert page: ", slog.Any("error", err))
		return 0
	}
	return max(page, 0)
}
