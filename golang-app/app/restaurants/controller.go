package restaurants

import (
	"context"
	"test/app/registry"
)

func AddRestaurant(ctx context.Context, opts ...Option) (*Restaurant, error) {
	restaurantService, err := registry.GetServiceClient(ctx, "restaurant_service")
	if err != nil {
		return nil, err
	}

	restaurantServiceClient := restaurantService.(*RestaurantServiceClient)

	restaurant, err := restaurantServiceClient.AddNewRestaurant(opts...)
	if err != nil {
		return nil, err
	}

	restaurantMap := registry.GetEntityMap(ctx, "restaurants").(map[string]*Restaurant)
	restaurantMap[restaurant.Id] = restaurant
	registry.UpdateEntityMap(ctx, "restaurants", restaurantMap)

	return restaurant, nil
}
