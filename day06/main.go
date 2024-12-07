package main

import (
	"regexp"
	"fmt"
	"os"
	"bufio"
	"strings"
)

func parseInput(filename string) ([]string, [2]int) {
	file, _ := os.Open(filename)
    defer file.Close()
	scanner := bufio.NewScanner(file)
	
	regexPos := *regexp.MustCompile("\\^")
	
	puzzle := []string{}
	var pos [2]int
	
	i := 0

    for scanner.Scan() {
        line := scanner.Text()
		puzzle = append(puzzle, strings.TrimSuffix(line, "\n"))

		indexPos := regexPos.FindStringIndex(line)
		if indexPos != nil {
			pos = [2]int{i, indexPos[0]}
		}

		i += 1
    }

	return puzzle, pos
}

func hasEscaped(puzzle []string, visitedPos map[[2]int]bool, pos [2]int, dimToChange int, dir int, extraWall [2]int) (bool, map[[2]int]bool, [2]int) {
	var maxLen int	
	if dimToChange == 0 {
		maxLen = len(puzzle)
	} else {
		maxLen = len(puzzle[0])
	}

	var oldPos [2]int

	for {
		oldPos = pos
		pos[dimToChange] += dir
		
		if pos[dimToChange] == maxLen || pos[dimToChange] == -1 {
			return true, visitedPos, oldPos
		} else if puzzle[pos[0]][pos[1]] == '#' || pos == extraWall {
			return false, visitedPos, oldPos
		}
			
		visitedPos[pos] = true
	}
}

func Part1(puzzle []string, pos [2]int) map[[2]int]bool {
	i := 0
	visitedPos := map[[2]int]bool{pos: true}
	noExtraWall := [2]int{-1, -1}

	var dimToChange int
	var dir int
	var escaped bool

	for {

		switch i {
			case 0:
				dimToChange, dir = 0, -1
				break
			case 1:
				dimToChange, dir = 1, 1
				break
			case 2:
				dimToChange, dir = 0, 1
				break
			case 3:
				dimToChange, dir = 1, -1
				break
		}

		escaped, visitedPos, pos = hasEscaped(puzzle, visitedPos, pos, dimToChange, dir, noExtraWall)
		
		if escaped {
			return visitedPos
		}

		i = (i + 1) % 4
	}
}

func Part2(puzzle []string, visitedPos map[[2]int]bool, originalPos [2]int) int {
	obstructions := 0

	var dimToChange int
	var dir int
	var escaped bool
	var i int
	var states map[[4]int]bool
	var pos [2]int
	emptyVisitedPos := map[[2]int]bool{}

	for p := range visitedPos {
		if p == originalPos {
			continue
		}

		states = map[[4]int]bool{}
		i = 0

		pos = originalPos
		for {

			switch i {
				case 0:
					dimToChange, dir = 0, -1
					break
				case 1:
					dimToChange, dir = 1, 1
					break
				case 2:
					dimToChange, dir = 0, 1
					break
				case 3:
					dimToChange, dir = 1, -1
					break
			}
	
			escaped, _, pos = hasEscaped(puzzle, emptyVisitedPos, pos, dimToChange, dir, p)

			if escaped {
				break
			}

			state := [4]int{pos[0], pos[1], dimToChange, dir}

			if states[state] {
				obstructions += 1
				break
			}
			states[state] = true

			i = (i + 1) % 4
		}
	}

	return obstructions
}

func main() {
	puzzle, pos := parseInput("input.txt")
	visitedPos := Part1(puzzle, pos)
	fmt.Println(len(visitedPos))
	fmt.Println(Part2(puzzle, visitedPos, pos))
}