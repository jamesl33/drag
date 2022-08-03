package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func aggregate() (time.Duration, error) {
	var (
		scanner        = bufio.NewScanner(os.Stdin)
		line    uint64 = 1
		agg     time.Duration
	)

	for scanner.Scan() {
		dur, err := time.ParseDuration(scanner.Text())
		if err != nil {
			return 0, fmt.Errorf("failed to parse duration on line %d: %w", line, err)
		}

		line++
		agg += dur
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("failed to scan line %d: %w", line, err)
	}

	return agg, nil
}

func run() error {
	agg, err := aggregate()
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", agg)

	return nil
}

func main() {
	err := run()
	if err == nil {
		return
	}

	defer os.Exit(1)

	fmt.Printf("Error: %s\n", err)
}
