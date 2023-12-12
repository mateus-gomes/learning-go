package model

import "time"

type ShoppingList struct {
	MarketName string
	ShopDate   time.Time
	Items      []ShopItem
}

type ShopItem struct {
	ItemName string
}

func NewShop(marketName string, itemsNames []string) ShoppingList {
	var items []ShopItem

	for _, name := range itemsNames {
		items = append(items, ShopItem{ItemName: name})
	}

	return ShoppingList{
		MarketName: marketName,
		ShopDate:   time.Now().Local(),
		Items:      items,
	}
}
