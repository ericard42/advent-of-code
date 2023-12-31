package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func parseLine(line string) int {
	index := strings.Index(line, ":")
	line = line[index+2:]
	index = strings.Index(line, "|")
	winners := strings.Split(line[:index], " ")
	checks := line[index+2:]

	total := 0
	for _, winner := range winners {
		if winner == "" {
			continue
		}
		re := regexp.MustCompile(` ` + winner + ` | ` + winner + `$|^` + winner + ` `)
		matches := re.MatchString(checks)
		if matches {
			if total == 0 {
				total = 1
			} else {
				total *= 2
			}
		}
	}
	return total
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
