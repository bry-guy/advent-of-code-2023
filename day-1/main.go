package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"time"
)

// --- Day 1: Trebuchet?! ---
// The calibration document consists of lines of text; each line originally contained a specific value.
// On each line, **the value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number**.

// For example:

// 1abc2
// pqr3stu8vwx
// a1b2c3d4e5f
// treb7uchet
// In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

// Consider your entire calibration document. What is the sum of all of the calibration values?

func main() {
	start := time.Now()

	// Read file
	if len(os.Args) < 2 {
		fmt.Println("error: requires one argument (filepath)")
	}

	filepath := os.Args[1]
	fmt.Printf("filepath: %s\n", filepath)

	buf, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("error: unable to read file at %s\n", filepath)
	}

	reader := bufio.NewReader(bytes.NewReader(buf))

	var sum int
	if len(os.Args) == 3 && os.Args[2] == "--async" {
		sum = day1_async(reader)
	} else {
		sum = day1(reader)
	}

	fmt.Printf("sum: %v\n", sum)

	end := time.Now()
	duration := end.UnixMicro() - start.UnixMicro()

	fmt.Printf("duration: %v us\n", duration)
}
