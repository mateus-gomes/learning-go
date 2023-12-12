package main

import (
	model "exercises-struct/model"
	"fmt"
)

func main() {
	var shoppingList = model.NewShop("Carrefour", []string{"Rice", "Beans"})

	fmt.Println(shoppingList)
}
