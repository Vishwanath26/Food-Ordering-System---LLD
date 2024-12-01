package restaurants

type Restaurant struct {
	Id      string
	Name    string
	Contact string
	Rating  float32
	MenuID  string
}

type Option func(r *Restaurant)

func WithName(name string) Option {
	return func(r *Restaurant) {
		r.Name = name
	}
}

func WithContact(contact string) Option {
	return func(r *Restaurant) {
		r.Contact = contact
	}
}

func WithRating(rating float32) Option {
	return func(r *Restaurant) {
		r.Rating = rating
	}
}
