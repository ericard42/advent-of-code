package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseLine(line string, lineBefore string, lineAfter string) int {
	sum := 0

	re := regexp.MustCompile(`(\d+)`)
	numbersIndex := re.FindAllStringIndex(line, -1)

	for i := 0; i < len(numbersIndex); i++ {
		indexBefore := numbersIndex[i][0] - 1
		indexAfter := numbersIndex[i][1]
		if indexBefore < 0 {
			indexBefore = 0
		}
		if indexAfter >= len(line) {
			indexAfter = len(line) - 1
		}
		re := regexp.MustCompile(`[^\d.]`)
		testLine := re.MatchString(line[indexBefore : indexAfter+1])
		testLineBefore := false
		testLineAfter := false
		if lineBefore != "" {
			testLineBefore = re.MatchString(lineBefore[indexBefore : indexAfter+1])
		}
		if lineAfter != "" {
			testLineAfter = re.MatchString(lineAfter[indexBefore : indexAfter+1])
		}
		if testLine || testLineBefore || testLineAfter {
			numberToSum, err := strconv.Atoi(line[numbersIndex[i][0]:numbersIndex[i][1]])
			errorCheck(err)
			sum += numberToSum
		}
	}
	return sum
}

func main() {
	file, err := os.Open("./input.txt")
	errorCheck(err)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	res := 0
	lineBefore := ""
	line := ""
	lineAfter := ""
	for fileScanner.Scan() {
		if line != "" {
			lineAfter = fileScanner.Text()
			res += parseLine(line, lineBefore, lineAfter)
			lineBefore = line
		}
		line = fileScanner.Text()
	}
	lineAfter = ""
	res += parseLine(line, lineBefore, lineAfter)
	fmt.Println(res)
}
