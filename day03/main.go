package main

import (
	"regexp"
	"fmt"
	"io/ioutil"
	"strconv"
)

func Solve(file string, part int) int {
	res := 0
	regex := *regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	strings := regex.FindAllStringSubmatch(file, -1)
	canMultiply := true
	for _, x := range strings {
		switch x[0] {
			case "do()":
				canMultiply = true
				break
			case "don't()":
				canMultiply = false
				break
			default:
				if part == 2 && !canMultiply {
					continue
				}
				opLeft, _ := strconv.Atoi(x[1])
				opRight, _ := strconv.Atoi(x[2])
				res += opLeft * opRight
		}
	}
	return res
}

func main() {
    content, _ := ioutil.ReadFile("input.txt")
	contentStr := string(content)
	fmt.Println(Solve(contentStr, 1))
	fmt.Println(Solve(contentStr, 2))
}
