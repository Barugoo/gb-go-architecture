package repository

import (
	"fmt"
	"shop/models"
	"testing"
)

var example = &models.Item{
	ID:    int32(194),
	Name:  "example",
	Price: 1098,
}

func TestDBCreateItem(t *testing.T) {
	mapDB := mapDB{
		db:    make(map[int32]*models.Item, 1),
		maxID: 0,
	}
	newItem := &models.Item{
		Name:  "test",
		Price: 1,
	}
	createdItem, err := mapDB.CreateItem(newItem)
	if err != nil {
		t.Error("some expected create error")
	}
	if createdItem.Name != newItem.Name {
		t.Errorf("exp name == %s, got %s", newItem.Name, createdItem.Name)
	}
	if createdItem.Price != createdItem.Price {
		t.Errorf("exp price == %d, got %d", newItem.Price, createdItem.Price)
	}

	if createdItem == nil {
		t.Error("createdItem is nil")
	}
}

func TestDBGetItem(t *testing.T) {
	mapDB := mapDB{
		db:    make(map[int32]*models.Item, 1),
		maxID: 0,
	}
	existedItem, err := mapDB.CreateItem(example)
	if err != nil {
		t.Error("unexpected Err filling test map")
	}
	gotItem, err := mapDB.GetItem(existedItem.ID)
	if gotItem == nil {
		if err == fmt.Errorf("Item with ID: %d is not found", gotItem.ID) {
			t.Error("Item with ID is not found")
			return
		}
		t.Errorf("cant get item with id: %d", existedItem.ID)
	}
	if gotItem.Name != existedItem.Name {
		t.Errorf("Name error. expected =  %s, got = %s", gotItem.Name, existedItem.Name)
	}

	if gotItem.Price != existedItem.Price {
		t.Errorf("Price error. expected = %d, have = %d", gotItem.Price, existedItem.Price)
	}
}

func TestDBUpdateItem(t *testing.T) {
	mapDB := mapDB{
		db:    make(map[int32]*models.Item, 1),
		maxID: 0,
	}
	existedItem, err := mapDB.CreateItem(example)
	if err != nil {
		t.Error("unexpected Err filling test map")
	}

	newName := "newName"
	newPrice := int32(100500)

	var newDataItem = &models.Item{ //new data for existing item in DB
		ID:    existedItem.ID,
		Name:  newName,
		Price: newPrice,
	}
	newItem, err := mapDB.UpdateItem(newDataItem)
	if newItem == nil {
		if err == fmt.Errorf("Item with ID: %d is not found", existedItem.ID) {
			t.Errorf("Cant find existing item with ID: %d", existedItem.ID)
			return
		}
		t.Error("Updating error")
	}

	if newItem.Name != newName {
		t.Errorf("expected name = %s, got %s", newItem.Name, newName)
	}

	if newItem.Price != newPrice {
		t.Errorf("expected name = %d, got %d", newItem.Price, newPrice)
	}
}

func TestMapDBDeleteItem(t *testing.T) {
	mapDB := mapDB{
		db:    make(map[int32]*models.Item, 1),
		maxID: 0,
	}

	existedItem, err := mapDB.CreateItem(example)
	if err != nil {
		t.Error("unexpected Err filling test map")
	}

	err = mapDB.DeleteItem(existedItem.ID)
	if err != nil {
		t.Error("Del error")
	}
	_, err = mapDB.GetItem(existedItem.ID)
	if err == nil {
		t.Errorf("Item with ID: %d is not deleted", existedItem.ID)
	}
}
