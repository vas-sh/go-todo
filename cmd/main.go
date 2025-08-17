package main

import (
	"log"
	"log/slog"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
	"github.com/vas-sh/todo/internal/bot"
	"github.com/vas-sh/todo/internal/config"
	"github.com/vas-sh/todo/internal/db"
	"github.com/vas-sh/todo/internal/handlers"
	"github.com/vas-sh/todo/internal/handlers/taskhandlers"
	"github.com/vas-sh/todo/internal/handlers/userhandlers"
	"github.com/vas-sh/todo/internal/handlers/wshandlers"
	"github.com/vas-sh/todo/internal/mail"
	"github.com/vas-sh/todo/internal/repo/taskrepo"
	"github.com/vas-sh/todo/internal/repo/userrepo"
	"github.com/vas-sh/todo/internal/services/jwttoken"
	"github.com/vas-sh/todo/internal/services/task"
	"github.com/vas-sh/todo/internal/services/user"
)

func main() {
	cfg := config.Config
	mailSrv, err := mail.New(cfg.MailLogin, cfg.MailPassword, cfg.MailHost, cfg.MailPort)
	if err != nil {
		panic(err)
	}
	databace, err := db.New(cfg.DB)
	if err != nil {
		panic(err)
	}
	sqlDB, err := databace.DB()
	if err != nil {
		panic(err)
	}
	defer func() {
		err := sqlDB.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	databace = databace.Debug()
	err = db.Migrate(databace)
	if err != nil {
		panic(err)
	}

	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	logger := slog.New(h)

	taskRepo := taskrepo.New(databace)
	taskSrv := task.New(taskRepo)

	userRepo := userrepo.New(databace)
	userSrv := user.New(userRepo, mailSrv)

	userFetcher := jwttoken.New(cfg.SecretJWT)
	server := handlers.New(userFetcher)
	anonRouter := server.AnonRouter()
	authRouter := server.AuthRouter()
	wsSrv := wshandlers.New(userFetcher)
	wsSrv.Register(anonRouter)

	tgBotAPI, err := tgbotapi.NewBotAPI(cfg.TgBotToken)
	if err != nil {
		panic(err)
	}
	tgBot := bot.New(userSrv, taskSrv, tgBotAPI, logger)
	go tgBot.Updates()

	taskhandlers.New(taskSrv, wsSrv).Register(authRouter)
	userhandlers.New(userSrv, cfg.SecretJWT, userFetcher).Register(anonRouter, authRouter)

	err = server.Run()
	if err != nil {
		panic(err)
	}
}
