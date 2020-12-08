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

func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func findint(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func main() {

	list, _ := readLines("input8.txt")
	var visited []int
	var current int = 0
	var accum int = 0
	var instruction string = ""
	var number int

	for true {
		// if we've visited somewhere that's already been visted
		if _, found := findint(visited, current); found {
			break
		}

		visited = append(visited, current)
		if current > len(list)-1 || current < 0 {
			fmt.Println("Theres been an error")
			break
		}

		line := list[current]
		lineRegex := regexp.MustCompile("(.*) (.*)")
		lineParse := lineRegex.FindAllStringSubmatch(line, -1)

		instruction = lineParse[0][1]
		number, _ = strconv.Atoi(lineParse[0][2])
		switch instruction {
		case "nop":
			current++
		case "acc":
			accum = accum + number
			current++
		case "jmp":
			current = current + number
		}
	}
	fmt.Println(accum)
}
