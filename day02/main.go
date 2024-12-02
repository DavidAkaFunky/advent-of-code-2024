package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func parseInput(filename string) [][]int {
	file, _ := os.Open(filename)
    defer file.Close()
	scanner := bufio.NewScanner(file)
	res := [][]int{}
    for scanner.Scan() {
        line := scanner.Text()
		lineArray := []int{}
		for _, num := range strings.Fields(line) {
			numInt, _ := strconv.Atoi(num)
			lineArray = append(lineArray, numInt)
		}
		res = append(res, lineArray)
    }
	return res
}

func getSlopeSign(num1 int, num2 int) int {
	slope := num2 - num1
	switch {
		case slope > 0:
			return 1
		case slope == 0:
			return 0
		default:
			return -1
	}
}

func isOk(line []int) bool {
	slopeSign := getSlopeSign(line[0], line[1])
	for i := range len(line) - 1 {
		slope := slopeSign * (line[i+1] - line[i])
		if slope < 1 || slope > 3 {
			return false
		}
	}
	return true
}

func joinArrays(a []int, b []int) []int {
	la := len(a)
	c := make([]int, la, la + len(b))
	_ = copy(c, a)
	return append(c, b...)
}

func Part1(nums [][]int) int {
	res := 0
	for _, line := range nums {
		if isOk(line) {
			res += 1
		}
	}
	return res
}

func Part2(nums [][]int) int {
	res := 0
	for _, line := range nums {
		if isOk(line) {
			res += 1
			continue
		}
		for i := range line {
			if isOk(joinArrays(line[:i], line[i+1:])) {
				res += 1
				break
			}
		}
	}
	return res
}

func main() {
    nums := parseInput("input.txt")
	fmt.Println(Part1(nums))
	fmt.Println(Part2(nums))
}
