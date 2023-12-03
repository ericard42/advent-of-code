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
	goodValues := map[string]int{
		"red":   12,
		"blue":  14,
		"green": 13,
	}
	index := strings.Index(line, ":")
	numberOfGame := line[:index]
	line = line[index+2:]
	splitedLine := strings.Split(line, "; ")
	for i := 0; i < len(splitedLine); i++ {
		tmpSplit := strings.Split(splitedLine[i], ", ")
		for j := 0; j < len(tmpSplit); j++ {
			keyValue := strings.Split(tmpSplit[j], " ")
			value, err := strconv.Atoi(keyValue[0])
			errorCheck(err)
			if goodValues[keyValue[1]] < value {
				return 0
			}
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
