package first_last_test

import (
	"fmt"
	"testing"

	"github.com/bry-guy/advent-of-code-2023/internal/first_last"
	"github.com/stretchr/testify/assert"
)

func TestFindValue(t *testing.T) {
	tests := []struct {
		line []byte
		err  error
		want int
	}{
		{
			line: []byte("abc"),
			err:  fmt.Errorf("Parse: no numbers found"),
			want: 0,
		},
		{
			line: []byte(""),
			err:  fmt.Errorf("Parse: no numbers found"),
			want: 0,
		},
		{
			line: []byte("1abc2"),
			err:  nil,
			want: 12,
		},
		{
			line: []byte("pqr3stu8vwx"),
			err:  nil,
			want: 38,
		},
		{
			line: []byte("a1b2c3d4e5f"),
			err:  nil,
			want: 15,
		},
		{
			line: []byte("treb7uchet"),
			err:  nil,
			want: 77,
		},
	}

	for _, tc := range tests {
		result, err := first_last.Parse(tc.line)

		assert.Equal(t, tc.err, err)
		assert.Equal(t, tc.want, result)
	}
}
