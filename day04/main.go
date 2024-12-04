package main

import (
	"regexp"
	"fmt"
	"os"
	"bufio"
	"slices"
)

func parseInput(filename string) ([]string, [][2]int, [][2]int, int, int) {
	file, _ := os.Open(filename)
    defer file.Close()
	scanner := bufio.NewScanner(file)
	
	regexX := *regexp.MustCompile("X")
	regexA := *regexp.MustCompile("A")
	
	resX := [][2]int{}
	resA := [][2]int{}
	puzzle := []string{}
	
	maxX := -1
	i := 0
    for scanner.Scan() {
        line := scanner.Text()
		puzzle = append(puzzle, line)
		if maxX == -1 {
			maxX = len(line) - 1
		}

		indicesX := regexX.FindAllStringIndex(line, -1)
		for _, j := range indicesX {
			resX = append(resX, [2]int{i, j[0]})
		}

		indicesA := regexA.FindAllStringIndex(line, -1)
		for _, j := range indicesA {
			resA = append(resA, [2]int{i, j[0]})
		}

		i += 1
    }

	return puzzle, resX, resA, maxX, i - 1
}

func getValidNeighbours(puzzle []string, pos [2]int, maxY int, maxX int) string {
	res := ""

	for _, y := range [2]int{pos[0] - 1, pos[0] + 1} {
		if (y < 0 || y > maxY) {
			return ""
		}

		for _, x := range [2]int{pos[1] - 1, pos[1] + 1} {
			if (x < 0 || x > maxX) {
				return ""
			}
			res += string(puzzle[y][x])
		}
	}

	return res
}

func isValid(puzzle []string, pos [2]int, maxY int, maxX int, incY int, incX int,) int {
	y, x := pos[0], pos[1]
	for _, letter := range [3]byte{'M', 'A', 'S'} {
		y += incY
		x += incX
		if (x < 0 || x > maxX || y < 0 || y > maxY || puzzle[y][x] != letter) {
			return 0
		}
	}
	return 1
}

func Part1(puzzle []string, xPos [][2]int, maxY int, maxX int) int {
	res := 0
	for _, pos := range xPos {
		res += isValid(puzzle, pos, maxY, maxX, -1, -1) +
			   isValid(puzzle, pos, maxY, maxX, -1, 0) +
			   isValid(puzzle, pos, maxY, maxX, -1, 1) +
			   isValid(puzzle, pos, maxY, maxX, 0, -1) +
			   isValid(puzzle, pos, maxY, maxX, 0, 1) +
			   isValid(puzzle, pos, maxY, maxX, 1, -1) +
			   isValid(puzzle, pos, maxY, maxX, 1, 0) +
			   isValid(puzzle, pos, maxY, maxX, 1, 1)
	}
	return res
}

func Part2(puzzle []string, aPos [][2]int, maxY int, maxX int) int {
	res := 0
	valid := []string{"MMSS", "SMSM", "SSMM", "MSMS"}
	for _, pos := range aPos {
		if slices.Contains(valid, getValidNeighbours(puzzle, pos, maxY, maxX)) {
			res += 1
		}
	}
	return res
}

func main() {
	puzzle, xPos, aPos, maxY, maxX := parseInput("input.txt")
	fmt.Println(Part1(puzzle, xPos, maxY, maxX))
	fmt.Println(Part2(puzzle, aPos, maxY, maxX))
}