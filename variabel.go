package main

import "fmt"

func main() {
	var name string

	name = "Menther Goreng"
	fmt.Println(name)

	name = "Menther Susu"
	fmt.Println(name)

	var SecondName = "Susu"
	fmt.Println(SecondName)

	var age = 40
	fmt.Println(age)

	country := "Indopride"
	fmt.Println(country)

	var (
		firstName = "Menther"
		lastName  = "Goreng"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
