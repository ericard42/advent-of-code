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
	for i := 0; i < len(array1); i++ {
		howMany := 0
		for j := 0; j < len(array2); j++ {
			if array1[i] == array2[j] {
				howMany++
			}
		}
		sum += array1[i] * howMany

	}
	fmt.Println(sum)
}
