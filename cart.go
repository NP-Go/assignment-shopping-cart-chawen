package main

/*
External functions for main.go
? How to remove error in VC
*/

import (
	"fmt"
	"strconv"
	"strings"
)

func cart() {
	fmt.Println("Start cart.go")
}

var mainMenu = []string{
	"1. View entire shopping list.",
	"2. Generate shopping list report",
	"3. Add item",
	"4. Modify items",
	"5. Delete item",
	"6. Print current data",
	"7. Add new category name",
	"8. Quit",
}

func listMenu() {
	fmt.Println("Shopping List App")
	fmt.Println("--------------------")

	for i := range mainMenu {
		fmt.Println(mainMenu[i])
	}

	fmt.Println("--------------------")
	fmt.Println("Select your choice:")
	// fmt.Scanln(&getVal)
}

var subMenu = []string{
	"1. Total cost of each category",
	"2. List of item by category",
	"3. Main menu",
}

func genMenu() {

	fmt.Println("Generate Report")
	fmt.Println("--------------------")

	for j := range subMenu {
		fmt.Println(subMenu[j])
	}

	fmt.Println("--------------------")
}

func makeMenuRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func chkInt(m []int, c int) bool {
	for i := range m {
		if m[i] == c {
			return true
		}
	}

	return false
}

func checkVal(trimVal string) (int, bool) {
	trimVal = strings.TrimSpace(trimVal)
	// Convert string to int
	intVal, err := strconv.Atoi(trimVal)

	if err != nil {
		return -1, false
	}

	return intVal, true
}

// ? Group work around
func rtnBack() {
	fmt.Println("")
	fmt.Println("Press enter to go back main menu")
	fmt.Scanln()
}

// Case 1
func viewShopList() {
	fmt.Println("Shopping List Contents:")
	for key, value := range itemMap {
		fmt.Println(itemCat[value.itemCat] + ": " + key + " - Item: " + strconv.Itoa(value.Quantity) + " Unit Cost: " + fmt.Sprintf("%g", value.Unit_Cost))
	}
}

// Case 2
func genShopReport() {
	var getIf int
	subChoices := makeMenuRange(1, len(subMenu))

	genMenu()
	fmt.Scanln(&getIf)

	for !chkInt(subChoices, getIf) {
		genMenu()
		fmt.Scanln(&getIf)
	}
	// ? Use if
	if getIf == 1 {
		catCost := make([]float64, len(itemCat))
		for k := range itemCat {
			for _, value := range itemMap {
				if value.itemCat == k {
					catCost[k] = catCost[k] + value.Unit_Cost*float64(value.Quantity)
				}
			}
		}
		for l := range itemCat {
			fmt.Println(itemCat[l] + " cost: " + fmt.Sprintf("%g", catCost[l]))
		}

		rtnBack()
		genShopReport()

	} else if getIf == 2 {

		fmt.Println("List by Category")
		for k := range itemCat {
			for key, value := range itemMap {
				if value.itemCat == k {
					fmt.Println(itemCat[value.itemCat] + ": " + key + " - Item: " + strconv.Itoa(value.Quantity) + " Unit Cost: " + fmt.Sprintf("%g", value.Unit_Cost))
				}
			}
		}

		rtnBack()
		genShopReport()

	} else if getIf == 3 {
		main()
		return

	}
}

// Start of Group help
// Case 3

func addItem() {
	var input31, input32 string
	var input33 int
	var input34 float64
	fmt.Println("Name the item you wish to add...?")
	fmt.Scanln(&input31)
	fmt.Println("What category does it belong to?")
	fmt.Scanln(&input32)
	fmt.Println("How many units are there?")
	fmt.Scanln(&input33)
	fmt.Println("How much does it cost per unit?")
	fmt.Scanln(&input34)
	addItems(input31, input32, input33, input34)
}

