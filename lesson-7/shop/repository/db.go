package repository

import (
	"context"
	"fmt"

	"shop/models"

	"go.uber.org/zap"
)

type Repository interface {
	CreateItem(ctx context.Context, item *models.Item) (*models.Item, error)
	GetItem(ctx context.Context, ID int32) (*models.Item, error)
	DeleteItem(ctx context.Context, ID int32) error
	UpdateItem(ctx context.Context, item *models.Item) (*models.Item, error)

	CreateOrder(ctx context.Context, Order *models.Order) (*models.Order, error)
	GetOrder(ctx context.Context, ID int32) (*models.Order, error)
}

type mapDB struct {
	logger zap.Logger

	itemsTable  *itemsTable
	ordersTable *ordersTable
}

type itemsTable struct {
	items map[int32]*models.Item
	maxID int32
}

type ordersTable struct {
	orders map[int32]*models.Order
	maxID  int32
}

func NewMapDB() Repository {
	return &mapDB{
		itemsTable: &itemsTable{
			items: make(map[int32]*models.Item),
			maxID: 0,
		},
		ordersTable: &ordersTable{
			orders: make(map[int32]*models.Order),
			maxID:  0,
		},
	}
}

func (m *mapDB) CreateItem(ctx context.Context, item *models.Item) (*models.Item, error) {
	m.itemsTable.maxID++

	newItem := &models.Item{
		ID:    m.itemsTable.maxID,
		Price: item.Price,
		Name:  item.Name,
	}

	m.itemsTable.items[newItem.ID] = newItem

	return &models.Item{
		ID:    newItem.ID,
		Name:  newItem.Name,
		Price: newItem.Price,
	}, nil
}

func (m *mapDB) GetItem(ctx context.Context, ID int32) (*models.Item, error) {
	item, ok := m.itemsTable.items[ID]
	if !ok {
		return nil, fmt.Errorf("Item with ID: %d is not found", ID)
	}

	return &models.Item{
		ID:    item.ID,
		Name:  item.Name,
		Price: item.Price,
	}, nil
}

func (m *mapDB) DeleteItem(ctx context.Context, ID int32) error {
	delete(m.itemsTable.items, ID)
	return nil
}

func (m *mapDB) UpdateItem(ctx context.Context, item *models.Item) (*models.Item, error) {
	updateItem, ok := m.itemsTable.items[item.ID]
	if !ok {
		return nil, fmt.Errorf("Item with ID: %d is not found", item.ID)
	}
	updateItem.Name = item.Name
	updateItem.Price = item.Price

	return &models.Item{
		ID:    updateItem.ID,
		Name:  updateItem.Name,
		Price: updateItem.Price,
	}, nil
}

func (m *mapDB) CreateOrder(ctx context.Context, order *models.Order) (*models.Order, error) {
	m.ordersTable.maxID++

	newOrder := &models.Order{
		ID:      m.ordersTable.maxID,
		Phone:   order.Phone,
		Email:   order.Email,
		ItemIDs: order.ItemIDs,
	}

	m.ordersTable.orders[newOrder.ID] = newOrder

	return &models.Order{
		ID:      newOrder.ID,
		Phone:   newOrder.Phone,
		Email:   newOrder.Email,
		ItemIDs: newOrder.ItemIDs,
	}, nil
}

func (m *mapDB) GetOrder(ctx context.Context, ID int32) (*models.Order, error) {
	order, ok := m.ordersTable.orders[ID]
	if !ok {
		return nil, fmt.Errorf("Order with ID: %d is not found", ID)
	}

	return &models.Order{
		ID:      order.ID,
		Phone:   order.Phone,
		ItemIDs: order.ItemIDs,
	}, nil
}
