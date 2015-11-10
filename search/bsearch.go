package main

import "fmt"

func Bsearch(s []int, value int, low int, high int) int {
	// Not Found
	if low > high {
		return -1
	}

	middle := (high + low) / 2

	// check if the middle item is the value we are searching for
	if s[middle] == value {
		return middle
	}

	// check if the value is in the top or bottom half
	if s[middle] > value {
		return Bsearch(s, value, low, middle-1)
	} else {
		return Bsearch(s, value, middle+1, high)
	}
}

func Reverse(s string) string {
	var result string
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	low := 0
	high := 5
	value := -1

	pos := Bsearch(s, value, low, high)
	fmt.Println(pos)

	reversed := Reverse("hello")
	fmt.Println(reversed)
}
