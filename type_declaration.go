package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpMter NoKTP = "3300001202000001"
	var marriedStatus Married = false
	fmt.Println(noKtpMter)
	fmt.Println(marriedStatus)
}
