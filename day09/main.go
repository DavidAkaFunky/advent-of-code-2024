package main

import (
	"fmt"
	"os"
	"bufio"
)

func parseInput(filename string) [][2]int {
	file, _ := os.Open(filename)
    defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	id := 0
	files := [][2]int{}
	var size int
	for i, char := range line {
		size = int(char - '0')

		if i % 2 == 0 {
			files = append(files, [2]int{id, size})
			id += 1
		} else {
			files = append(files, [2]int{-1, size})
		}
	}

	return files
}

func Part1(files [][2]int) int {
	res := 0
	l, r := 0, len(files) - 1
	i := 0
	var leftId, leftSize, minSize int
	for {
		if l > r {
			break
		}

		leftId, leftSize = files[l][0], files[l][1]

		if leftId == -1 {
			for {
				minSize = min(leftSize, files[r][1])
				res += files[r][0] * minSize * (2*i + minSize - 1) / 2
				i += minSize
				leftSize -= minSize
				files[r][1] -= minSize
				if files[r][1] == 0 {
					r -= 2
					if l > r {
						break
					}
				}
				if leftSize == 0 {
					break
				}
			}
		} else {
			res += leftId * leftSize * (2*i + leftSize - 1) / 2
			i += leftSize
		}

		l++
	}

	return res
}

func Part2(files [][2]int) int {
	res := 0
	i := 0
	var r, leftId, leftSize int
	used := make(map[int]bool)
	for l, leftFile := range files {

		leftId, leftSize = leftFile[0], leftFile[1]

		if leftId == -1 {
			for r = len(files) - 1; r > l; r -= 2 {
				if !used[r] && files[r][1] <= leftSize {
					res += files[r][0] * files[r][1] * (2*i + files[r][1] - 1) / 2
					i += files[r][1]
					leftSize -= files[r][1]
					used[r] = true
				}
			}
		} else if !used[l] {
			res += leftId * leftSize * (2*i + leftSize - 1) / 2
		}
		
		i += leftSize
		l++
	}

	return res
}

func main() {
	files := parseInput("input.txt")
	filesCopy := make([][2]int, len(files))
	copy(filesCopy, files)
	fmt.Println(Part1(files))
	fmt.Println(Part2(filesCopy))
}