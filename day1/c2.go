package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var input = `1721
979
366
299
675
1456`

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

func main() {

	list, _ := readLines("input1.txt")

	for _, i := range list {
		i2, _ := strconv.Atoi(i)
		for _, j := range list {
			j2, _ := strconv.Atoi(j)
			for _, l := range list {
				l2, _ := strconv.Atoi(l)
				if i2+j2+l2 == 2020 {
					fmt.Println(i2, j2, l2, i2*j2*l2)
				}
			}

		}
	}
}
