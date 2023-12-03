package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseLine(line string) int {
	index := strings.Index(line, ":")
	line = line[index+2:]
	minValue := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	re := regexp.MustCompile(`(\d+)\s(\w+)`)
	keyValue := re.FindAllStringSubmatch(line, -1)
	for i := 0; i < len(keyValue); i++ {
		value, err := strconv.Atoi(keyValue[i][1])
		errorCheck(err)
		if minValue[keyValue[i][2]] < value {
			minValue[keyValue[i][2]] = value
		}
	}
	return minValue["red"] * minValue["green"] * minValue["blue"]
}

func main() {
	file, err := os.Open("./list.txt")
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
