package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := readFile("input.txt")

	totalPart1 := part1(input)
	fmt.Println(totalPart1)

	totalPart2 := part2(input)
	fmt.Println(totalPart2)

}

func readFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	total := []string{}
	for scanner.Scan() {
		curr := scanner.Text()
		total = append(total, curr)

	}
	return total
}

func part1(input []string) int {
	total := 0
	for i, line := range input {
		for j, _ := range line {
			total += checkIfWordExists(input, i, j)
		}
	}
	return total
}

func checkIfWordExists(input []string, i, j int) int {
	total := 0
	word := "XMAS"
	curr := input[i][j]
	if curr != word[0] {
		return 0
	}

	x := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	y := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for direction := 0; direction < 8; direction++ {
		currX := i + x[direction]
		currY := j + y[direction]
		k := 0

		for k = 1; k < len(word); k++ {
			if currX < 0 || currY < 0 || currX >= len(input) || currY >= len(input[0]) {
				break
			}
			if input[currX][currY] != word[k] {
				break
			}
			currX += x[direction]
			currY += y[direction]
		}
		if k == len(word) {
			total++
		}
	}
	return total
}

func part2(input []string) int {
	total := 0
	for i, line := range input {
		for j, _ := range line {
			if i == 0 || j == 0 || i >= len(input)-1 || j >= len(input)-1 {
				continue
			}
			if checkPattern(input, i, j) {
				total++
			}
		}
	}
	return total
}

func checkPattern(input []string, i, j int) bool {

	curr := input[i][j]
	if curr != 'A' {
		return false
	}

	diagCoords := [][]int{
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}

	if input[i+diagCoords[0][0]][j+diagCoords[0][1]] == 'M' {
		if input[i+diagCoords[1][0]][j+diagCoords[1][1]] == 'M' {
			if input[i+diagCoords[2][0]][j+diagCoords[2][1]] == 'S' {
				if input[i+diagCoords[3][0]][j+diagCoords[3][1]] == 'S' {
					fmt.Println(i, j)
					return true
				}
			}
		}
	}
	if input[i+diagCoords[0][0]][j+diagCoords[0][1]] == 'M' {
		if input[i+diagCoords[1][0]][j+diagCoords[1][1]] == 'S' {
			if input[i+diagCoords[2][0]][j+diagCoords[2][1]] == 'M' {
				if input[i+diagCoords[3][0]][j+diagCoords[3][1]] == 'S' {
					fmt.Println(i, j)

					return true
				}
			}
		}
	}
	if input[i+diagCoords[0][0]][j+diagCoords[0][1]] == 'S' {
		if input[i+diagCoords[1][0]][j+diagCoords[1][1]] == 'S' {
			if input[i+diagCoords[2][0]][j+diagCoords[2][1]] == 'M' {
				if input[i+diagCoords[3][0]][j+diagCoords[3][1]] == 'M' {
					fmt.Println(i, j)

					return true
				}
			}
		}
	}
	if input[i+diagCoords[0][0]][j+diagCoords[0][1]] == 'S' {
		if input[i+diagCoords[1][0]][j+diagCoords[1][1]] == 'M' {
			if input[i+diagCoords[2][0]][j+diagCoords[2][1]] == 'S' {
				if input[i+diagCoords[3][0]][j+diagCoords[3][1]] == 'M' {
					fmt.Println(i, j)

					return true
				}
			}
		}
	}
	return false

}
