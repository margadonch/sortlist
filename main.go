package main

import (
    "fmt"
    "sort"
)

func main() {
    // Example array
    numbers := []int{42, 23, 16, 15, 8, 4}

    // Sorting the array in ascending order
    sort.Ints(numbers)

    // Print the sorted array
    fmt.Println("Sorted array:", numbers)
}
