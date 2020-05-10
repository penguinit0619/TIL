package main

import "fmt"

//Use var to DECLARE three VARIABLES. The variables should have package level scope.
//Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).
//identifier “x” type int
//identifier “y” type string
//identifier “z” type bool

var a = 42
var b = "Jinyong kim"
var c = true

func main() {
	s := fmt.Sprintf("%d", a) + " " + fmt.Sprintf("%s", b) + " " + fmt.Sprintf("%t", c)
	fmt.Println(s)
}