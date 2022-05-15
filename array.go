package main

import "fmt"

func main() {
	var names [4]string

	names[0] = "Menther"
	names[1] = "Goreng"
	names[2] = "Susu"
	names[3] = "Verse"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[3])

	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [9]string

	fmt.Println(len(lagi))
}
