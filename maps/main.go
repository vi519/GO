package main

import "fmt"

func main() {
	fmt.Println("Map is golang")
	superheroes := make(map[string]string)

	superheroes["JuticeLeague"] = "Batman"
	superheroes["Marvel"] = "thor"
	superheroes["IndianHeroes"] = "shaktiman"

	fmt.Println("\nHere are my superhero", superheroes)
	fmt.Println("\nNeed marvel one ", superheroes["Marvel"])

	for key, val := range superheroes {
		fmt.Printf("For key %v , value is %v\n", key, val)
	}

}
