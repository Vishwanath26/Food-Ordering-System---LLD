package items

type Item struct {
	ID       string
	Name     string
	Category string
	Veg      bool
}

type Option func(item *Item)

func WithName(name string) Option {
	return func(item *Item) {
		item.Name = name
	}
}

func WithCategory(category string) Option {
	return func(item *Item) {
		item.Category = category
	}
}

func WithVeg(isVeg bool) Option {
	return func(item *Item) {
		item.Veg = isVeg
	}
}
