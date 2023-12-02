package main

import (
	"fmt"
	"os"
	"strconv"
  "bufio"
)

// check checks for errors and exits the program if an error is encountered.
func check(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func main() {
	// Open the file
	file, err := os.Open("./data/data.txt")
	check(err)
	defer file.Close()

	// Read the file content
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())

	// Process each line
	sum := 0
	for _, line := range lines {
		fmt.Println(line)
		var first, last int
		fmt.Println()
		for _, ch := range line {
			if num, err := strconv.Atoi(string(ch)); err == nil {
				if first == 0 {
					first = num
				}
				last = num
				fmt.Printf("%q looks like a number.\n", ch)
			}
		}
		fmt.Printf("first: %d, last: %d\n", first, last)
		weight := first*10 + last
		fmt.Printf("Weight: %d\n", weight)
		sum += weight
	}
	fmt.Printf("Sum: %d\n", sum)
}

