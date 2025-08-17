package bot

import (
	"context"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *srv) sendTextMessage(ctx context.Context, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := s.bot.Send(msg)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed send message: ", slog.Any("error", err))
	}
}

func (s *srv) sendSticker(ctx context.Context, chatID int64, sticker string) {
	msg := tgbotapi.NewSticker(chatID, tgbotapi.FileID(sticker))
	_, err := s.bot.Send(msg)
	if err != nil {
		s.logger.ErrorContext(ctx, "failed send sticker: ", slog.Any("error", err))
	}
}
