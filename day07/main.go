package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type stack [][2]int

func (s stack) Push(v [2]int) stack {
    return append(s, v)
}

func (s stack) Pop() (stack, [2]int) {
    l := len(s)
    return s[:l-1], s[l-1]
}

type line struct {
    result int
	numbers []int
}

func parseInput(filename string) []line {
	file, _ := os.Open(filename)
    defer file.Close()
	scanner := bufio.NewScanner(file)
	res := []line{}

    for scanner.Scan() {

        text := scanner.Text()
		var result int
		numbers := []int{}

		for i, num := range strings.Fields(text) {
			if i == 0 {
				result, _ = strconv.Atoi(num[:len(num)-1])
			} else {
				numInt, _ := strconv.Atoi(num)
				numbers = append(numbers, numInt)
			}
		}

		res = append(res, line{result: result, numbers: numbers})
    }

	return res
}

func concat(n1 int, n2 int) int {
	numInt, _ := strconv.Atoi(strconv.Itoa(n1) + strconv.Itoa(n2))
	return numInt
}

func Solve(lines []line, part int) int {
	res := 0
	var s stack
	var el [2]int

	for _, l := range(lines) {
		
		result := l.result
		numbers := l.numbers
		s = stack{[2]int{0, numbers[0]}}

		for {

			if len(s) == 0 {
				break
			}
			
			s, el = s.Pop()
			i, number := el[0], el[1]
			i += 1
			sum := number + numbers[i]
			product := number * numbers[i]
			conc := concat(number, numbers[i])

			if i == len(numbers) - 1 {
				if sum == result || product == result || (part == 2 && conc == result) {
					res += result
					break
				}
				continue
			}

			if sum <= result {
				s = s.Push([2]int{i, sum})
			}

			if product <= result {
				s = s.Push([2]int{i, product})
			}

			if part == 2 && conc <= result {
				s = s.Push([2]int{i, conc})
			}
		}
	}

	return res
}

func main() {
	lines := parseInput("input.txt")
	fmt.Println(Solve(lines, 1))
	fmt.Println(Solve(lines, 2))
}