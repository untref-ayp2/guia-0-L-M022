package main

import "fmt"

func Swap(px, py *int) {
	aux := *px
	*px = *py
	*py = aux
}
func main() {
	px := 7
	py := -3
	fmt.Printf("Pre Swap: %v | %v\n", px, py)
	Swap(&px, &py)
	fmt.Printf("Post Swap: %v | %v", px, py)
}
