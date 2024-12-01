package items

import (
	"github.com/google/uuid"
)

type ItemServiceClient struct {
}

func (ic ItemServiceClient) ServiceName() string {
	return "item_service"
}

func NewItemServiceClient() *ItemServiceClient {
	return &ItemServiceClient{}
}

func (ic ItemServiceClient) create(options ...Option) (*Item, error) {
	newItem := Item{
		ID: uuid.New().String(),
	}

	for _, opt := range options {
		opt(&newItem)
	}
	return &newItem, nil
}
