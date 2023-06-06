package main

import "fmt"

func main(){

	//Create

	//1st type of declaration
	//var population map[string]int

	//2nd way
	//population := make(map[string]int)

	//population["Mumbai"] := 10167

	//3rd way
	population := map[string]int {
		"Hyderabad":1000,
		"Chennai":1119,
		"Mumbai": 1200,
		"Lucknow": 567,
	}

	//Retreive
	fmt.Println(population)
	fmt.Println(population["Chennai"])

	//Update 
	population["Hyderabad"] = 1500 
	fmt.Println(population["Hyderabad"])

	//delete
	delete(population, "Chennai")
	fmt.Println(population["Chennai"])


	printMap(population)

}

func printMap(m map[string]int){
	for city, population := range m{
		fmt.Println(city, population)
	}
}