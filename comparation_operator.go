package main

import "fmt"

func main() {
	var name1 = "Eko"
	var name2 = "Mter"
	var name3 = "Mter"

	var result bool = name1 == name2
	fmt.Println(result)
	result = name1 == name3
	fmt.Println(result)
	result = name2 == name3
	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
