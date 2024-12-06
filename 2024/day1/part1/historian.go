package main

import (
	"bufio"
	"fmt"
	"os"
)

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func parseLine(line string) [2]int {
	var nums [2]int
	_, err := fmt.Sscanf(line, "%d   %d", &nums[0], &nums[1])
	errorCheck(err)
	return nums
}

func findLowestIndex(array []int) int {
	lowest := 0
	for i := 1; i < len(array); i++ {
		if array[i] < array[lowest] {
			lowest = i
		}
	}
	return lowest
}

func main() {
	file, err := os.Open("./list.txt")
	errorCheck(err)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var array1 = []int{}
	var array2 = []int{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		nums := parseLine(line)
		array1 = append(array1, nums[0])
		array2 = append(array2, nums[1])
	}

	sum := 0
	lenght := len(array1)
	for i := 0; i < lenght; i++ {
		lowestIndex1 := findLowestIndex(array1)
		lowestIndex2 := findLowestIndex(array2)

		numberToAdd := 0
		if array1[lowestIndex1] < array2[lowestIndex2] {
			numberToAdd = array2[lowestIndex2] - array1[lowestIndex1]
		} else {
			numberToAdd = array1[lowestIndex1] - array2[lowestIndex2]
		}

		sum += numberToAdd

		array1 = append(array1[:lowestIndex1], array1[lowestIndex1+1:]...)
		array2 = append(array2[:lowestIndex2], array2[lowestIndex2+1:]...)
	}
	fmt.Println(sum)
}
