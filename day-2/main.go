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
	fmt.Println("test")

	lines := readFile("input.txt")

	safeLines := 0

	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		fmt.Println(splitLine)

		var increase, safe bool

		for i, curr := range splitLine {
			if i == 0 {
				continue
			}

			current, _ := strconv.Atoi(curr)
			prev, _ := strconv.Atoi(splitLine[i-1])
			diff := current - prev

			if i == 1 {
				if diff > 0 {
					increase = true
				}
			}

			if increase && diff > 0 && diff < 4 {
				safe = true
			} else if increase {
				safe = false
				break
			} else if !increase && diff < 0 && diff > -4 {
				safe = true
			} else {
				safe = false
				break
			}
		}
		if safe {
			fmt.Println("valid")
			safeLines++
		}
	}
	fmt.Println(safeLines)

}

func readFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success")

	defer file.Close()
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
	return lines

}
