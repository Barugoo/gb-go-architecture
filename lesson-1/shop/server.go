package main

import (
	"encoding/json"
	"log"
	"net/http"
	"shop/models"
	"shop/repository"
	"strconv"

	"github.com/gorilla/mux"
)

type shopHandler struct {
	db repository.Repository
}

func (s *shopHandler) createItemHandler(w http.ResponseWriter, r *http.Request) {
	item := new(models.Item)
	err := json.NewDecoder(r.Body).Decode(item)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		return
	}

	item, err = s.db.CreateItem(item)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		return
	}
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func (s *shopHandler) getItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	itemID, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		return
	}

	item, err := s.db.GetItem(int32(itemID))
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		return
	}

	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		return
	}
}

func (s *shopHandler) deleteItemHandler(w http.ResponseWriter, r *http.Request) {

}
