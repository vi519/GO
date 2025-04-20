package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[2] = "Leachyee"
	fruitList[3] = "Mango"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

}
