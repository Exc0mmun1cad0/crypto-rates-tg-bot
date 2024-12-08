package main

import (
	"log"
	"os"

	"github.com/telegram_bot/bot/config"
	eventconsumer "github.com/telegram_bot/bot/consumer/event_consumer"
	telegram "github.com/telegram_bot/bot/events/tg_event_processor"
	e "github.com/telegram_bot/bot/lib/error_wrapping"
	sqlite3 "github.com/telegram_bot/bot/storage/sqlite"
	client "github.com/telegram_bot/bot/tg_client"
)

const (
	configPath = "configs/bot_config.yaml"
	tokenPath  = "configs/token.txt"
)

func main() {
	cfg, err := config.NewBotConfig(configPath)
	if err != nil {
		log.Fatal("cannot read bot config:", err)
	}

	token, err := mustToken(tokenPath)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlite3.New("db/tokens.db")
	if err != nil {
		log.Fatal("cannot connect to storage:", err)
	}

	if err := db.Init(); err != nil {
		log.Fatal("cannot init storage:", err)
	}

	eventProcessor := telegram.New(
		client.New(cfg.Host, token),
		db,
	)

	log.Print("bot is started")

	consumer := eventconsumer.New(eventProcessor, eventProcessor, cfg.BatchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("bot is stopped", err)
	}
}

// Get token from .txt file
func mustToken(tokenPath string) (string, error) {
	token, err := os.ReadFile(tokenPath)
	if err != nil {
		return "", e.Wrap("cannot read token", err)
	}

	return string(token), nil
}
