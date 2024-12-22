package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func test(rules map[string]map[string]bool, update []string) bool {
	for i, num1 := range update {
		for j := i + 1; j < len(update); j++ {
			num2 := update[j]
			if !rules[num1][num2] {
				return false
			}
		}
	}
	return true
}

func parseInput(filename string) (map[string]map[string]bool, [][]string) {

	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rules := make(map[string]map[string]bool)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		result := strings.Split(line, "|")
		num1, num2 := result[0], result[1]

		if _, ok := rules[num1]; !ok {
			rules[num1] = make(map[string]bool)
		}

		rules[num1][num2] = true
	}

	updates := [][]string{}

	for scanner.Scan() {
		updates = append(updates, strings.Split(scanner.Text(), ","))
	}

	return rules, updates
}

func Part1(rules map[string]map[string]bool, updates [][]string) int {

	res := 0

	for _, update := range updates {
		
		if test(rules, update) {
			value, _ := strconv.Atoi(update[(len(update) - 1) / 2])
			res += value
		}
	}

	return res
}

func Part2(rules map[string]map[string]bool, updates [][]string) int {

	res := 0
	var valid bool
	var swapped bool

	for _, update := range updates {
		
		valid = true
		midPoint := (len(update) - 1) / 2
		
		for n := len(update); n > midPoint; n-- {
			swapped = false

			for i := 1; i < n; i++ {
				num1, num2 := update[i-1], update[i]
				if !rules[num1][num2] {
					valid = false
					swapped = true
					update[i-1], update[i] = num2, num1
				}
			}
			
			if !swapped {
				break
			}
		}
		
		if !valid {
			fmt.Println(update)
			value, _ := strconv.Atoi(update[midPoint])
			res += value
		}
	}

	return res
}

func main() {
   	rules, updates := parseInput("input.txt")
	fmt.Println(Part1(rules, updates))
	fmt.Println(Part2(rules, updates))
}