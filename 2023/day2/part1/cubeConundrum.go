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
	goodValues := map[string]int{
		"red":   12,
		"blue":  14,
		"green": 13,
	}
	index := strings.Index(line, ":")
	numberOfGame := line[:index]
	line = line[index+2:]
	re := regexp.MustCompile(`(\d+)\s(\w+)`)
	keyValue := re.FindAllStringSubmatch(line, -1)
	for i := 0; i < len(keyValue); i++ {
		value, err := strconv.Atoi(keyValue[i][1])
		errorCheck(err)
		if goodValues[keyValue[i][2]] < value {
			return 0
		}
	}
	ret, err := strconv.Atoi(numberOfGame[5:])
	errorCheck(err)
	return ret
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
