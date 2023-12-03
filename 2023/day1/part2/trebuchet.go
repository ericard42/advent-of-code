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

func parseLine(line string) int {
	mapNumbers := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	re := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	end := false
	var num []string
	for !end {
		ret := string(re.Find([]byte(line)))
		if ret != "" {
			num = append(num, ret)
			index := re.FindIndex([]byte(line))
			line = line[index[0]+1:]
		} else {
			end = true
		}
	}

	resString := []string{"", ""}
	res := ""
	if len(num) > 0 {
		resString[0] = num[0]
		resString[1] = num[len(num)-1]
		if len(resString[0]) > 1 {
			resString[0] = mapNumbers[resString[0]]
		}
		if len(resString[1]) > 1 {
			resString[1] = mapNumbers[resString[1]]
		}
		res = resString[0] + resString[1]
		intRet, err := strconv.Atoi(res)
		errorCheck(err)
		return intRet
	}
	return 0
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
