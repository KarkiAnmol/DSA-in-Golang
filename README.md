# DSA in Golang

## Overview
This repository showcases Data Structures and Algorithms (DSA) implemented in Go. It includes the implementations of sorting algorithms such as Insertion Sort and Merge Sort, along with performance measurement scripts that record and visualize each algorithm's execution time in response to various input sizes.

## Getting Started

### Prerequisites
Make sure you have Go installed on your system. You can download it from [the Go website](https://golang.org/dl/).

### Installing Dependencies
Some scripts in this repository require external Go packages. Install these dependencies by running:

```sh
go get -u gonum.org/v1/plot
```

This command retrieves the `gonum/plot` package which is used for generating graphical data representations.

## Directory Structure

- `/insertion`: Contains the Insertion Sort algorithm and its test cases.
- `/merge`: Contains the Merge Sort algorithm and its test cases.
- `/performanceMeasurement`: Contains scripts for performance testing.
- `/plot`: Contains scripts for generating plots from the performance data.

## Usage Instructions

### Sorting Algorithms
Navigate to a sorting algorithm's directory to execute its implementation and associated tests.

#### Insertion Sort
```sh
cd insertion
go run insertion_sort.go # Executes the Insertion Sort algorithm
go test                  # Runs the test suite
```

#### Merge Sort
```sh
cd merge
go run merge_sort.go # Executes the Merge Sort algorithm
go test              # Runs the test suite
```

### Performance Analysis
Run the performance test script to measure the sorting algorithms' performance:

```sh
cd performanceMeasurement
go run performance_test.go
```
This will output a `performance_data.csv` file with the timing data.

### Graph Generation
To visualize the performance data, generate graphs with the plotting script:

```sh
cd plot
go run plot.go
```
The script will produce a PNG image plotting execution time against input size.

## Contributing
If you wish to contribute to this project, feel free to fork the repository, make your changes, and submit a pull request.

## License
The code in this project is open source.

---
