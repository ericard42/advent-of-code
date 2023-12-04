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

var addCards []int

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
			total++
		}
	}
	nowAdd := 0
	if len(addCards) > 0 {
		nowAdd = addCards[0]
		addCards = addCards[1:]
	}
	for i := 0; i < total; i++ {
		if i >= len(addCards) {
			addCards = append(addCards, nowAdd+1)
		} else {
			addCards[i] = addCards[i] + nowAdd + 1
		}
	}

	return nowAdd + 1
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
