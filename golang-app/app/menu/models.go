package menu

type Menu struct {
	ID                string
	RestaurantID      string
	ItemsPriceMapping map[string]float32
}

type Option func(menu *Menu)

func WithRestaurantID(restaurantID string) Option {
	return func(menu *Menu) {
		menu.RestaurantID = restaurantID
	}
}

func WithItemsPriceMapping(itemsPriceMapping map[string]float32) Option {
	return func(menu *Menu) {
		menu.ItemsPriceMapping = itemsPriceMapping
	}
}
