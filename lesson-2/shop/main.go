package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"shop/pkg/sendmail"
	"shop/pkg/tgbot"
	"shop/repository"
	"shop/service"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	tg, err := tgbot.NewTelegramAPI("1538307948:AAEJSbHwgf2AVz0bdCWzKH8cC6cT56lDBaA", 201882409)
	if err != nil {
		log.Fatal("Unable to init telegram bot")
	}

	sm := sendmail.NewSentMail(os.Getenv("FROM_EMAIL"), os.Getenv("HOST_EMAIL"), os.Getenv("PASSWORD_EMAIL"))

	db := repository.NewMapDB()

	service := service.NewService(tg, db, sm)
	handler := &shopHandler{
		service: service,
		db:      db,
	}

	router := mux.NewRouter()

	router.HandleFunc("/item", handler.createItemHandler).Methods("POST")
	router.HandleFunc("/item/{id}", handler.getItemHandler).Methods("GET")
	router.HandleFunc("/item/{id}", handler.deleteItemHandler).Methods("DELETE")
	router.HandleFunc("/item/{id}", handler.updateItemHandler).Methods("PUT")

	router.HandleFunc("/order", handler.createOrderHandler).Methods("POST")
	router.HandleFunc("/order/{id}", handler.getOrderHandler).Methods("GET")

	srv := &http.Server{
		Addr:         ":8080",
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
