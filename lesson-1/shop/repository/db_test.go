package repository

import (
	"fmt"
	"gb/lesson-1/shop/models"
	"testing"
)

func TestCreateItem(t *testing.T) {

	db := NewMapDB()

	data := &models.Item{
		Name: "Product",
		Price: 100,
	}

	dataBad := &models.Item{
		Name: "Product",
		Price: -100,
	}

	res, err := db.CreateItem(data)
		if err != nil {
			t.Error(err)
		}

	resBad, err := db.CreateItem(dataBad)
		if err != nil {
			t.Error(err)
		}

		fmt.Printf("Expected price value (> 0 ): resulting value: %v", res)
		fmt.Printf("Expected price value (> 0 ): resulting value: %v", resBad)



}