package main

import (
	"fmt"
	"log"
	"os"

	"github.com/incrypto32/airdropbot/handler"
	"github.com/incrypto32/airdropbot/router"
	"github.com/incrypto32/airdropbot/tgbotgo"
	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.Llongfile)

	// Load ENV FIle
	if err := godotenv.Load(); err != nil {
		log.Panic("Failed loading env variables", err)
	}

	port := os.Getenv("PORT")

	// echo instance
	r := router.New()

	// tgbot
	bot, err := tgbotgo.New(os.Getenv("BOT_TOKEN"), "https://api.telegram.org")
	if err != nil {
		log.Panic(err)
	}

	handler.New(r, bot)
	log.Println(fmt.Sprintf("server started at http://localhost:%s", port))
	r.Logger.Fatal(r.Start(":" + port))
}
