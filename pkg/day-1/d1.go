package day1

import (
	"bufio"
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/spf13/cobra"
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

// var Day1Cmd = &cobra.Command{
// 	Use:   "day1 -f/--filepath {path-to-input} -a/--async",
// 	Short: "Run day1",
// 	Long:  `Run day1 long`,
// 	Args:  cobra.MinimumNArgs(1),
// 	Run:   d1,
// }

type options struct {
	filepath string
	async    bool
}

func NewCmd() *cobra.Command {
	var opts options

	cmd := &cobra.Command{
		Use:   "day1 -f/--filepath {path-to-input} -a/--async",
		Short: "Run day1",
		Long:  `Run day1`,
		Run: func(cmd *cobra.Command, args []string) {
			d1(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.filepath, "filepath", "f", "", "filepath to input (required)")
	err := cmd.MarkFlagRequired("filepath")
	if err != nil {
		slog.Error("unable to mark filepath required", "err", err)
		os.Exit(1)
	}

	cmd.Flags().BoolVarP(&opts.async, "async", "a", false, "enable async processing")

	return cmd
}

func d1(opts *options) {
	start := time.Now()

	buf, err := os.ReadFile(opts.filepath)
	if err != nil {
		fmt.Printf("error: unable to read file at %s\n", opts.filepath)
	}

	reader := bufio.NewReader(bytes.NewReader(buf))

	var sum int
	if opts.async {
		sum = day1_async(reader)
	} else {
		sum = day1(reader)
	}

	fmt.Printf("sum: %v\n", sum)

	end := time.Now()
	duration := end.UnixMicro() - start.UnixMicro()

	fmt.Printf("duration: %v us\n", duration)
}
