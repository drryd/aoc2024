package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	levels := parseFile("input.txt")
	numSafeLevels := 0
	
	levelLoop:
	for i := range levels {
		if levels[i][1] - levels[i][0] == 0 {
			continue levelLoop
		}

		isIncreasing := false

		if levels[i][1] - levels[i][0] > 0 {
			isIncreasing = true
		}

		for j := 1; j < len(levels[i]); j++ {
			currItem := levels[i][j]
			prevItem := levels[i][j-1]
			if isIncreasing {
				amountIncreased := currItem - prevItem
				if amountIncreased < 1 || amountIncreased > 3 {
					continue levelLoop
				}
			} else {
				amountDecreased := prevItem - currItem
				if amountDecreased < 1 || amountDecreased > 3 {
					continue levelLoop
				}
			}
		}

		numSafeLevels++
	}

	fmt.Println(numSafeLevels)
}

func parseFile(filename string) ([][]int) {
        var levels [][]int

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

		var nums []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				return nil
			}
			nums = append(nums, num)
		}

		levels = append(levels, nums)
        }

        if err := scanner.Err(); err != nil {
                fmt.Println("Error reading file")
		return nil
        }

        return levels
}
