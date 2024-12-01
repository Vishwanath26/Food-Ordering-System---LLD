package registry

import (
	"context"
	"errors"
)

type IService interface {
	ServiceName() string
}

func GetServiceClient(ctx context.Context, serviceName string) (IService, error) {
	services := ctx.Value("registry").(map[string]IService)

	if service, found := services[serviceName]; !found {
		return nil, errors.New("restaurant service not found")
	} else {
		return service, nil
	}
}

func GetEntityMap(ctx context.Context, entityName string) interface{} {
	datastore := ctx.Value("datastore").(map[string]interface{})
	return datastore[entityName]
}

func UpdateEntityMap(ctx context.Context, entityName string, updatedEntityMap interface{}) {
	datastore := ctx.Value("datastore").(map[string]interface{})
	datastore[entityName] = updatedEntityMap
}