// Case 4
func modItem() {
	fmt.Println("Modify items")
	var input41 string
	var modItem, modCat, modQty, modUC bool
	fmt.Println("Which item do you wish to modify?")
	fmt.Scanln(&input41)
	_, found := itemMap[input41]
	if found {
		modItem, modCat, modQty, modUC = modifyItem(input41)
		fmt.Println("Item Name is changed: " + strconv.FormatBool(modItem))
		fmt.Println("Category is changed: " + strconv.FormatBool(modCat))
		fmt.Println("Quantity is changed: " + strconv.FormatBool(modQty))
		fmt.Println("Unit-cost is changed: " + strconv.FormatBool(modUC))
	} else {
		fmt.Println("Item is not found. Nothing to modify!")
	}
}

// Case 5
func delItem() {
	fmt.Println("Delete items")
	var input51 string
	fmt.Println("Enter the item to delete")
	fmt.Scanln(&input51)
	_, found := itemMap[input51]
	if found {
		delete(itemMap, input51)
		fmt.Println("Deleted " + input51)
	} else {
		fmt.Println("Item not found. Nothing to delete")
	}
	rtnBack()
}

// Case 6

// Case 7

func viewCur() {
	var newCategory string
	fmt.Println("Add New Category Name")
	fmt.Println("What is the New Category Name to add?")
	fmt.Scanln(&newCategory)
	_ = addNewCategory(newCategory)
	rtnBack()
}

// ? Package
// More of Group help

func findCategory(C []string, c string) (int, bool) {
	for i := range C {
		if strings.ToLower(C[i]) == strings.ToLower(c) {
			return i, true
		}
	}
	return -1, false
}

// add an item to map items
func addItems(itemName, itemCategory string, quantity int, unit_cost float64) {
	var input1 string
	_, categoryFound := findCategory(itemCat, itemCategory)
	if !categoryFound && itemCategory != "" {
		fmt.Printf("\nCategory, %s does not exists. These are the existing categories...:\n", itemCategory)
		for i := range itemCat {
			fmt.Println(" - " + strconv.Itoa(i) + ". " + itemCat[i])
		}
		fmt.Println("Enter 'a' to add " + itemCategory + " in, or 'enter' to ignore this.")
		fmt.Scanln(&input1)
		if input1 == "a" {
			_ = addNewCategory(itemCategory)
		}
	}

	_, found := itemMap[itemName]
	if found {
		_ = addItemQty(itemName, quantity)
		_ = updateItemUnitCost(itemName, unit_cost)
	} else {
		if itemCategory != "" {
			// check if user added the category earlier or not
			index, categoryFound := findCategory(itemCat, itemCategory)
			if categoryFound {
				itemMap[itemName] = item{index, quantity, unit_cost}
			} else {
				// itemCategory is blank and item is new
				fmt.Println("Unable to proceed... item is new, but category wasn't added earlier.")
			}
		} else {
			// itemCategory is blank and item is new
			fmt.Println("Unable to proceed, item is new, but category of it is unknown.")
		}
	}
}

