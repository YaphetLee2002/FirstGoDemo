package main

import "fmt"

func main() {
	var stockade = "%d and %d"
	var b = 456
	var c = 789
	var string = fmt.Sprintf(stockade, b, c)
	fmt.Println(string)
}
