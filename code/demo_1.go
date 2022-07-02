package main

import "fmt"

func main() {

	const name string = "AAA"
	fmt.Println(name)

	const name1, name2 string = "AA", "BB"
	fmt.Println(name1, name2)

	const name3, name4 = "AA", 123
	fmt.Println(name3, name4)

	fmt.Println("-----------------------------")

	var a_1 uint8 = 10
	var a_2 = 20
	a_3 := 30
	fmt.Println(a_1, a_2, a_3)

	fmt.Println("-----------------------------")

	var a_4, a_5, a_6 int = 40, 50, 60
	fmt.Println(a_4,a_5,a_6)

	fmt.Println("-----------------------------")


}
