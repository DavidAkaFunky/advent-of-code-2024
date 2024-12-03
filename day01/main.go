package main

import (
	"regexp"
	"fmt"
	"io/ioutil"
	"strconv"
	"sort"
	"math"
)

func getLists(file string) ([]int, []int) {
	regex := *regexp.MustCompile(`(?s)(\d+)   (\d+)`)
	res := regex.FindAllStringSubmatch(file, -1)
	left := make([]int, len(res))
	right := make([]int, len(res))
	for _, x := range res {
		iLeft, _ := strconv.Atoi(x[1])
		iRight, _ := strconv.Atoi(x[2])
		left = append(left, iLeft)
		right = append(right, iRight)
	}
	sort.Sort(sort.IntSlice(left))
	sort.Sort(sort.IntSlice(right))
	return left, right
}

func Part1(left []int, right []int) int {
	res := 0
	for i := range len(left) {
        res += int(math.Abs(float64(left[i] - right[i])))
    }
	return res
}

func Part2(left []int, right []int) int {
	counter := make(map[int]int)
	res := 0
	for _, iRight := range right {
        counter[iRight] += 1
    }
	for _, iLeft := range left {
        res += counter[iLeft] * iLeft
    } 
	return res
}

func main() {
    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        fmt.Println("Err")
    }
    left, right := getLists(string(content))
	fmt.Println(Part1(left, right))
	fmt.Println(Part2(left, right))
}
