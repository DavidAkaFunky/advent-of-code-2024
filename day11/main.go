package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func parseInput(filename string) map[string]int {
	file, _ := os.Open(filename)
    defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	numbers := make(map[string]int)
	for _, num := range strings.Fields(line) {
		numbers[num]++
	}
	return numbers
}

func getTotal(numbers map[string]int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}

func Solve(numbers map[string]int, blinks int) int {
	var newNumbers map[string]int
	for range(blinks) {
		newNumbers = make(map[string]int)
		for k, v := range numbers {
			if k == "0" {
				newNumbers["1"] += v
			} else if len(k) % 2 == 0 {
				numIntL, _ := strconv.Atoi(k[:len(k)/2])
				numIntR, _ := strconv.Atoi(k[len(k)/2:])
				newNumbers[strconv.Itoa(numIntL)] += v
				newNumbers[strconv.Itoa(numIntR)] += v
			} else {
				numInt, _ := strconv.Atoi(k)
				newNumbers[strconv.Itoa(numInt * 2024)] += v
			}
		}
		numbers = newNumbers
	}
	return getTotal(numbers)
}

func main() {
	numbers := parseInput("input.txt")
	fmt.Println(Solve(numbers, 25))
	fmt.Println(Solve(numbers, 75))
}