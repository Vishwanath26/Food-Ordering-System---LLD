package restaurants

import "github.com/google/uuid"

type RestaurantServiceClient struct {
}

func (rs RestaurantServiceClient) ServiceName() string {
	return "restaurant_service"
}

func NewRestaurantServiceClient() *RestaurantServiceClient {
	return &RestaurantServiceClient{}
}

func (rs RestaurantServiceClient) AddNewRestaurant(opts ...Option) (*Restaurant, error) {
	restaurant := &Restaurant{
		Id: uuid.New().String(),
	}

	for _, opt := range opts {
		opt(restaurant)
	}

	return restaurant, nil
}
