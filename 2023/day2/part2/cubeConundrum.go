package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	splitedLine := strings.Split(line, "; ")
	minValue := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for i := 0; i < len(splitedLine); i++ {
		tmpSplit := strings.Split(splitedLine[i], ", ")
		for j := 0; j < len(tmpSplit); j++ {
			keyValue := strings.Split(tmpSplit[j], " ")
			value, err := strconv.Atoi(keyValue[0])
			if minValue[keyValue[1]] < value {
				minValue[keyValue[1]] = value
			}
			errorCheck(err)
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
