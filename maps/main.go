package main

import "fmt"

func main() {
	shoppingList := make(map[string]int) // default value for integer is 0

	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1 // default value (0) + 1 = 1

	shoppingList["eggs"] += 1 // 11 + 1 = 12
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("Milk deleted, new list:", shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]
	fmt.Println("Need cereal?")
	if !found {
		fmt.Println("Nope")
	} else {
		fmt.Println("Yes, we need", cereal, "boxes of cereal")
	}

	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}

	fmt.Println("Total items:", totalItems)
}
