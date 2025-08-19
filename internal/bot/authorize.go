package bot

import (
	"context"
	"strings"
)

func (s *srv) authorize(ctx context.Context, chatID int64, auth string) error {
	parts := strings.Split(auth, " ")
	if len(parts) != 2 {
		s.logger.ErrorContext(ctx, "incorrect command format")
		return s.sendSticker(chatID, s.sticker.Confused)
	}
	token := parts[1]
	res, err := s.userSrv.FindBotUser(ctx, token)
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error())
		return s.sendSticker(chatID, s.sticker.Confused)
	}
	err = s.userSrv.AddTelegramID(ctx, res.UserID, chatID)
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error())
		return s.sendSticker(chatID, s.sticker.Confused)
	}
	return s.sendSticker(chatID, s.sticker.Congratulations)
}
