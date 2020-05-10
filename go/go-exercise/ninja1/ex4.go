package main

import "fmt"

type penguin int
var superX penguin
var superY int

func main () {
	fmt.Println(superX)
	fmt.Printf("superX type: %T\n", superX)

	superX = 42
	fmt.Println(superX)

	superY = int(superX)
	fmt.Println(superY)
	fmt.Printf("superY type: %T\n", superY)
}
