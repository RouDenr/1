package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/heroku/x/hmetrics/onload"
)

const (
	webhook = "https://afternoon-wave-60151.herokuapp.com/"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.Run(":" + port)

	bot, err := tgbotapi.NewBotAPI("1114854329:AAGIXlmgZzvmdz16Av-w-wcd2-EVOjZnMSU")
	if err != nil {
		log.Fatal("cr bot: ", err)
	}
	log.Println("bot cr")

	if _, err := bot.SetWebhook(tgbotapi.NewWebhook(webhook)); err != nil {
		log.Fatalf("setting webhook %v; error: %v", webhook, err)
	}
	log.Println("webhook set")

	updates := bot.ListenForWebhook("/")
	for updates := range updates {
		if _, err := bot.Send(tgbotapi.NewMessage(updates.Message.Chat.ID, "эй")); err != nil {
			log.Print(err)
		}
	}
}