func modifyItem(itemToMod string) (itemNameUpdated, categoryUpdated, qtyUpdated, ucUpdated bool) {
	var itemName, itemCategory, quantity, unit_cost string
	value, _ := itemMap[itemToMod]
	fmt.Println("Current item name is " + itemToMod + " - Category is " + itemCat[value.itemCat] + " - Quantity is " + strconv.Itoa(value.Quantity) + " - Unit Cost " + fmt.Sprintf("%g", value.Unit_Cost))
	fmt.Println("Enter new name. Enter for no change.")
	fmt.Scanln(&itemName)
	itemNameUpdated = updateItemName(itemToMod, itemName)
	if itemNameUpdated {
		itemToMod = itemName
	}

	fmt.Println("Enter new Category. Enter for no change.")
	fmt.Scanln(&itemCategory)
	categoryUpdated = updateCategoryName(itemToMod, itemCategory)

	fmt.Println("Enter new Quantity. Enter for no change.")
	fmt.Scanln(&quantity)
	if quantity != "" {
		qty, err := strconv.Atoi(quantity)
		for err != nil {
			fmt.Println("Enter new Quantity. Enter for no change.")
			fmt.Scanln(&quantity)
			qty, err = strconv.Atoi(quantity)
		}
		qtyUpdated = updateItemQty(itemToMod, qty)
	}

	fmt.Println("Enter new Unit cost. Enter for no change.")
	fmt.Scanln(&unit_cost)
	cost, err := strconv.ParseFloat(unit_cost, 64)
	if unit_cost != "" {
		for err != nil {
			fmt.Println("Enter new Unit cost. Enter for no change.")
			fmt.Scanln(&unit_cost)
			cost, err = strconv.ParseFloat(unit_cost, 64)
		}
		ucUpdated = updateItemUnitCost(itemToMod, cost)
	}
	return itemNameUpdated, categoryUpdated, qtyUpdated, ucUpdated
}

func addNewCategory(newCategory string) int {
	i, found := findCategory(itemCat, newCategory)
	if !found {
		itemCat = append(itemCat, newCategory)
		i, _ = findCategory(itemCat, newCategory)
		fmt.Println("New category: " + newCategory + " added at index " + strconv.Itoa(i))
		return i // return the index position of the new category
	} else {
		fmt.Println("Category: " + newCategory + " already exists at index " + strconv.Itoa(i))
		return i
	}
}

// returns true if s is an integer
func isInt(s string) (int, bool) {
	s = strings.TrimSpace(s)
	intValue, err := strconv.Atoi(s)
	if err != nil {
		return -1, false
	}
	return intValue, true
}

// returns true if s is a string
func isString(s string) (string, bool) {
	s = strings.TrimSpace(s)
	_, err := strconv.Atoi(s)
	if err != nil {
		return s, true
	}
	return "", false
}

// add Item Qty adds the new qty to the item's qty
func addItemQty(itemName string, qty int) bool {
	for key, value := range itemMap {
		if key == itemName {
			value.Quantity = value.Quantity + qty
			itemMap[key] = item{value.itemCat, value.Quantity, value.Unit_Cost}
			return true
		}
	}
	return false
}

func updateItemName(oldName, newName string) bool {
	value, _ := itemMap[oldName]
	if newName != "" {
		itemMap[newName] = item{value.itemCat, value.Quantity, value.Unit_Cost}
		delete(itemMap, oldName)
		return true
	}
	return false
}

func updateCategoryName(itemName, newCat string) bool {
	if newCat != "" {
		i, found := findCategory(itemCat, newCat)
		if found {
			value, ok := itemMap[itemName]
			if ok {
				itemMap[itemName] = item{i, value.Quantity, value.Unit_Cost}
				return true
			}
			return false
		}
		return false
	}
	return false
}

// update Item Qty replaces the item qty with new value
func updateItemQty(itemName string, qty int) bool {
	value, ok := itemMap[itemName]
	if ok {
		itemMap[itemName] = item{value.itemCat, qty, value.Unit_Cost}
		return true
	}
	return false
}

// update Item Unit-Cost replaces the item Unit-Cost with new value
func updateItemUnitCost(itemName string, uc float64) bool {
	value, ok := itemMap[itemName]
	if ok {
		itemMap[itemName] = item{value.itemCat, value.Quantity, uc}
		return true
	}
	return false
}

// End of Group help

func printDataInMem() {
	fmt.Println("Print Current Data.")
	if len(itemMap) != 0 {
		for key, value := range itemMap {
			fmt.Println(key + " - {" + strconv.Itoa(value.itemCat) + " " + strconv.Itoa(value.Quantity) + " " + fmt.Sprintf("%g", value.Unit_Cost) + "}")
		}
	} else {
		fmt.Println("No data found")
	}
}
