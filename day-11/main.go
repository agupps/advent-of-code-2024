package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// part 1
	input := readFile("input.txt")
	fmt.Println(input)

	ints := convertToInts(input)
	output := ints
	fmt.Println(output)
	for i := 0; i < 25; i++ {
		fmt.Println(i)
		output = applyRules(output)
	}
	fmt.Println(len(output))

	stones := map[int]int{}

	for _, val := range ints {
		stones[val]++
	}
	fmt.Println(stones)
	fmt.Println(secondPart(stones, 25))
	fmt.Println(secondPart(stones, 75))

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
		input = strings.Split(curr, " ")
	}
	return input
}

func convertToInts(input []string) []int {
	var output []int
	for _, str := range input {
		curr, _ := strconv.Atoi(str)
		output = append(output, curr)
	}
	return output
}

func applyRules(input []int) []int {
	var output []int
	for _, curr := range input {
		//fmt.Println("current number")
		//fmt.Println(curr)
		if curr == 0 {
			output = append(output, 1)
		} else if len(strconv.Itoa(curr))%2 == 0 {
			halfwayPoint := len(strconv.Itoa(curr)) / 2
			tens := math.Pow10(halfwayPoint)

			firstNum := curr / int(tens)
			output = append(output, firstNum)

			secondNum := curr % int(tens)
			output = append(output, secondNum)

		} else {
			output = append(output, curr*2024)
		}
	}
	return output
}

func secondPart(stones map[int]int, numberOfRuns int) int {
	for i := 0; i < numberOfRuns; i++ {
		nextStep := map[int]int{}
		for k, v := range stones {
			if k == 0 {
				nextStep[1] += v
			} else if len(strconv.Itoa(k))%2 == 0 {
				halfwayPoint := len(strconv.Itoa(k)) / 2
				tens := math.Pow10(halfwayPoint)
				firstNum := k / int(tens)
				nextStep[firstNum] += v

				secondNum := k % int(tens)
				nextStep[secondNum] += v
			} else {
				nextStep[k*2024] += v
			}
		}
		stones = nextStep
	}
	result := 0
	for _, v := range stones {
		result += v
	}
	return result
}
