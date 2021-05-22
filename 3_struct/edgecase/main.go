package main

import "fmt"

func main() {
	str := []string{"Hi", "There", "How", "are", "you?"}

	updateHeadSlice(str, "Bye")
	fmt.Println(str)
}

func updateHeadSlice(s []string, update string) {
	s[0] = update
}
