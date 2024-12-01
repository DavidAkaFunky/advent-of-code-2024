package main

import (
	"regexp"
	"fmt"
	"io/ioutil"
	"strconv"
	"sort"
	"math"
)

func get_lists(file string) ([]int, []int) {
	left := []int{}
	right := []int{}
	regex := *regexp.MustCompile(`(?s)(\d+)   (\d+)`)
	res := regex.FindAllStringSubmatch(file, -1)
	for i := range res {
		left_i, _ := strconv.Atoi(res[i][1])
		right_i, _ := strconv.Atoi(res[i][2])
		left = append(left, left_i)
		right = append(right, right_i)
	}
	sort.Sort(sort.IntSlice(left))
	sort.Sort(sort.IntSlice(right))
	return left, right
}

func part1(left []int, right []int) int {
	res := 0
	for i := 0; i < len(left); i++ {
        res += int(math.Abs(float64(left[i] - right[i])))
    }
	return res
}

func part2(left []int, right []int) int {
	counter := make(map[int]int)
	res := 0
	for i := 0; i < len(right); i++ {
        counter[right[i]] += 1
    }
	for i := 0; i < len(left); i++ {
        res += counter[left[i]] * left[i]
    } 
	return res
}

func main() {
    content, err := ioutil.ReadFile("input.txt")
    if err != nil {
        fmt.Println("Err")
    }
    left, right := get_lists(string(content))
	fmt.Println(part1(left, right))
	fmt.Println(part2(left, right))
}
