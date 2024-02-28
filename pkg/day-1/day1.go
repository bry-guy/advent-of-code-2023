package day1

import (
	"bufio"
	"fmt"

	"github.com/bry-guy/advent-of-code-2023/internal/first_last"
)

func day1(reader *bufio.Reader) int {
	var sum int
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}

		val, err := first_last.Parse(line)
		if err != nil {
			fmt.Printf("error: %e\n", err)
		}

		sum += val
	}

	return sum
}
