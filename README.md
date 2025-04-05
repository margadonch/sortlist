# Project SortList

## Description
**SortList** is a Go-based program designed to sort numerical data in ascending or descending order. It supports interactive user input, data processing from multiple sources, and visualization of initial, intermediate, and final results. The program is flexible and can be extended to handle streaming data from multiple providers and deliver results to multiple consumer services.

## Features
- Accepts numerical input from users or data streams.
- Supports ascending or descending sorting based on user preference.
- Visualizes data at different stages (initial, intermediate, and final).
- Can be integrated with multiple data sources and consumer services.

## How It Works
1. **Input Data:** The program accepts numerical data from the user or streaming sources.
2. **Sorting:** Data is sorted in ascending or descending order based on user input.
3. **Visualization:** The program visualizes the data at different stages:
   - **Initial Data:** Before sorting.
   - **Intermediate Data:** During processing (if applicable).
   - **Final Data:** After sorting.
4. **Streaming Integration:** The program can be extended to handle data from multiple providers and send results to multiple consumer services.

## Example Usage
### Input from User
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

### Example with Three Data Sources and Three Consumer Services
#### Data Sources:
1. **Provider A:** Sends a stream of random numbers.
2. **Provider B:** Sends a stream of user-generated data.
3. **Provider C:** Sends a stream of sensor data.

#### Consumer Services:
1. **Service X:** Receives sorted data for analytics.
2. **Service Y:** Receives sorted data for visualization.
3. **Service Z:** Receives sorted data for storage.

#### Workflow:
1. Data is streamed from the three providers into the program.
2. The program processes and sorts the data.
3. Sorted data is sent to the three consumer services.

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/sortlist.git
   cd sortlist
   ```
2. Ensure you have Go installed on your system. You can download it from [Go's official website](https://go.dev/).

## Configuration
To integrate with multiple data sources and consumer services:
1. Modify the `main.go` file to include logic for streaming data from providers.
2. Add HTTP or WebSocket clients to send sorted data to consumer services.

## License
This project is licensed under the GNU General Public License v3.0. You may copy, distribute, and modify the software as long as you track changes/dates in source files. Any modifications to this project must also be open-sourced under the same license.

For more details, see the [LICENSE](https://www.gnu.org/licenses/gpl-3.0.en.html) file.

## Contact
If you have any questions or suggestions, feel free to reach out to me via GitHub.
