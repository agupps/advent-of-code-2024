package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input := readFile("input.txt")
	fmt.Println(input)
	maxX := len(input)
	maxY := len(input[0])

	// loop and get all the unique antennas
	maps := getUniques(input)
	fmt.Println(maps)

	var antiNodesLocs []coords
	for _, coordList := range maps {
		antiNodesLocs = append(antiNodesLocs, getAntiNodesLocationsForAntennaSet(coordList, maxX, maxY)...)
	}

	fmt.Println(antiNodesLocs)

	lenX := len(input)
	lenY := len(input[0])

	var grid [][]bool

	for i := 0; i < lenX; i++ {
		grid = append(grid, make([]bool, lenY))
	}

	for _, curr := range antiNodesLocs {
		grid[curr.x][curr.y] = true
	}
	fmt.Println(grid)

	for _, coordinates := range maps {
		for _, coord := range coordinates {
			grid[coord.x][coord.y] = true
		}
	}

	total := 0
	for _, line := range grid {
		for _, b := range line {
			if b {
				total++
			}
		}
	}

	fmt.Println(total)

}

func readFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	var input []string

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curr := scanner.Text()
		input = append(input, curr)
	}
	return input
}

type coords struct {
	x, y int
}

func getUniques(input []string) map[rune][]coords {
	output := map[rune][]coords{}
	for x, line := range input {
		for y, curr := range line {
			if curr != '.' {
				val, ok := output[curr]
				currCoords := coords{x: x, y: y}
				if ok {
					output[curr] = append(val, currCoords)
				} else {
					output[curr] = []coords{currCoords}
				}
			}
		}
	}
	return output
}

func getAntiNodesLocationsForAntennaSet(antennaLocs []coords, maxX, maxY int) []coords {
	var antiNodeLocs []coords
	directions := [][]int{{1, 1}, {-1, 1}, {1, -1}, {-1, -1}}
	for _, antennaLoc1 := range antennaLocs {
		for _, antennaLoc2 := range antennaLocs {
			if antennaLoc1 == antennaLoc2 {
				continue
			}
			distanceMain := dist(antennaLoc1, antennaLoc2)

			for _, dir := range directions {
				antiNodeLoc := coords{}
				antiNodeLoc.x = antennaLoc1.x + dir[0]*(antennaLoc1.x-antennaLoc2.x)
				antiNodeLoc.y = antennaLoc1.y + dir[1]*(antennaLoc1.y-antennaLoc2.y)
				distance1To1 := dist(antiNodeLoc, antennaLoc1)
				distance1To2 := dist(antiNodeLoc, antennaLoc2)
				if distanceMain == distance1To1 && 4*distanceMain == distance1To2 {
					if isInBound(antiNodeLoc, maxX, maxY) {
						antiNodeLocs = append(antiNodeLocs, antiNodeLoc)
						antiNodeLocs = append(antiNodeLocs, getNextAntiNode(antiNodeLoc, antennaLoc1, maxX, maxY, dir)...)
					}
				}
			}
		}
	}
	return antiNodeLocs
}

func getNextAntiNode(antennaLoc1, antennaLoc2 coords, maxX, maxY int, dir []int) []coords {
	distanceMain := dist(antennaLoc1, antennaLoc2)
	var antiNodeLocs []coords
	antiNodeLoc := coords{}
	antiNodeLoc.x = antennaLoc1.x + dir[0]*(antennaLoc1.x-antennaLoc2.x)
	antiNodeLoc.y = antennaLoc1.y + dir[1]*(antennaLoc1.y-antennaLoc2.y)
	distance1To1 := dist(antiNodeLoc, antennaLoc1)
	distance1To2 := dist(antiNodeLoc, antennaLoc2)
	if distanceMain == distance1To1 && 4*distanceMain == distance1To2 {
		if isInBound(antiNodeLoc, maxX, maxY) {
			antiNodeLocs = append(antiNodeLocs, antiNodeLoc)
			antiNodeLocs = append(antiNodeLocs, getNextAntiNode(antiNodeLoc, antennaLoc1, maxX, maxY, dir)...)
		}
	}
	return antiNodeLocs
}

func isInBound(coord coords, maxX, maxY int) bool {
	return coord.x < maxX && coord.x >= 0 && coord.y < maxY && coord.y >= 0
}

func dist(a, b coords) int {
	return (a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)
}
