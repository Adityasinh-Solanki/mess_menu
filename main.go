package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

type Menu struct {
	Day   string     `json:"day"`
	Date  string     `json:"date"`
	Meal  string     `json:"meal"`
	Items [10]string `json:"items"`
}

func print_items(cols [][]string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the name of the day :- ")
	scanner.Scan()
	day := scanner.Text()
	fmt.Println("Enter the name of the meal :- ")
	scanner.Scan()
	meal := scanner.Text()
	day = strings.ToUpper(day)
	meal = strings.ToUpper(meal)
	day = strings.TrimSpace(day)
	meal = strings.TrimSpace(meal)
	var i int
	for i = 0; i < 7; i++ {
		if cols[i][0] == day {
			break
		}
	}
	if i == 7 {
		fmt.Println()
		fmt.Println("Invalid input , words are misspelled or not in the Sample Menu")
		return
	}
	var j int
	for j = 3; j < len(cols[i]); j++ {
		if cols[i][j-1] == meal {
			break
		}
	}
	if j == len(cols[i]) {
		fmt.Println()
		fmt.Println("Invalid input , words are misspelled or not in the Sample Menu")
		return
	}
	for k := j; k < len(cols[i]); k++ {
		if cols[i][k] == day || cols[i][k] == "" {
			break
		}
		fmt.Println("\t", cols[i][k])
	}
}

func return_items(cols [][]string) int {
	count := 0
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the name of the day :- ")
	scanner.Scan()
	day := scanner.Text()
	fmt.Println("Enter the name of the meal :- ")
	scanner.Scan()
	meal := scanner.Text()
	day = strings.ToUpper(day)
	meal = strings.ToUpper(meal)
	day = strings.TrimSpace(day)
	meal = strings.TrimSpace(meal)
	var i int
	for i = 0; i < 7; i++ {
		if cols[i][0] == day {
			break
		}
	}
	if i == 7 {
		fmt.Println()
		fmt.Println("Invalid input , words are misspelled or not in the Sample Menu")
		return 0
	}
	var j int
	for j = 3; j < len(cols[i]); j++ {
		if cols[i][j-1] == meal {
			break
		}
	}
	if j == len(cols[i]) {
		fmt.Println()
		fmt.Println("Invalid input , words are misspelled or not in the Sample Menu")
		return 0
	}
	for k := j; k < len(cols[i]); k++ {
		if cols[i][k] == day || cols[i][k] == "" {
			break
		}
		count++
	}
	return count
}

func is_item(cols [][]string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the name of the day :- ")
	scanner.Scan()
	day := scanner.Text()
	fmt.Println("Enter the name of the meal :- ")
	scanner.Scan()
	meal := scanner.Text()
	fmt.Println("Enter the name of the dish you want to search :- ")
	scanner.Scan()
	key := scanner.Text()
	day = strings.ToUpper(day)
	meal = strings.ToUpper(meal)
	key = strings.ToUpper(key)
	day = strings.TrimSpace(day)
	meal = strings.TrimSpace(meal)
	key = strings.TrimSpace(key)
	var i int
	for i = 0; i < 7; i++ {
		if cols[i][0] == day {
			break
		}
	}
	if i == 7 {
		fmt.Println()
		fmt.Println("Invalid input , words are misspelled or not in the Sample Menu")
		return
	}
	var j int
	for j = 3; j < len(cols[i]); j++ {
		if cols[i][j-1] == meal {
			break
		}
	}
	if j == len(cols[i]) {
		fmt.Println()
		fmt.Println("Invalid input , words are misspelled or not in the Sample Menu")
		return
	}
	for k := j; k < len(cols[i]); k++ {
		if cols[i][k] == day || cols[i][k] == "" {
			break
		}
		if key == cols[i][k] {
			fmt.Println()
			fmt.Println("The item exists in the given meal")
			return
		}
	}
	fmt.Println()
	fmt.Println("The item does not exist in the given meal")
}

func xlsx_json(cols [][]string) {
	var menu [21]Menu
	count := 0
	for i := 0; i < 7; i++ {
		row_count := 2
		for j := 0; j < 3; j++ {
			menu[count].Day = cols[i][0]
			menu[count].Date = cols[i][1]
			menu[count].Meal = cols[i][row_count]
			row_count++
			for k := 0; k < 10; k++ {
				if row_count == len(cols[i]) || cols[i][row_count] == cols[i][0] {
					row_count++
					break
				}
				menu[count].Items[k] = cols[i][row_count]
				row_count++
			}
			count++
		}
	}
	file, err := os.Create("Sample-Menu.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(menu)
	if err != nil {
		panic(err)
	}
}

func struct_json() {
	var menu1 [21]Menu
	file, err := os.Open("Sample-Menu.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	temp := json.NewDecoder(file)
	err = temp.Decode(&menu1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Printing the structures made from Sample-Menu.json")
	fmt.Println()
	for i := 0; i < 21; i++ {
		fmt.Println("\t", menu1[i].Day)
		fmt.Println("\t", menu1[i].Date)
		fmt.Println("\t", menu1[i].Meal, " :")
		for j := 0; j < 10; j++ {
			fmt.Println("\t\t", menu1[i].Items[j])
		}
		fmt.Println()
	}
}

func main() {
	f, err := excelize.OpenFile("Sample-Menu.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	cols, err := f.GetCols("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Executing the functions
	{
		fmt.Println("Function to print a particular meal")
		fmt.Println()
		print_items(cols)
		fmt.Println()
		fmt.Println("Function to return number of items in a meal")
		fmt.Println()
		temp := return_items(cols)
		if temp != 0 {
			fmt.Println("\t", temp)
		}
		fmt.Println()
		fmt.Println("Function to check if an item is the part of the given meal")
		fmt.Println()
		is_item(cols)
		fmt.Println()
		fmt.Println("Creating Sample-Menu.json file")
		fmt.Println()
		xlsx_json(cols)
		fmt.Println("Sample-Menu.json created")
		fmt.Println()
		fmt.Println("Creating an array of structures from the Sample-Menu.json file")
		fmt.Println()
		struct_json()
		fmt.Println("Structures created succesfully")
	}
}

