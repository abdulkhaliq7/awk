# Go AWK Package

## Overview
The Go AWK package provides a simple and efficient way to process and manipulate string data in a manner similar to the `awk` command-line tool. This package is designed to handle large amounts of text data efficiently using Go's standard libraries.

## Features
- **Data Splitting**: Split strings by specified delimiters and extract specific columns.
- **String Replacement**: Replace occurrences of a substring within the processed data.

## Installation
To install the package, run:
```go
go get -u github.com/abdulkhaliq/awk
```
## Usage

### 1. Import the Package
```go
import "github.com/abdulkhaliq7/awk"
```
### 2. Create an AWK Instance
```go
data := "John|30|New York"
awkInstance := awk.NewAwk(data)
```
### 3. Split Data

Split data based on a delimiter, optionally choosing specific columns or all columns:
```go
// Split data and include all columns

splitData, err := awkInstance.DataSplit("|", " ")
if err != nil {
    log.Fatalf("DataSplit error: %v", err)
}
fmt.Println("All columns:\n", splitData.Data)
```
```go
// Split data and include only specific columns (e.g., columns 0 and 2)

splitData, err = awkInstance.DataSplit("|", " ", "0", "2")
if err != nil {
    log.Fatalf("DataSplit error: %v", err)
}
fmt.Println("Specific columns (0 and 2):\n", splitData.Data)
```
### 4. Replace Substrings

Replace occurrences of specific substrings in the data:
```go
replacedData := splitData.Replace("John", "abdulkhaliq")
fmt.Println(replacedData.Data)
```

## Contributing
Contributions are welcome! If you have any ideas for new features or improvements, feel free to open an issue or submit a pull request.







