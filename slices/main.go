package main

import "fmt"

func main() {
	fmt.Println("welcome to learn go lang")
	//slices
	var fruitList = []string{"Apple", "Mango", "Peach"}
	fmt.Println("Type of fruits %T \n", fruitList)

	//var fruitList1 = []string{"Apple","Mango","Peach"}

	//to do append in slice
	//fruitList = append(fruitList, "lychee", "Pompogrenet", "Watermelons", "MaskMelons")
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

}
