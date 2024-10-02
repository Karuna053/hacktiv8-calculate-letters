package main

import "fmt"

func main() {
	inputString := "selamat malam"
	mapResult := make(map[string]int)

	for _, char := range inputString { // Use _ to ignore the index, also for some reason iterating over a string per character produces rune?
		fmt.Println(string(char))
		mapResult[string(char)]++
	}

	fmt.Println(mapResult)

	fmt.Println("Press Enter to exit program!")
	fmt.Scanln() // Wait for input
}
