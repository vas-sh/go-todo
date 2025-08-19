package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *srv) sendTextMessage(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := s.bot.Send(msg)
	return err
}

func (s *srv) sendSticker(chatID int64, sticker string) error {
	msg := tgbotapi.NewSticker(chatID, tgbotapi.FileID(sticker))
	_, err := s.bot.Send(msg)
	return err
}
