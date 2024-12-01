package order

import (
	"context"
	"github.com/google/uuid"
	"test/app/menu"
	"test/app/registry"
	"test/app/restaurants"
)

type OrderServiceClient struct {
}

func (oc *OrderServiceClient) ServiceName() string {
	return "order_service"
}

func NewOrderServiceClient() *OrderServiceClient {
	return &OrderServiceClient{}
}

func (oc *OrderServiceClient) create() (*Order, error) {
	order := &Order{
		ID: uuid.New().String(),
	}

	return order, nil
}

func (oc *OrderServiceClient) addItemToOrder(ctx context.Context, restaurantID string, itemID string, quantity int, order *Order) (*Order, error) {
	restaurantMapping := registry.GetEntityMap(ctx, "restaurants").(map[string]*restaurants.Restaurant)
	r := restaurantMapping[restaurantID]

	menuMapping := registry.GetEntityMap(ctx, "menu").(map[string]*menu.Menu)
	m := menuMapping[r.MenuID]
	price := m.ItemsPriceMapping[itemID]

	order.Cost = order.Cost + (float32(quantity) * price)
	order.ItemsQuantityMapping[itemID] = quantity

	orderMapping := registry.GetEntityMap(ctx, "order").(map[string]*Order)
	orderMapping[order.ID] = order

	registry.UpdateEntityMap(ctx, "order", orderMapping)
	return order, nil
}
