package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "sort"
    "strings"
    "time"
)

// Data represents a single data point from a provider.
type Data struct {
    Provider string
    Value    int
}

// BubbleSort implements the bubble sort algorithm.
func BubbleSort(numbers []int) []int {
    sorted := make([]int, len(numbers))
    copy(sorted, numbers)
    n := len(sorted)
    for i := 0; i < n; i++ {
        for j := 0; j < n-i-1; j++ {
            if sorted[j] > sorted[j+1] {
                sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
            }
        }
    }
    return sorted
}

// SortList sorts a slice of integers using the specified algorithm.
func SortList(numbers []int, algorithm string) []int {
    switch algorithm {
    case "bubble":
        return BubbleSort(numbers)
    case "quick":
        sort.Ints(numbers) // QuickSort (default in Go)
        return numbers
    case "reverse":
        sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
        return numbers
    case "custom1":
        // Placeholder for a custom sorting algorithm
        sort.Ints(numbers)
        return numbers
    case "custom2":
        // Placeholder for another custom sorting algorithm
        sort.Ints(numbers)
        return numbers
    default:
        sort.Ints(numbers) // Default to QuickSort
        return numbers
    }
}

// SimulateProvider simulates a data provider sending random data.
func SimulateProvider(providerName string, dataChannel chan Data) {
    for i := 0; i < 10; i++ {
        dataChannel <- Data{Provider: providerName, Value: rand.Intn(100)}
        time.Sleep(time.Millisecond * 500) // Simulate delay
    }
}

// SendToConsumer sends sorted data to a web portal (consumer).
func SendToConsumer(sortedData []int) {
    // Simulate sending data to a web portal
    fmt.Println("Sending sorted data to web portal:", sortedData)
    // Example HTTP POST request (replace with actual endpoint)
    _, err := http.Post("http://example.com/consumer", "application/json", nil)
    if err != nil {
        fmt.Println("Error sending data to consumer:", err)
    }
}

func main() {
    // Seed random number generator
    rand.Seed(time.Now().UnixNano())

    // Create a channel for streaming data
    dataChannel := make(chan Data)

    // Simulate three providers
    go SimulateProvider("ProviderA", dataChannel)
    go SimulateProvider("ProviderB", dataChannel)
    go SimulateProvider("ProviderC", dataChannel)

    // Collect data from providers
    var numbers []int
    go func() {
        for data := range dataChannel {
            fmt.Printf("Received data from %s: %d\n", data.Provider, data.Value)
            numbers = append(numbers, data.Value)
        }
    }()

    // Wait for user input to sort
    time.Sleep(5 * time.Second) // Simulate data collection period
    fmt.Println("Choose sorting algorithm (bubble, quick, reverse, custom1, custom2):")
    var algorithm string
    fmt.Scanln(&algorithm)

    // Sort the numbers
    sorted := SortList(numbers, algorithm)

    // Print and send the sorted array
    fmt.Println("Sorted array:", sorted)
    SendToConsumer(sorted)
}
