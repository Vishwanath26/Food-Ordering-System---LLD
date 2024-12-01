package order

type Order struct {
	ID                   string
	RestaurantID         string
	ItemsQuantityMapping map[string]int
	Cost                 float32
}
