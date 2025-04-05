# Project SortList

## Description
**SortList** is a Go-based program designed to sort numerical and textual data in ascending or descending order. The program supports manual input, database integration (MySQL, PostgreSQL, Firebird, SQLite), and provides modular sorting functionality. It is designed to be flexible, extensible, and user-friendly, making it suitable for a variety of use cases.

## Features
- **Sorting Capabilities:**
  - Supports sorting of integers, floating-point numbers, and text.
  - Allows sorting in ascending or descending order.
- **Database Integration:**
  - Fetches data from databases (MySQL, PostgreSQL, Firebird, SQLite).
  - Saves sorted results back to the database.
- **Error Handling:**
  - Handles invalid input gracefully.
  - Provides meaningful error messages for database and input issues.
- **Modular Design:**
  - Reusable `SortList` function for sorting integers.
  - Extendable to support additional data types and sources.

## How It Works
1. **Input Data:**
   - The program prompts the user to enter numbers or text separated by spaces.
   - Alternatively, it can fetch data from a connected database.
2. **Sorting Order:**
   - The user specifies whether the data should be sorted in ascending or descending order.
3. **Sorting Logic:**
   - The program sorts the data using Go's built-in sorting functions.
4. **Output:**
   - The sorted data is displayed in the console.
   - If database integration is enabled, the sorted data is saved back to the database.

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/sortlist.git
   cd sortlist
   ```
2. Install Go from [Go's official website](https://go.dev/).
3. Install the required database drivers:
   ```bash
   go get -u github.com/go-sql-driver/mysql
   go get -u github.com/lib/pq
   go get -u github.com/nakagami/firebirdsql
   go get -u github.com/mattn/go-sqlite3
   ```

## Configuration
### Database Setup
1. Create a table named `numbers` in your database with the following schema:
   ```sql
   CREATE TABLE numbers (value INTEGER);
   ```
2. Update the database connection string in the code:
   - **MySQL:**
     ```go
     db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
     ```
   - **PostgreSQL:**
     ```go
     db, err := sql.Open("postgres", "user=postgres password=yourpassword dbname=yourdb sslmode=disable")
     ```
   - **Firebird:**
     ```go
     db, err := sql.Open("firebirdsql", "sysdba:masterkey@localhost:3050/yourdb.fdb")
     ```
   - **SQLite:**
     ```go
     db, err := sql.Open("sqlite3", "example.db")
     ```

### Running the Program
1. Run the program:
   ```bash
   go run main.go
   ```
2. Follow the prompts to enter data or fetch it from the database.

## Example Usage
### Manual Input
1. Enter numbers or text separated by spaces:
   ```
   Enter numbers separated by spaces:
   42 3.14 apple 23 banana 16.5
   ```
2. Choose sorting order:
   ```
   Sort in ascending order? (yes/no):
   yes
   ```
3. Output:
   ```
   Sorted numbers: [3.14 16.5 23 42]
   Sorted text: [apple banana]
   ```

### Database Integration
1. Fetch data from the database:
   ```
   SELECT value FROM numbers;
   ```
2. Sort the data and save the results back to the database.

## Extending Functionality
1. **Additional Sorting Algorithms:**
   - Add new algorithms to the `SortList` function.
2. **Support for Mixed Data Types:**
   - Extend the logic to handle mixed sorting of numbers and text.
3. **Visualization:**
   - Integrate with visualization libraries like `go-echarts` or build a web interface.

## License
This project is licensed under the GNU General Public License v3.0. You may copy, distribute, and modify the software as long as you track changes/dates in source files. Any modifications to this project must also be open-sourced under the same license.

For more details, see the [LICENSE](https://www.gnu.org/licenses/gpl-3.0.en.html) file.

## Contact
If you have any questions or suggestions, feel free to reach out to me via GitHub.
