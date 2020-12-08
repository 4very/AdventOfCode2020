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
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

type bags map[string]map[string]int
type subbag map[string]int

func main() {

	list, _ := readLines("input7.txt")
	var bags bags = make(bags)
	var name string
	nameReg := regexp.MustCompile("(.*) bags contain")
	emptyReg := regexp.MustCompile("no other bags.")
	contReg := regexp.MustCompile(`(\d+) (.*?) bags?`)

	for _, line := range list {

		name = nameReg.FindAllStringSubmatch(line, -1)[0][1]
		bags[name] = make(subbag)

		if emptyReg.MatchString(line) {
			continue
		} else {
			for _, elt := range contReg.FindAllStringSubmatch(line, -1) {
				numBags, _ := strconv.Atoi(elt[1])
				bags[name][elt[2]] = numBags
			}
		}
		// fmt.Println("-----")
		fmt.Println(bags[name])
	}

	var explored []string
	var good []string
	var exploring []string

	exploring = append(exploring, "shiny gold")
	for true {
		if len(exploring) == 0 { // if we explored everything
			break
		}
		// fmt.Println(exploring)

		exploreTemp := exploring
		exploring = exploring[:0]

		for _, explore := range exploreTemp {
			if _, found := Find(explored, explore); found {
				continue
			}

			for key, elt := range bags {
				if _, ok := elt[explore]; ok { // if bag current exploring can fit in bag
					exploring = append(exploring, key)
					good = append(good, key)
				}
			}
		}

		explored = append(explored, exploreTemp...)
	}

	fmt.Println(len(good))

	// for _, bag := range queue {
	// 	if _, ok := bags[bag]["shiny golden"]; ok { // if shiny golden is in these bags
	// 		good = append(good, bag)
	// 		queue = queue[1:]
	// 	} else if len(bags[bag]) == 0 { // if no bags cannot fit in it
	// 		bad = append(bad, bag)
	// 	} else { // if all bags are not golden
	// 		for subbag := range bags[bag] {
	// 			if _, found := Find(bad, subbag); found {

	// 			}
	// 			queue = append(queue, subbag)
	// 		}
	// 	}
	// }

}
