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

func parseLine(line string) int {
	return 0
}

func main() {
	file, err := os.Open("./input.txt")
	errorCheck(err)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	res := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		res += parseLine(line)
	}
	fmt.Println(res)
}
