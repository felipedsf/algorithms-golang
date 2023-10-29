package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	bubbleSort(slice)
	printSlice(slice, 40)

	checkSorted(slice)
}

func bubbleSort(slice []int) {
	sorted := false

	for !sorted {
		changes := 0
		for i := 1; i < len(slice); i++ {
			if slice[i-1] > slice[i] {
				slice[i-1], slice[i] = slice[i], slice[i-1]
				changes++
			}
		}
		if changes > 0 {
			changes = 0
		} else {
			sorted = true
		}
	}
}

func makeRandomSlice(numItems, max int) []int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	slice := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		slice[i] = random.Intn(max)
	}
	return slice
}

func printSlice(slice []int, numItems int) {
	if len(slice) <= numItems {
		fmt.Println(slice)
	} else {
		fmt.Println(slice[:numItems])
	}
}

func checkSorted(slice []int) {
	for i := 0; i < len(slice); i++ {
		if i > 0 {
			if slice[i-1] > slice[i] {
				fmt.Println("The slice is NOT sorted!")
				return
			}
		}
	}
	fmt.Println("The slice is sorted!")

}
