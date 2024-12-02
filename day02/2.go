package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	reports := parseFile("input.txt")
	numSafeReports := 0
	
	reportLoop:
	for i := range reports {
		// Iterate over each item in the report, and see if it's safe with any item removed.
		// We could do better by more selectively removing elements, but this is good enough for day 2.
		for j := 0; j < len(reports[i]); j++ {
			reportWithLevelJRemoved := removeAtIndex(reports[i], j)

			if isReportSafe(reportWithLevelJRemoved) {
				numSafeReports++
				continue reportLoop
			}
		}
	}

	fmt.Println(numSafeReports)
}

func isReportSafe(level []int) bool {
	if level[1] - level[0] == 0 {
		return false
	}

	isIncreasing := false

	if level[1] - level[0] > 0 {
		isIncreasing = true
	}

	for i := 1; i < len(level); i++ {
		currItem := level[i]
		prevItem := level[i-1]

		if isIncreasing {
			amountIncreased := currItem - prevItem
			if amountIncreased < 1 || amountIncreased > 3 {
				return false
			}
		} else {
			amountDecreased := prevItem - currItem
			if amountDecreased < 1 || amountDecreased > 3 {
				return false
			}
		}
	}

	return true
}

func parseFile(filename string) ([][]int) {
        var reports [][]int

        file, err := os.Open(filename)

        if err != nil {
                fmt.Println("Error opening file:", err)
                return nil
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)

        for scanner.Scan() {
                line := scanner.Text()

                line = strings.TrimSpace(line)
		parts := strings.Fields(line)

		var levels []int
		for _, part := range parts {
			level, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				return nil
			}
			levels = append(levels, level)
		}

		reports = append(reports, levels)
        }

        if err := scanner.Err(); err != nil {
                fmt.Println("Error reading file")
		return nil
        }

        return reports
}

func removeAtIndex(slice []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, slice[:index]...)

	// If we are removing the last item, there's nothing at slice[index+1]
	if (index < len(slice) - 1) {
		ret = append(ret, slice[index+1:]...)
	}
	
	return ret
}
