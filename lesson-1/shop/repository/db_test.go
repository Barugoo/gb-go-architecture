package repository

import (
	"reflect"
	"shop/models"
	"testing"
)

func Test_Repository(t *testing.T) {
	wanted := map[int32]*models.Item{
		1: {1, "test1", 10},
		2: {2, "test2", 20},
		3: {3, "test3", 30},
	}

	t.Run("Test CreateItem", func(t *testing.T) {
		store := mapDB{
			db: make(map[int32]*models.Item),
		}

		store.CreateItem(&models.Item{1, "test1", 10})
		store.CreateItem(&models.Item{2, "test2", 20})
		store.CreateItem(&models.Item{3, "test3", 30})

		if len(store.db) != len(wanted) {
			t.Fatalf("got %v want %v", store.db, wanted)
		}

		for k, g := range store.db {
			w, ok := wanted[k]
			if !ok {
				t.Fatalf("key %v not exists", k)
			}

			if !reflect.DeepEqual(g, w) {
				t.Fatalf("got %v want %v", g, w)
			}
		}
	})

	t.Run("Test getItem", func(t *testing.T) {
		store := mapDB{
			db: map[int32]*models.Item{
				1: {1, "test1", 10},
				2: {2, "test2", 20},
				3: {3, "test3", 30},
			},
		}

		for k, w := range wanted {
			got, err := store.GetItem(k)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(got, w) {
				t.Fatalf("got %v want  %v", got, w)
			}
		}

	})

	t.Run("Test deleteItem", func(t *testing.T) {
		store := mapDB{
			db: map[int32]*models.Item{
				1: {1, "test1", 10},
				2: {2, "test2", 20},
				3: {3, "test3", 30},
			},
		}

		store.DeleteItem(1)
		store.DeleteItem(2)
		store.DeleteItem(3)

		if len(store.db) != 0 {
			t.Fatalf("got %v want %v", len(store.db), 0)
		}
	})

	t.Run("Test updateItem", func(t *testing.T) {
		wanted := map[int32]*models.Item{
			1: {1, "update", 100},
			2: {2, "update", 100},
			3: {3, "update", 100},
		}
		store := mapDB{
			db: map[int32]*models.Item{
				1: {1, "need update", 1},
				2: {2, "need update", 1},
				3: {3, "need update", 1},
			},
		}
		for _, v := range wanted {
			got, err := store.UpdateItem(v)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, v) {
				t.Fatalf("got %v want %v", got, v)
			}
		}

	})
}
