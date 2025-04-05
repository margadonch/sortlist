package main

import (
    "database/sql"
    "fmt"
    "sort"
    "strings"

    _ "github.com/go-sql-driver/mysql"      // MySQL driver
    _ "github.com/lib/pq"                   // PostgreSQL driver
    _ "github.com/nakagami/firebirdsql"     // Firebird driver
    _ "github.com/mattn/go-sqlite3"         // SQLite driver
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

// FetchDataFromDB fetches data from a database table.
func FetchDataFromDB(db *sql.DB, query string) ([]int, error) {
	rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var numbers []int
    for rows.Next() {
        var num int
        if err := rows.Scan(&num); err != nil {
            return nil, err
        }
        numbers = append(numbers, num)
    }
    return numbers, nil
}

// SaveDataToDB saves sorted data into a database table.
func SaveDataToDB(db *sql.DB, tableName string, data []int) error {
    // Clear the table before inserting new data
    _, err := db.Exec(fmt.Sprintf("DELETE FROM %s", tableName))
    if err != nil {
        return err
    }

    // Insert sorted data
    for _, num := range data {
        _, err := db.Exec(fmt.Sprintf("INSERT INTO %s (value) VALUES (?)", tableName), num)
        if err != nil {
            return err
        }
    }
    return nil
}

func main() {
    // Connect to the database (example: SQLite)
    db, err := sql.Open("sqlite3", "example.db")
    if err != nil {
        fmt.Println("Error connecting to the database:", err)
        return
    }
    defer db.Close()

    // Fetch data from the database
    query := "SELECT value FROM numbers"
    numbers, err := FetchDataFromDB(db, query)
    if err != nil {
        fmt.Println("Error fetching data from the database:", err)
        return
    }

    // If no data in the database, prompt user for input
    if len(numbers) == 0 {
        fmt.Println("Enter numbers separated by spaces:")
        var input string
        fmt.Scanln(&input)

        // Parse input into a slice of integers
        for _, s := range strings.Fields(input) {
            var num int
            fmt.Sscanf(s, "%d", &num)
            numbers = append(numbers, num)
        }
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

    // Save sorted data back to the database
    err = SaveDataToDB(db, "numbers", sorted)
    if err != nil {
        fmt.Println("Error saving data to the database:", err)
        return
    }
    fmt.Println("Sorted data saved to the database.")
}
