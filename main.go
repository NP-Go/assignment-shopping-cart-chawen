package main

// "strconv"
// "math"

import "fmt"

// Category int to itemCat

type item struct {
	itemCat   int
	Quantity  int
	Unit_Cost float64
}

var itemCat []string

var itemMap map[string]item

// Preload
func init() {
	itemCat = []string{"Household", "Food", "Drinks"}

	itemMap = make(map[string]item)

	itemMap["Fork"] = item{0, 4, 3.0}
	itemMap["Plates"] = item{0, 4, 3.0}
	itemMap["Cups"] = item{0, 5, 3.0}

	itemMap["Bread"] = item{1, 2, 2.0}
	itemMap["Cake"] = item{1, 3, 1.0}

	itemMap["Coke"] = item{2, 5, 2.0}
	itemMap["Sprite"] = item{2, 5, 2.0}

	// ? Advance items
}

func main() {

	/*
		    fmt.Println("Shopping Cart")
				fmt.Println("Start main.go")
				cart()
	*/

	var getVal string
	var getCase int

	menuChoices := makeMenuRange(0, len(mainMenu))

	listMenu()

	fmt.Scanln(&getVal)

	getCase, _ = checkVal(getVal)

	for !chkInt(menuChoices, getCase) {
		// For infinity loop
		listMenu()
		fmt.Scanln(&getVal)
		getCase, _ = checkVal(getVal)
	}

	switch getCase {

	case 1:
		viewShopList()
		rtnBack()

	case 2:
		genShopReport()

	case 3:
		addItem()
		rtnBack()

	case 4:
		modItem()
		rtnBack()

	case 5:
		delItem()
		rtnBack()

	case 6:
		// viewCur()
		fmt.Println("Error")
		rtnBack()

	case 7:
		// addNewCat()
		fmt.Println("Error")
		rtnBack()

	case 8:
		fmt.Println("Quit")

		//break
		return
	}

	main()

}
