package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
)

func parseInput(filename string) (map[byte][][2]int, int, int) {
	file, _ := os.Open(filename)
    defer file.Close()
	scanner := bufio.NewScanner(file)
	
	regex := *regexp.MustCompile("[a-zA-Z\\d]")
	antennas := make(map[byte][][2]int)
	
	i := 0
	var maxJ int
	
    for scanner.Scan() {
        line := scanner.Text()
		maxJ = len(line)
		matches := regex.FindAllStringIndex(line, -1)
		for _, j := range matches {
			antennas[line[j[0]]] = append(antennas[line[j[0]]], [2]int{i, j[0]})
		}

		i += 1
    }

	return antennas, i, maxJ
}

func newPositions(pos1 [2]int, pos2 [2]int, maxI int, maxJ int, limit int) [][2]int {
	diffI := pos1[0] - pos2[0]
	diffJ := pos1[1] - pos2[1]

	positions := [][2]int{}

	newPosI := pos1[0]
	newPosJ := pos1[1]
	
	for ; limit != 0; limit-- {
		newPosI += diffI
		newPosJ += diffJ
		
		if newPosI < 0 || newPosI >= maxI || newPosJ < 0 || newPosJ >= maxJ {
			break
		}

		positions = append(positions, [2]int{newPosI, newPosJ})
	}

	return positions
}

func Solve(allAntennas map[byte][][2]int, maxI int, maxJ int, limit int) int {
	positions := make(map[[2]int]bool)

	if limit < 0 {
		for _, antennas := range allAntennas {
			for _, antenna := range antennas {
				positions[antenna] = true
			}
		}
	}

	for _, antennas := range allAntennas {
		for i, antennaI := range antennas {
			for j, antennaJ := range antennas {
				if i == j {
					continue
				}
				newPositions := newPositions(antennaI, antennaJ, maxI, maxJ, limit)
				for _, newPosition := range newPositions {
					positions[newPosition] = true
				}
			}
		}
	}

	return len(positions)
}

func main() {
	antennas, maxI, maxJ := parseInput("input.txt")
	fmt.Println(Solve(antennas, maxI, maxJ, 1))
	fmt.Println(Solve(antennas, maxI, maxJ, -1))
}