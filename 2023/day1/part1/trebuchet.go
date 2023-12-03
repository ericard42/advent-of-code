package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseLine(line string) int {
	num := ""
	for i := 0; i < len(line); i++ {
		char := string(line[i])
		if char > "0" && char <= "9" {
			num += char
		}
	}
	if len(num) > 1 {
		num = string(num[0]) + string(num[len(num)-1])
	}
	if len(num) == 1 {
		num += num
	}
	if len(num) == 0 {
		num = "0"
	}
	intRet, err := strconv.Atoi(num)
	errorCheck(err)
	return intRet
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
