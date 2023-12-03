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

func checkBefore(line string, indexBefore int) int {
	numberToAdd := ""
	for indexBefore >= 0 && line[indexBefore] >= '0' && line[indexBefore] <= '9' {
		numberToAdd = string(line[indexBefore]) + numberToAdd
		indexBefore--
	}
	num, err := strconv.Atoi(numberToAdd)
	errorCheck(err)
	return num
}

func checkAfter(line string, indexAfter int) int {
	numberToAdd := ""
	for indexAfter < len(line) && line[indexAfter] >= '0' && line[indexAfter] <= '9' {
		numberToAdd += string(line[indexAfter])
		indexAfter++
	}
	num, err := strconv.Atoi(numberToAdd)
	errorCheck(err)
	return num
}

func checkBeforeAfter(line string, index int) int {
	numberToAdd := ""
	indexTmp := index
	for indexTmp >= 0 && line[indexTmp] >= '0' && line[indexTmp] <= '9' {
		numberToAdd = string(line[indexTmp]) + numberToAdd
		indexTmp--
	}
	index++
	for index < len(line) && line[index] >= '0' && line[index] <= '9' {
		numberToAdd += string(line[index])
		index++
	}
	num, err := strconv.Atoi(numberToAdd)
	errorCheck(err)
	return num
}

func parseLine(line string, lineBefore string, lineAfter string) int {
	sum := 0

	re := regexp.MustCompile(`\*`)
	gearsIndex := re.FindAllStringIndex(line, -1)

	for i := 0; i < len(gearsIndex); i++ {
		var numbers []int
		indexBefore := gearsIndex[i][0] - 1
		indexAfter := gearsIndex[i][1]
		if indexBefore < 0 {
			indexBefore = 0
		}
		if indexAfter >= len(line) {
			indexAfter = len(line) - 1
		}
		if line[indexBefore] >= '0' && line[indexBefore] <= '9' {
			numbers = append(numbers, checkBefore(line, indexBefore))
		}
		if line[indexAfter] >= '0' && line[indexAfter] <= '9' {
			numbers = append(numbers, checkAfter(line, indexAfter))
		}
		re := regexp.MustCompile(`\d+`)
		if lineBefore != "" {
			fragmentBefore := lineBefore[indexBefore : indexAfter+1]
			check := re.FindAllStringIndex(fragmentBefore, -1)
			if check != nil {
				for j := 0; j < len(check); j++ {
					numbers = append(numbers, checkBeforeAfter(lineBefore, check[j][0]+indexBefore))
				}
			}
		}
		if lineAfter != "" {
			fragmentAfter := lineAfter[indexBefore : indexAfter+1]
			check := re.FindAllStringIndex(fragmentAfter, -1)
			if check != nil {
				for j := 0; j < len(check); j++ {
					numbers = append(numbers, checkBeforeAfter(lineAfter, check[j][0]+indexBefore))
				}
			}
		}

		if len(numbers) >= 2 {
			sum += numbers[0] * numbers[1]
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
