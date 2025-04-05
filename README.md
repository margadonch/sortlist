# Project SortList

## Description
**SortList** is a Go-based program designed to sort numerical data in ascending or descending order. The program allows users to input data manually, processes it based on the selected sorting order, and outputs the sorted result. It is designed to be simple, modular, and extensible, making it easy to integrate additional features such as advanced sorting algorithms, data visualization, and external data sources.

## Features
- Accepts numerical input from the user.
- Supports sorting in ascending or descending order.
- Modular design with a reusable `SortList` function.
- Easy to extend with additional sorting algorithms or data sources.

## How It Works
1. **Input Data:** The program prompts the user to enter a list of numbers separated by spaces.
2. **Sorting Order:** The user specifies whether the data should be sorted in ascending or descending order.
3. **Sorting Logic:** The program uses the `SortList` function to sort the data based on the selected order.
4. **Output:** The sorted data is displayed in the console.

## Configuration
The program is designed to be simple and requires no external configuration. However, it can be extended with the following:
- **Additional Sorting Algorithms:** Add new sorting algorithms to the `SortList` function.
- **External Data Sources:** Replace manual input with data from files, databases, or APIs.
- **Visualization:** Integrate with visualization libraries or tools to display data graphically.

## Example Usage
### Input and Sorting
1. Run the program:
   ```bash
   go run main.go
   ```
2. Enter numbers separated by spaces:
   ```
   Enter numbers separated by spaces:
   42 23 16 15 8 4
   ```
3. Choose sorting order:
   ```
   Sort in ascending order? (yes/no):
   yes
   ```
4. Output:
   ```
   Sorted array: [4 8 15 16 23 42]
   ```

### Code Example
The sorting logic is implemented in the `SortList` function:
```go
func SortList(numbers []int, ascending bool) []int {
    sorted := make([]int, len(numbers))
    copy(sorted, numbers)
    sort.Ints(sorted)
    if !ascending {
        for i, j := 0, len(sorted)-1; i < j; i, j = i+1, j-1 {
            sorted[i], sorted[j] = sorted[j], sorted[i]
        }
    }
    return sorted
}
```

## Extending Functionality
### Adding New Sorting Algorithms
To add a new sorting algorithm:
1. Define the algorithm in a new function.
2. Modify the `SortList` function to include the new algorithm as an option.

### Integrating External Data Sources
Replace the manual input logic in the `main` function with code to fetch data from:
- Files (e.g., CSV or JSON).
- Databases (e.g., MySQL, PostgreSQL).
- APIs (e.g., REST or GraphQL).

### Visualization
Integrate with visualization tools such as:
- **Charts:** Use libraries like `go-echarts` to create charts.
- **Web Interfaces:** Build a web interface using Go's `net/http` package or frameworks like Gin.

## Limitations
- Currently supports only integer sorting.
- Input is limited to manual entry via the console.
- No error handling for invalid input (e.g., non-numeric values).

## Future Enhancements
- Add support for floating-point numbers and strings.
- Implement additional sorting algorithms (e.g., merge sort, heap sort).
- Integrate with real-time data sources and consumer services.
- Add graphical visualization of sorting steps.

## License
This project is licensed under the GNU General Public License v3.0. You may copy, distribute, and modify the software as long as you track changes/dates in source files. Any modifications to this project must also be open-sourced under the same license.

For more details, see the [LICENSE](https://www.gnu.org/licenses/gpl-3.0.en.html) file.

## Contact
If you have any questions or suggestions, feel free to reach out to me via GitHub.
