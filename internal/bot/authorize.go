package bot

import (
	"context"
	"log/slog"
	"strings"
)

func (s *srv) authorize(ctx context.Context, chatID int64, auth string) {
	parts := strings.Split(auth, " ")
	if len(parts) != 2 {
		s.logger.InfoContext(ctx, "incorrect command format")
		return
	}
	token := parts[1]
	res, err := s.userSrv.FindBotUser(ctx, token)
	if err != nil {
		s.logger.InfoContext(ctx, "error: ", slog.Any("error", err))
		s.sendSticker(ctx, chatID, s.sticker.FailedStickerFileID)
		return
	}
	err = s.userSrv.AddTelegramID(ctx, res.UserID, res.TelegramID)
	if err != nil {
		s.logger.WarnContext(ctx, "failed to authorize: ", slog.Any("error", err))
		s.sendSticker(ctx, chatID, s.sticker.FailedStickerFileID)
		return
	}
	s.sendSticker(ctx, chatID, s.sticker.SuccessStickerFileID)
}
