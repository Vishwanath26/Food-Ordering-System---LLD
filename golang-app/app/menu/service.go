package menu

import (
	"context"
	"github.com/google/uuid"
	"test/app/items"
)

type MenuServiceClient struct {
}

func (mc MenuServiceClient) ServiceName() string {
	return "menu_service"
}

func NewMenuServiceClient() MenuServiceClient {
	return MenuServiceClient{}
}

func (mc MenuServiceClient) create(ctx context.Context, options ...Option) (*Menu, error) {
	menu := Menu{
		ID: uuid.New().String(),
	}

	for _, opt := range options {
		opt(&menu)
	}
	return &menu, nil
}

func (mc MenuServiceClient) AddItem(ctx context.Context, menuID string, item items.Item, price float32) (*Menu, error) {
	dataStore := ctx.Value("datastore").(map[string]interface{})
	menuList := dataStore["menu"].(map[string]*Menu)
	menu := menuList[menuID]

	menuItemsMapping := menu.ItemsPriceMapping
	menuItemsMapping[item.ID] = price
	menu.ItemsPriceMapping = menuItemsMapping

	return menu, nil
}
