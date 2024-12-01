package menu

import (
	"context"
	"test/app/items"
	"test/app/registry"
)

func AddNewItem(ctx context.Context, menuID string, item items.Item, price float32) (*Menu, error) {
	menuService, err := registry.GetServiceClient(ctx, "menu_service")
	if err != nil {
		return nil, err
	}

	menuServiceClient := menuService.(MenuServiceClient)

	updatedMenu, err := menuServiceClient.AddItem(ctx, menuID, item, price)
	if err != nil {
		return nil, err
	}

	return updatedMenu, nil
}

func CreateMenu(ctx context.Context, options ...Option) (*Menu, error) {
	menuService, err := registry.GetServiceClient(ctx, "menu_service")
	if err != nil {
		return nil, err
	}

	menuServiceClient := menuService.(MenuServiceClient)

	menu, err := menuServiceClient.create(ctx, options...)
	if err != nil {
		return nil, err
	}

	return menu, nil
}
