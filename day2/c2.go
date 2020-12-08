package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

type lineVal struct {
	n1   int
	n2   int
	char byte
	pass string
}

func main() {

	list, _ := readLines("input2.txt")
	reg := regexp.MustCompile(`(\d+)-(\d+) (.): (.*)`)

	var lineData lineVal
	var valid int
	var logic bool

	for _, line := range list {
		vals := reg.FindAllStringSubmatch(line, -1)

		lineData.n1, _ = strconv.Atoi(vals[0][1])
		lineData.n2, _ = strconv.Atoi(vals[0][2])
		lineData.char = vals[0][3][0]
		lineData.pass = vals[0][4]

		logic = (lineData.pass[lineData.n1-1] == lineData.char && lineData.pass[lineData.n2-1] != lineData.char) ||
			(lineData.pass[lineData.n1-1] != lineData.char && lineData.pass[lineData.n2-1] == lineData.char)

		if logic {
			valid++
		}
	}
	fmt.Println(valid)

}
