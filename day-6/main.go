package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input, startX, startY := readFile("input.txt")
	// fmt.Println(input)
	fmt.Println(startX)
	fmt.Println(startY)

	part1output := part1(input, startX, startY)
	fmt.Println(part1output)

	// part2output := part2(input)
	// fmt.Println(part2output)
}

func readFile(path string) ([][]bool, int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	var total [][]bool
	var initX, initY int
	x := 0
	for scanner.Scan() {
		curr := scanner.Text()
		var bs []bool
		for y, c := range curr {
			b := false
			if c == '#' {
				b = true
			} else if c == '^' {
				initX = x
				initY = y
			}
			bs = append(bs, b)
		}
		total = append(total, bs)
		x++

	}
	return total, initX, initY
}

func part1(input [][]bool, startX, startY int) int {
	xMax := len(input)
	yMax := len(input[0])

	var traversed [][]bool

	for _, currLine := range input {
		var traversedLine []bool
		for range currLine {
			traversedLine = append(traversedLine, false)
		}
		traversed = append(traversed, traversedLine)
	}

	// fmt.Println(traversed)

	// starting point
	traversed[startX][startY] = true

	dir := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	currX, currY := startX, startY
	keepMoving := true
	currDir := 0

	for keepMoving {
		dirX := dir[currDir%4][0]
		dirY := dir[currDir%4][1]
		// fmt.Println(dirX)
		// fmt.Println(dirY)

		if currX + dirX == xMax || currY + dirY == yMax || currX + dirX == -1 || currY + dirY == -1 {
			keepMoving = false
			break
		}
		if !input[currX+dirX][currY+dirY] {
			keepMoving = true
			currX = currX + dirX
			currY = currY + dirY
			fmt.Println(currX, currY)
			traversed[currX][currY] = true
		} else {
			currDir++
		}
	}

	fmt.Println(traversed)

	total := 0
	for _, currLine := range traversed {
		for _, i := range currLine {
			if i == true {
				total++
			}
		}
	}

	return total

}
