package main

import (
	"./models"
	"./repository"
	"encoding/json"
	"log"
	"net/http"
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
		err = json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		if err != nil {
			log.Println(err)
		}
		return
	}

	item, err = s.db.CreateItem(item)
	if err != nil {
		log.Println(err)
		err = json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		if err != nil {
			log.Println(err)
		}
		return
	}
	err = json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	if err != nil {
		log.Println(err)
	}
}

func (s *shopHandler) getItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	itemID, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		err = json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		if err != nil {
			log.Println(err)
		}
		return
	}

	item, err := s.db.GetItem(int32(itemID))
	if err != nil {
		log.Println(err)
		err = json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		log.Println(err)
		err = json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		if err != nil {
			log.Println(err)
		}
		return
	}
}

func (s *shopHandler) deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	itemID, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		err = json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = s.db.DeleteItem(int32(itemID))
	if err != nil {
		log.Println(err)
		err = json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		if err != nil {
			log.Println(err)
		}
		return
	}
	err = json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	if err != nil {
		log.Println(err)
	}
}

func (s *shopHandler) updateItemHandler(w http.ResponseWriter, r *http.Request) {
	updatedItem := new(models.Item)
	err := json.NewDecoder(r.Body).Decode(updatedItem)
	if err != nil {
		log.Println(err)
		err = json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		if err != nil {
			log.Println(err)
		}
		return
	}

	vars := mux.Vars(r)
	itemIDStr := vars["id"]

	var itemID int
	itemID, err = strconv.Atoi(itemIDStr)
	if err != nil {
		log.Println(err)
		err = json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		if err != nil {
			log.Println(err)
		}
		return
	}

	var item *models.Item
	item, err = s.db.GetItem(int32(itemID))
	if err != nil {
		log.Println(err)
		err = json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		if err != nil {
			log.Println(err)
		}
		return
	}

	item.Name = updatedItem.Name
	item.Price = updatedItem.Price

	item, err = s.db.UpdateItem(item)
	if err != nil {
		log.Println(err)
		err = json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		log.Println(err)
		err = json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		if err != nil {
			log.Println(err)
		}
		return
	}
}
