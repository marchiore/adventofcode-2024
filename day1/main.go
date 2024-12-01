package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	leftList := []int{}
	rightList := []int{}

	result := 0
	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), "   ")
		i, err := strconv.Atoi(numbers[0])

		if err == nil {
			leftList = append(leftList, i)
		}

		i, err = strconv.Atoi(numbers[1])

		if err == nil {
			rightList = append(rightList, i)
		}
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	for _, element := range leftList {

		totalOcurrences := 0
		for _, value := range rightList {
			if element == value {
				totalOcurrences++
			}
		}

		result += element * totalOcurrences
	}
	fmt.Println("%d", result)
}
