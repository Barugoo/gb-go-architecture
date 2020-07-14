package main

import (
	"./repository"
	"encoding/json"
<<<<<<< HEAD
	"fmt"
	"log"
	"net/http"
	"shop/models"
	"shop/repository"
	"shop/tools/tgbot"
=======
	"github.com/kaatinga/testModel"
	"log"
	"net/http"
>>>>>>> origin/master
	"strconv"

	"github.com/gorilla/mux"
)

type shopHandler struct {
	db  repository.Repository
	bot *tgbot.ShopTgBot
}

func (s *shopHandler) createOrderHandler(w http.ResponseWriter, r *http.Request) {
	order := new(models.Order)
	err := json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		return
	}

	for _, itemID := range order.ItemIDs {
		_, err := s.db.GetItem(int32(itemID))
		if err != nil {
			log.Println(err)
			json.NewEncoder(w).Encode(map[string]string{
				"status":  "not found",
				"message": fmt.Sprintf("item with ID %d not found", itemID),
			})
			return
		}
	}

	err = s.bot.SendOrderNotification(order)
	if err != nil {
		log.Println(err)
	}

	order, err = s.db.CreateOrder(order)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		return
	}
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func (s *shopHandler) getOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	orderID, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		return
	}

	order, err := s.db.GetOrder(int32(orderID))
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		return
	}

	err = json.NewEncoder(w).Encode(order)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		return
	}
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
