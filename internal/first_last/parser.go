package first_last

import (
	"fmt"
	"unicode"
)

func Parse(line []byte) (int, error) {
	var nums []int

	l := string(line)
	// fmt.Printf("Parse: found line: %v\n", l)

	for _, c := range l {
		if !unicode.IsNumber(c) {
			continue
		}

		num := int(c) - '0'
		// fmt.Printf("Parse: found number: %v\n", num)
		nums = append(nums, num)
	}

	if len(nums) < 1 {
		return 0, fmt.Errorf("Parse: no numbers found")
	}

	first := nums[0]
	last := nums[len(nums)-1]
	// fmt.Printf("Parse: first: %d last: %d\n", first, last)

	return first*10 + last, nil
}
