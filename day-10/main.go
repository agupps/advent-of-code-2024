package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coord struct {
	x, y int
}

func main() {
	input := readFile("input.txt")
	fmt.Println(input)

	part1answer := part1(input)
	fmt.Println(part1answer)

	part2answer := part2(input)
	fmt.Println(part2answer)
}

func readFile(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	var input [][]int
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curr := scanner.Text()
		var line []int
		for _, str := range curr {
			line = append(line, int(str-48))
		}
		input = append(input, line)
	}
	return input
}

func part1(input [][]int) int {
	result := 0
	seen := map[coord][]coord{}
	for i, line := range input {
		for j, _ := range line {
			seen = numTrailHeads(input, 0, i, j, seen, i, j)
		}
	}
	for _, v := range seen {
		result += len(v)
	}
	return result
}

func numTrailHeads(input [][]int, num, i, j int, seen map[coord][]coord, originalX, originalY int) map[coord][]coord {
	curr := input[i][j]
	//fmt.Println(curr)
	if curr != num {
		return seen
	}
	if curr == 9 {
		OG := coord{x: originalX, y: originalY}
		final := coord{x: i, y: j}
		coords, ok := seen[OG]
		if !ok {
			seen[OG] = []coord{final}
		} else {
			for _, currentCoord := range coords {
				if currentCoord == final {
					return seen
				}
			}
			seen[OG] = append(coords, final)
		}
		return seen
	}
	dirs := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	xMax := len(input)
	yMax := len(input[0])
	for _, dir := range dirs {
		newX := i + dir[0]
		newY := j + dir[1]
		if newX >= 0 && newX < xMax && newY >= 0 && newY < yMax {
			//fmt.Println(newX, " ", newY)
			seen = numTrailHeads(input, num+1, newX, newY, seen, originalX, originalY)
		}
	}
	return seen
}

func part2(input [][]int) int {
	result := 0
	for i, line := range input {
		for j, _ := range line {
			result += numTrailHeads2(input, 0, i, j)
		}
	}

	return result
}

func numTrailHeads2(input [][]int, num, i, j int) int {
	curr := input[i][j]
	//fmt.Println(curr)
	if curr != num {
		return 0
	}
	if curr == 9 {
		return 1
	}
	dirs := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	result := 0
	xMax := len(input)
	yMax := len(input[0])
	for _, dir := range dirs {
		newX := i + dir[0]
		newY := j + dir[1]
		if newX >= 0 && newX < xMax && newY >= 0 && newY < yMax {
			//fmt.Println(newX, " ", newY)
			result += numTrailHeads2(input, num+1, newX, newY)
		}
	}
	return result
}
