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
			Command:     string(models.TaskListCommand),
			Description: "Show task list",
		},
		{
			Command:     string(models.CreateTaskCommand),
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
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
		_, err := s.bot.Request(callback)
		if err != nil {
			s.logger.ErrorContext(ctx, err.Error())
		}
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
		err := s.list(ctx, message.From.ID)
		if err != nil {
			s.logger.ErrorContext(ctx, err.Error())
		}
	}
	if message.Text == string(models.CreateTaskCommand) {
		err := s.createTaskDruft(ctx, message.From.ID)
		if err != nil {
			s.logger.ErrorContext(ctx, err.Error())
		}
		return
	}
	err := s.manageTextMassage(ctx, *message)
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error())
	}
}

func (s *srv) handleCallback(ctx context.Context, callback *tgbotapi.CallbackQuery) {
	err := s.handleTaskCallback(ctx, callback)
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error())
		return
	}
}

func (s *srv) handleTaskCallback(ctx context.Context, callback *tgbotapi.CallbackQuery) error {
	data := callback.Data
	chatID := callback.Message.Chat.ID
	messageID := callback.Message.MessageID
	switch data {
	case string(models.CancelButtonType):
		err := s.deleteTaskDruft(ctx, chatID)
		if err != nil {
			return err
		}
		return s.deleteMessage(chatID, messageID)
	case string(models.NewStatus),
		string(models.InProgressStatus),
		string(models.DoneStatus),
		string(models.CanceledStatus):
		return s.addStatusToDruft(ctx, chatID, data, messageID)
	case string(models.CreateButtonType):
		return s.createFromDruft(ctx, chatID)
	}
	return s.checkCallbackPrefix(ctx, chatID, messageID, data)
}

func (s *srv) checkCallbackPrefix(ctx context.Context, chatID int64, messageID int, data string) error {
	handlers := []struct {
		prefixes []string
		handler  func(ctx context.Context, chatID int64, messageID int, data string) error
	}{
		{
			prefixes: []string{
				string(models.NextButtonType),
				string(models.PrevButtonType),
			},
			handler: func(ctx context.Context, chatID int64, messageID int, data string) error {
				page := s.getPage(ctx, data)
				return s.refreshTaskMessage(ctx, chatID, messageID, page)
			},
		},
		{
			prefixes: []string{string(models.YearButtonType)},
			handler:  s.selectMonth,
		},
		{
			prefixes: []string{string(models.MonthButtonType)},
			handler:  s.selectDay,
		},
		{
			prefixes: []string{
				string(models.NextYearButtonType),
				string(models.PrevYearButtonType),
			},
			handler: s.selectYear,
		},
		{
			prefixes: []string{string(models.DayButtonType)},
			handler:  s.selectHour,
		},
		{
			prefixes: []string{string(models.HourButtonType)},
			handler:  s.selectMinutes,
		},
		{
			prefixes: []string{string(models.MinutesButtonType)},
			handler:  s.addEstimateTimeToDruft,
		},
		{
			prefixes: []string{
				string(models.NextPageButtonType),
				string(models.PrevPageButtonType),
			},
			handler: s.updateMinutesMessage,
		},
	}
	for i := range handlers {
		for j := range handlers[i].prefixes {
			if strings.HasPrefix(data, handlers[i].prefixes[j]) {
				return handlers[i].handler(ctx, chatID, messageID, data)
			}
		}
	}
	return nil
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
