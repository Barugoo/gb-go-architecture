package repository

import (
	"gb-go-architecture/lesson-1/shop_new/models"
	"testing"
)

func TestMapDB_CreateItem(t *testing.T) {
	db := NewMapDB()

	input := &models.Item{
		Name:  "someName",
		Price: 10,
	}
	expected := &models.Item{
		ID:    1,
		Name:  input.Name,
		Price: input.Price,
	}

	result, err := db.CreateItem(input)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	testItems(t, expected, result)

	result, err = db.GetItem(expected.ID)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	testItems(t, expected, result)

	input = &models.Item{
		Name:  "someName2",
		Price: 20,
	}
	expected = &models.Item{
		ID:    2,
		Name:  input.Name,
		Price: input.Price,
	}

	result, err = db.CreateItem(input)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	testItems(t, expected, result)
}

func TestMapDB_UpdateItem(t *testing.T) {
	db := NewMapDB()
	initialData := &models.Item{
		Name:  "someName",
		Price: 10,
	}

	item, err := db.CreateItem(initialData)
	if err != nil {
		t.Error("unexpected error: ", err)
	}
	expectedData := &models.Item{
		ID:    item.ID,
		Name:  "someName 2",
		Price: 20,
	}
	updatedItem, err := db.UpdateItem(expectedData)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	testItems(t, expectedData, updatedItem)

}

func TestMapDB_ListItem(t *testing.T) {
	db := NewMapDB()
	itemsToCreate := []models.Item{
		{
			ID:    1,
			Name:  "Name 10",
			Price: 10,
		},
		{
			ID:    2,
			Name:  "Name 20",
			Price: 20,
		},
		{
			ID:    3,
			Name:  "Name 30",
			Price: 30,
		},
		{
			ID:    4,
			Name:  "Name 40",
			Price: 40,
		},
	}

	type testCases struct {
		Description   string
		Filter        *ItemFilter
		ExpectedItems []models.Item
	}

	tests := []testCases{
		{
			Description: "Count all to test setup",
			Filter: &ItemFilter{
				Limit:  1000,
				Offset: 0,
			},
			ExpectedItems: itemsToCreate,
		},
		{
			Description: "When offset is passed",
			Filter: &ItemFilter{
				Limit:  1000,
				Offset: 2,
			},
			ExpectedItems: itemsToCreate[2:],
		},
		{
			Description: "When limit is passed",
			Filter: &ItemFilter{
				Limit:  3,
				Offset: 0,
			},
			ExpectedItems: itemsToCreate[:3],
		},
		{
			Description: "When offset and limit are passed",
			Filter: &ItemFilter{
				Limit:  4,
				Offset: 1,
			},
			ExpectedItems: itemsToCreate[1:4],
		},
		{
			Description: "When PriceLeft is passed",
			Filter: &ItemFilter{
				PriceLeft: createInt64(20),
				Limit:     1000,
				Offset:    0,
			},
			ExpectedItems: itemsToCreate[1:],
		},
		{
			Description: "When PriceLeft and limit are passed",
			Filter: &ItemFilter{
				PriceLeft: createInt64(20),
				Limit:     1,
				Offset:    0,
			},
			ExpectedItems: itemsToCreate[1:2],
		},
		{
			Description: "When PriceLeft and Limit and Offset are passed",
			Filter: &ItemFilter{
				PriceLeft: createInt64(30),
				Limit:     1,
				Offset:    1,
			},
			ExpectedItems: itemsToCreate[3:4],
		},
		{
			Description: "When PriceRight and PriceLeft are passed",
			Filter: &ItemFilter{
				PriceRight: createInt64(40),
				PriceLeft:  createInt64(20),
				Limit:      1000,
				Offset:     0,
			},
			ExpectedItems: itemsToCreate[1:4],
		},
		{
			Description: "When PriceRight is passed",
			Filter: &ItemFilter{
				PriceRight: createInt64(30),
				Limit:      1000,
				Offset:     0,
			},
			ExpectedItems: itemsToCreate[:3],
		},
	}

	for _, item := range itemsToCreate {
		_, err := db.CreateItem(&item)
		if err != nil {
			t.Error("unexpected error: ", err)
		}
	}

	for _, testCase := range tests {
		t.Log(testCase.Description)
		items, err := db.ListItems(testCase.Filter)
		if err != nil {
			t.Error("can't return the list of elements: ", err)
		}
		if len(items) != len(testCase.ExpectedItems) {
			t.Errorf("wrong ammount, expected %d, got %d", len(testCase.ExpectedItems), len(items))
		}

		for i, item := range items {
			testItems(t, &testCase.ExpectedItems[i], item)
		}
	}
}

func TestMapDB_DeleteItem(t *testing.T) {
	db := NewMapDB()
	initialData := &models.Item{
		Name:  "someName",
		Price: 10,
	}
	filter := &ItemFilter{
		Limit:  100,
		Offset: 0,
	}

	item, err := db.CreateItem(initialData)
	if err != nil {
		t.Error("unexpected error: ", err)
	}
	items, _ := db.ListItems(filter)
	if len(items) != 1 {
		t.Errorf("wrong test setup, expected 1, got %d", len(items))
	}

	err = db.DeleteItem(item.ID)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	items, _ = db.ListItems(filter)
	if len(items) != 0 {
		t.Error("Item was not deleted")
	}
}

func testItems(t *testing.T, expectedItem, item *models.Item) {
	if expectedItem.ID != item.ID {
		t.Errorf("unexpected id: expected %d result: %d", expectedItem.ID, item.ID)
	}
	if expectedItem.Name != item.Name {
		t.Errorf("unexpected name: expected %s result: %s", expectedItem.Name, item.Name)
	}
	if expectedItem.Price != item.Price {
		t.Errorf("unexpected price: expected %d result: %d", expectedItem.Price, item.Price)
	}
}

func createInt64(x int64) *int64 {
	return &x
}
