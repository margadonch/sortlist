package main

import (
    "fmt"
    "sort"
    "strings"
)

// SortList sorts a slice of integers in ascending or descending order.
func SortList(numbers []int, ascending bool) []int {
    // Make a copy of the slice to avoid modifying the original
    sorted := make([]int, len(numbers))
    copy(sorted, numbers)

    // Sort in ascending order
    sort.Ints(sorted)

    // Reverse the order if descending is required
    if !ascending {
        for i, j := 0, len(sorted)-1; i < j; i, j = i+1, j-1 {
            sorted[i], sorted[j] = sorted[j], sorted[i]
        }
    }

    return sorted
}

func main() {
    // Prompt user for input
    fmt.Println("Enter numbers separated by spaces:")
    var input string
    fmt.Scanln(&input)

    // Parse input into a slice of integers
    var numbers []int
    for _, s := range strings.Fields(input) {
        var num int
        fmt.Sscanf(s, "%d", &num)
        numbers = append(numbers, num)
    }

    // Ask user for sorting order
    fmt.Println("Sort in ascending order? (yes/no):")
    var order string
    fmt.Scanln(&order)
    ascending := strings.ToLower(order) == "yes"

    // Sort the numbers
    sorted := SortList(numbers, ascending)

    // Print the sorted array
    fmt.Println("Sorted array:", sorted)
}
