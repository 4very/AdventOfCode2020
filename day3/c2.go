package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {

	list, _ := readLines("input3.txt")
	var map2 [][]rune
	var lineData []rune

	for _, line := range list {
		lineData = []rune{}
		for _, elt := range line {
			lineData = append(lineData, elt)
		}
		map2 = append(map2, lineData)
	}

	var vals []int
	var ds [][]int = [][]int{[]int{1, 1}, []int{3, 1}, []int{5, 1}, []int{7, 1}, []int{1, 2}}

	for _, delt := range ds {
		var xOF, yOF bool = false, false
		var trees int = 0
		var x, y int = 0, 0
		var dx, dy int = delt[0], delt[1]

		for true {

			x += dx
			y += dy

			if x >= len(map2[0])-1 {
				xOF = true
			}
			if y >= len(map2) {
				yOF = true
			}

			if x >= len(map2[0]) {
				x -= len(map2[0])
			}
			if y >= len(map2) {
				y -= len(map2)
			}

			if xOF && yOF {
				break
			}

			if map2[y][x] == 35 {
				trees++
			}

		}

		vals = append(vals, trees)
	}
	fmt.Println(vals)
	var ret int = 1
	for _, val := range vals {
		ret *= val
	}
	fmt.Println(ret)

}
