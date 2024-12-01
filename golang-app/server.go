package main

import (
	"context"
	"fmt"
	"log"
	"test/app/items"
	"test/app/menu"
	"test/app/order"
	"test/app/registry"
	"test/app/restaurants"
)

func bootServices() map[string]registry.IService {
	servicesRegistry := map[string]registry.IService{}

	restaurantService := restaurants.NewRestaurantServiceClient()
	itemService := items.NewItemServiceClient()
	menuService := menu.NewMenuServiceClient()
	orderService := order.NewOrderServiceClient()

	servicesRegistry[restaurantService.ServiceName()] = restaurantService
	servicesRegistry[itemService.ServiceName()] = itemService
	servicesRegistry[menuService.ServiceName()] = menuService
	servicesRegistry[orderService.ServiceName()] = orderService

	return servicesRegistry
}

func bootDataStore() map[string]interface{} {
	dataStore := map[string]interface{}{
		"restaurants": map[string]*restaurants.Restaurant{},
		"items":       map[string]*items.Item{},
		"menu":        map[string]*menu.Menu{},
		"order":       map[string]*order.Order{},
	}
	return dataStore
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "registry", bootServices())
	ctx = context.WithValue(ctx, "datastore", bootDataStore())

	// Create first Restaurant
	restaurantOptions := []restaurants.Option{
		restaurants.WithName("restaurant1"),
		restaurants.WithContact("contact1"),
		restaurants.WithRating(5.0),
	}

	restaurant1, err := restaurants.AddRestaurant(ctx, restaurantOptions...)
	if err != nil {
		log.Default().Print(err.Error())
	} else {
		fmt.Println("New restaurant added ", restaurant1)
	}

	// Create second Restaurant
	restaurantOptions = []restaurants.Option{
		restaurants.WithName("restaurant2"),
		restaurants.WithContact("contact2"),
		restaurants.WithRating(4.5),
	}

	restaurant2, err := restaurants.AddRestaurant(ctx, restaurantOptions...)
	if err != nil {
		log.Default().Print(err.Error())
	} else {
		fmt.Println("New restaurant added ", restaurant2)
	}

	// Create first item
	itemOptions := []items.Option{
		items.WithName("roti"),
		items.WithCategory("bread"),
		items.WithVeg(true),
	}
	item1, err := items.CreateNewItem(ctx, itemOptions...)
	if err != nil {
		log.Default().Print(err.Error())
	} else {
		fmt.Println("New item added ", *item1)
	}

	// Create second item
	itemOptions = []items.Option{
		items.WithName("sabji"),
		items.WithCategory("curry"),
		items.WithVeg(true),
	}
	item2, err := items.CreateNewItem(ctx, itemOptions...)
	if err != nil {
		log.Default().Print(err.Error())
	} else {
		fmt.Println("New item added ", *item2)
	}

	// Create first menu
	menuItemMapping := map[string]float32{}
	menuItemMapping[item1.ID] = 10.00
	menuItemMapping[item2.ID] = 20.00

	menuOptions := []menu.Option{
		menu.WithRestaurantID(restaurant1.Id),
		menu.WithItemsPriceMapping(menuItemMapping),
	}

	menu1, err := menu.CreateMenu(ctx, menuOptions...)
	if err != nil {
		log.Default().Print(err.Error())
	} else {
		fmt.Println("New menu added ", *menu1)
	}

	menu1, err = menu.AddNewItem(ctx, menu1.ID, *item1, 15.00)
	if err != nil {
		log.Default().Print(err.Error())
	} else {
		fmt.Println("Updated menu ", *menu1)
	}

	menu1, err = menu.AddNewItem(ctx, menu1.ID, *item2, 25.00)
	if err != nil {
		log.Default().Print(err.Error())
	} else {
		fmt.Println("Updated menu ", *menu1)
	}

	menuOptions = []menu.Option{
		menu.WithRestaurantID(restaurant2.Id),
		menu.WithItemsPriceMapping(menuItemMapping),
	}

	menu2, err := menu.CreateMenu(ctx, menuOptions...)
	if err != nil {
		log.Default().Print(err.Error())
	} else {
		fmt.Println("New menu added ", *menu2)
	}

	//Create Order 1
	order1, err := order.CreateNewOrder(ctx)
	if err != nil {
		log.Default().Print(err.Error())
	} else {
		fmt.Println("New order added ", *order1)
	}

	updatedOrder1, err := order.AddItem(ctx, restaurant1.Id, item1.ID, 2, order1)
	if err != nil {
		log.Default().Print(err.Error())
	} else {
		fmt.Println("Updated order added ", *updatedOrder1)
	}

	updatedOrder1, err = order.AddItem(ctx, restaurant1.Id, item2.ID, 1, order1)
	if err != nil {
		log.Default().Print(err.Error())
	} else {
		fmt.Println("Updated order added ", *updatedOrder1)
	}
}
