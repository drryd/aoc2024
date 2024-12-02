package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

func main() {

	list1, list2 := parseFile("input.txt")

	sort.Ints(list1)
	sort.Ints(list2)

	totalDistance := 0

	for i := 0; i < len(list1); i++ {
		totalDistance += abs(list1[i] - list2[i])
	}

	fmt.Println("distance:", totalDistance)
}

func parseFile(filename string) ([]int, []int) {
	var list1 []int
	var list2 []int

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		line = strings.TrimSpace(line)
		parts := strings.Fields(line)

		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Error converting string to integer:", err1, err2);
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file")
	}

	return list1, list2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
