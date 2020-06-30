package main

import (
	"log"
	"net/http"
	"net/smtp"
	"os"
	"shop/repository"
	"shop/tools/tgbot"
	"shop/utils/sendmail"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var (
	isDebug      *bool
	webAddr      string
	tgBotToken   string
	chatID       int64
	mailHost     string
	mailFrom     string
	mailPassword string
)

func init() {
	var ok bool
	webAddr, ok = os.LookupEnv("WEB_SERVER_ADDR")
	if !ok {
		log.Fatal("WEB_SERVER_ADDR env not set")
	}
	tgBotToken, ok = os.LookupEnv("TG_BOT_TOKEN")
	if !ok {
		log.Fatal("TG_BOT_TOKEN env not set")
	}
	tgChatID, ok := os.LookupEnv("TG_CHAT_ID")
	if !ok {
		log.Fatal("TG_CHAT_ID env not set")
	}
	var err error
	chatID, err = strconv.ParseInt(tgChatID, 10, 64)
	if err != nil {
		log.Fatal("Unable to parse chat ID")
	}

	mailHost, ok = os.LookupEnv("MAIL_HOST")
	if !ok {
		log.Fatal("MAIL_HOST env not set")
	}
	mailFrom, ok = os.LookupEnv("MAIL_FROM")
	if !ok {
		log.Fatal("MAIL_FROM env not set")
	}
	mailPassword, ok = os.LookupEnv("MAIL_PASSWORD")
	if !ok {
		log.Fatal("MAIL_FROM env not set")
	}

}

func main() {

	sendmail := sendmail.NewSendmail(mailFrom, mailHost, smtp.PlainAuth("", mailFrom, mailPassword, mailHost))

	bot, err := tgbot.NewShopTgBot(tgBotToken, chatID)
	if err != nil {
		log.Fatal("Unable to init tg bot")
	}

	handler := &shopHandler{
		bot:  bot,
		mail: sendmail,
	}
	if *isDebug {
		handler.db = repository.NewMapDB()
	}

	router := mux.NewRouter()

	router.HandleFunc("/item", handler.createItemHandler).Methods("POST")
	router.HandleFunc("/item/{id}", handler.getItemHandler).Methods("GET")
	router.HandleFunc("/item/{id}", handler.deleteItemHandler).Methods("DELETE")
	router.HandleFunc("/item/{id}", handler.updateItemHandler).Methods("PUT")

	router.HandleFunc("/order", handler.createOrderHandler).Methods("POST")
	router.HandleFunc("/order/{id}", handler.getOrderHandler).Methods("GET")

	srv := &http.Server{
		Addr:         webAddr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
