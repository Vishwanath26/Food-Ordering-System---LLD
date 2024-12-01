package order

import (
	"context"
	"test/app/registry"
)

func CreateNewOrder(ctx context.Context) (*Order, error) {
	orderService, err := registry.GetServiceClient(ctx, "order_service")
	if err != nil {
		return nil, err
	}
	orderServiceClient := orderService.(*OrderServiceClient)

	order, err := orderServiceClient.create()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func AddItem(ctx context.Context, restaurantID string, itemID string, quantity int, order *Order) (*Order, error) {
	orderService, err := registry.GetServiceClient(ctx, "order_service")
	if err != nil {
		return nil, err
	}
	orderServiceClient := orderService.(*OrderServiceClient)

	order, err = orderServiceClient.addItemToOrder(ctx, restaurantID, itemID, quantity, order)
	if err != nil {
		return nil, err
	}
	return order, nil
}
