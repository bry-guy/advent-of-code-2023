package main

import (
	"bufio"
	"context"
	"fmt"
	"sync"

	"github.com/bry-guy/advent-of-code-2023/internal/first_last"
	"golang.org/x/sync/errgroup"
)

func day1_async(reader *bufio.Reader) int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var sum int

	lineChan := make(chan []byte)
	parseChan := make(chan int)

	g, _ := errgroup.WithContext(context.Background())

	go func() {
		for {
			bytes, err := reader.ReadBytes('\n')
			if err != nil {
				break
			}
			lineChan <- bytes
		}
		close(lineChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range parseChan {
			mu.Lock()
			sum += val
			mu.Unlock()
		}
	}()

	for line := range lineChan {
		line := line
		g.Go(func() error {
			val, err := first_last.Parse(line)
			if err != nil {
				return err
			}
			parseChan <- val
			return nil
		})
	}

	// Wait for all goroutines to finish and check for errors
	if err := g.Wait(); err != nil {
		fmt.Printf("error processing lines: %v\n", err)
	}
	close(parseChan)

	wg.Wait()

	return sum
}
