package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("test")
	lines := readFile("input.txt")

	var first, second []int

	for _, line := range lines {
		split := strings.Split(line, "   ")
		fmt.Println(split)
		fmt.Println(len(split))

		firstInt, _ := strconv.Atoi(split[0])
		first = append(first, firstInt)
		secondInt, _ := strconv.Atoi(split[1])
		second = append(second, secondInt)

	}

	fmt.Println(first)
	fmt.Println(second)

	sort.Ints(first)
	sort.Ints(second)

	totalSum := 0.0

	for i, _ := range lines {
		totalSum += math.Abs(float64(first[i] - second[i]))
	}

	// answer to part 1
	fmt.Println(fmt.Sprintf("%f", totalSum))

	similarity := 0

	for _, firstNum := range first {
		count := 0
		for _, secondNum := range second {
			if firstNum == secondNum {
				count++
			}
		}
		similarity+=count*firstNum
	}

	fmt.Println(fmt.Sprintf("%d", similarity))


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
