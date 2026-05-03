package main

import "fmt"

func main() {

	numbers := []int{5, 9, 25, 100}
	total := 0

	for _, number := range numbers {

		total = total + number

	}
	fmt.Println(total)
}
