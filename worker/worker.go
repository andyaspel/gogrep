package worker

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Result struct {
	Line   string
	LineNr int
	Path   string
}
type Results struct {
	Inner []Result
}

func NewResult(line string, lineNr int, path string) Result {
	return Result{line, lineNr, path}
}

/* this is where I have to add functionality
   that implements various search patterns */

func FindInFile(path string, find string) *Results {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	results := Results{make([]Result, 0)}

	scanner := bufio.NewScanner(file)
	lineNr := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), find) {
			r := NewResult(scanner.Text(), lineNr, path)
			results.Inner = append(results.Inner, r)
		}
		lineNr += 1
	}

	if len(results.Inner) == 0 {
		return nil
	} else {
		return &results
	}
}
