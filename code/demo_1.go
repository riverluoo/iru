package main

import "fmt"

func main() {

	const name string = "AAA"
	fmt.Println(name)

	const name1, name2 string = "AA", "BB"
	fmt.Println(name1, name2)

	const name3, name4 = "AA", 123
	fmt.Println(name3, name4)

}
