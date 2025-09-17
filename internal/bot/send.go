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

func (s *srv) sendTextMassageWithKeyboard(chatID int64, text string, keyboard tgbotapi.InlineKeyboardMarkup) error {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = keyboard
	_, err := s.bot.Send(msg)
	return err
}

func (s *srv) refreshTextMassageWithKeyboard(
	chatID int64, keyboard tgbotapi.InlineKeyboardMarkup, messageID int, text string,
) error {
	msg := tgbotapi.NewEditMessageText(chatID, messageID, text)
	msg.ReplyMarkup = &keyboard
	_, err := s.bot.Send(msg)
	return err
}

func (s *srv) refreshMassage(chatID int64, messageID int, text string) error {
	msg := tgbotapi.NewEditMessageText(chatID, messageID, text)
	_, err := s.bot.Send(msg)
	return err
}

func (s *srv) deleteMessage(chatID int64, messageID int) error {
	deleteMsg := tgbotapi.NewDeleteMessage(chatID, messageID)
	_, err := s.bot.Request(deleteMsg)
	return err
}
