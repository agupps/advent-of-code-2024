package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	firstPart, secondPart := readFile("day-5/input-sample.txt")

	part1result := part1(firstPart, secondPart)
	fmt.Println(part1result)

}

func readFile(path string) ([][]int, [][]int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	var secondPart [][]int
	var firstPart [][]int
	hitNewLine := false
	for scanner.Scan() {
		curr := scanner.Text()

		if curr == "" {
			hitNewLine = true
			continue
		}
		if !hitNewLine {
			split := strings.Split(curr, "|")
			var intermediate []int

			for _, s := range split {
				in, _ := strconv.Atoi(s)
				intermediate = append(intermediate, in)
			}
			firstPart = append(firstPart, intermediate)
		} else {
			split := strings.Split(curr, ",")
			var intermediate []int
			for _, s := range split {
				in, _ := strconv.Atoi(s)
				intermediate = append(intermediate, in)
			}
			secondPart = append(secondPart, intermediate)
		}
	}
	return firstPart, secondPart
}

func part1(rules [][]int, input [][]int) int {
	var validOnes [][]int
	for _, list := range input {
		fmt.Println(list)
		isValid := false
		for i, num := range list {
			if i == len(list)-1 {
				validOnes = append(validOnes, list)
			} else {
				for _, rule := range rules {
					if num == rule[0] && list[i+1] == rule[1] {
						isValid = true
						break
					}
				}

				if !isValid {
					break
				}
			}
		}
	}
	fmt.Println(validOnes)
	return 0
}
