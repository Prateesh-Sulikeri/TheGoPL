package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	processedFiles := []string{}
	if len(files) == 0 {
		countFiles(os.Stdin, counts)
	} else {
		for _, args := range files {
			f, err := os.Open(args)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countFiles(f, counts)
			f.Close()
		}
	}

	for _, f := range processedFiles {
		fmt.Println("\nFiles Processed: ", f)
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t: %s\n", n, line)
		}
	}
}

func countFiles(f *os.File, counts map[string]int) {
	inputs := bufio.NewScanner(f)
	for inputs.Scan() {
		counts[inputs.Text()]++
	}
}
