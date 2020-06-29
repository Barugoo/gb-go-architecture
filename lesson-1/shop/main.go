package main

import (
	"./repository"
	"flag"
	"log"
	"net/http"
	"os"
<<<<<<< HEAD
	"shop/repository"
	"shop/tools/tgbot"
	"strconv"
=======
>>>>>>> origin/master
	"time"

	"github.com/gorilla/mux"
)

func main() {
	isDebug := flag.Bool("mode", true, "runs app in debug mode")
	flag.Parse()

	webAddr, ok := os.LookupEnv("WEB_SERVER_ADDR")
	if !ok {
		log.Fatal("WEB_SERVER_ADDR env not set")
	}
	tgBotToken, ok := os.LookupEnv("TG_BOT_TOKEN")
	if !ok {
		log.Fatal("WEB_SERVER_ADDR env not set")
	}
	tgChatID, ok := os.LookupEnv("TG_CHAT_ID")
	if !ok {
		log.Fatal("WEB_SERVER_ADDR env not set")
	}
	chatID, err := strconv.ParseInt(tgChatID, 10, 64)
	if err != nil {
		log.Fatal("Unable to parse chat ID")
	}

	bot, err := tgbot.NewShopTgBot(tgBotToken, chatID)
	if err != nil {
		log.Fatal("Unable to init tg bot")
	}

	handler := &shopHandler{
		bot: bot,
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
