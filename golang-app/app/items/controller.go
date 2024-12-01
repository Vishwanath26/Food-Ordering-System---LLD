package items

import (
	"context"
	"test/app/registry"
)

func CreateNewItem(ctx context.Context, options ...Option) (*Item, error) {
	itemService, err := registry.GetServiceClient(ctx, "item_service")
	if err != nil {
		return nil, err
	}

	itemServiceClient := itemService.(*ItemServiceClient)

	item, err := itemServiceClient.create(options...)
	if err != nil {
		return nil, err
	}

	itemsMap := registry.GetEntityMap(ctx, "items").(map[string]*Item)
	itemsMap[item.ID] = item
	registry.UpdateEntityMap(ctx, "items", itemsMap)

	return item, nil
}
