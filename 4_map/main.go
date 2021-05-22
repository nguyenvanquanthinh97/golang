package main

import "fmt"

func main() {
	// Declare map variable - 3 ways
	
	// Method - 1
	// colors := map[string]string{
	// 	"red": "ff0000",
	// 	"green": "00ff00",
	// 	"blue": "0000ff",
	// }

	// Method - 2
	// Map types are reference types, like pointers or slices, and so the value of m above is nil;
	// it doesn't point to an initialized map. A nil map behaves like an empty map when reading,
	// but attempts to write to a nil map will cause a runtime panic; don't do that.
	// To initialize a map, use the built in make function.
	// var colors map[string]string

	// Method - 3
	colors := make(map[string]string)
	
	colors["red"] = "f00"
	colors["green"] = "0f0"
	colors["blue"] = "00f"

	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, "-", value)
	}
}
