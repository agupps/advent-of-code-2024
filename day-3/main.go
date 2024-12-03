package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := readFile("input.txt")

	part1output := part1(input)
	fmt.Println(part1output)

	part2output := part2(input)
	fmt.Println(part2output)
}

func readFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	total := ""
	for scanner.Scan() {
		curr := scanner.Text()
		total += curr

	}
	return total
}

func part1(input string) int {
	re3 := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)")
	output := re3.FindAllString(input, -1)

	// fmt.Println(output)
	total := 0

	for _, curr := range output {
		split := strings.Split(curr, ",")
		// fmt.Println(split)
		first, _ := strconv.Atoi(split[0][4:])
		second, _ := strconv.Atoi(split[1][:len(split[1])-1])

		// fmt.Println(first)
		// fmt.Println(second)
		total += first * second
	}

	return total
}

func part2(input string) int {
	re3 := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)|do\\(\\)|don't\\(\\)")
	output := re3.FindAllString(input, -1)

	// fmt.Println(output)
	total := 0
	enabled := true

	for _, curr := range output {
		if curr == "do()" {
			enabled = true
		} else if curr == "don't()" {
			enabled = false
		} else if enabled {
			split := strings.Split(curr, ",")
			// fmt.Println(split)
			first, _ := strconv.Atoi(split[0][4:])
			second, _ := strconv.Atoi(split[1][:len(split[1])-1])

			// fmt.Println(first)
			// fmt.Println(second)
			total += first * second
		}
	}

	return total
}
