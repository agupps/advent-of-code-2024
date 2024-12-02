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
	set := readFile("input.txt")

	fmt.Println("part 1")
	part1Answer := part1(set)

	fmt.Println(part1Answer)

	fmt.Println("part 2")
	part2Answer := part2(set)
	fmt.Println(part2Answer)

}

func part1(input [][]int) int {
	safeLines := 0

	for _, line := range input {

		var increase, safe bool

		for i, curr := range line {
			if i == 0 {
				continue
			}

			prev := line[i-1]
			diff := curr - prev

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
			safeLines++
		}
	}
	return safeLines
}

func part2(input [][]int) int {
	var advancedOutput int
	for _, report := range input {
		if checkAdvanced(report) {
			advancedOutput++
		}
	}
	return advancedOutput
}

func checkBasic(list []int) bool {
	var increased int
	var decreased int
	for i := 0; i < (len(list) - 1); i++ {
		ss := list[i : i+2]
		distance := ss[0] - ss[1]
		if distance == 0 {
			return false
		}
		if distance < 0 {
			decreased++
			distance = distance * -1
		} else {
			increased++
		}
		if distance > 3 {
			return false
		}
	}
	if (increased != 0) && (decreased != 0) {
		return false
	}
	return true
}

func checkAdvanced(list []int) bool {
	if checkBasic(list) {
		return true
	}
	for i := 0; i < len(list); i++ {
		listCopy := make([]int, len(list))
		_ = copy(listCopy, list)
		adjustedList := append(listCopy[:i], listCopy[i+1:]...)

		if checkBasic(adjustedList) {
			return true
		}
	}
	return false
}

// for _, line := range lines {
// 	splitLine := strings.Split(line, " ")
// 	fmt.Println(splitLine)

// 	var increase, safe, usedDamper bool
// 	skipVal := -1

// 	for i, curr := range splitLine {
// 		fmt.Println(curr)
// 		if i == 0 {
// 			continue
// 		}
// 		prevIndex := i - 1
// 		if i-1 == skipVal {
// 			prevIndex = i - 2
// 		}

// 		current, _ := strconv.Atoi(curr)
// 		prev, _ := strconv.Atoi(splitLine[prevIndex])
// 		diff := current - prev

// 		if prevIndex == 0 {
// 			if diff > 0 {
// 				increase = true
// 			} else {
// 				increase = false
// 			}
// 		}

// 		if increase && diff > 0 && diff < 4 {
// 			safe = true
// 		} else if increase {
// 			if !usedDamper {
// 				fmt.Println("using")

// 				usedDamper = true
// 				skipVal = i
// 				safe = true
// 			} else {
// 				safe = false
// 				break
// 			}
// 		} else if !increase && diff < 0 && diff > -4 {
// 			safe = true
// 		} else {
// 			if !usedDamper {
// 				fmt.Println(diff)
// 				fmt.Println("using")

// 				usedDamper = true
// 				skipVal = i
// 				safe = true
// 			} else {
// 				safe = false
// 				break
// 			}
// 		}
// 	}
// 	if usedDamper {
// 		fmt.Println("used damp")
// 	}
// 	if safe {
// 		fmt.Println("valid")
// 		safeLines++
// 	}
// }
// fmt.Println(safeLines)

func readFile(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success")

	defer file.Close()
	scanner := bufio.NewScanner(file)
	var total [][]int
	for scanner.Scan() {
		currentLine := strings.Fields(scanner.Text())
		var curr []int
		for _, currStage := range currentLine {
			number, _ := strconv.Atoi(currStage)
			curr = append(curr, number)
		}

		total = append(total, curr)

	}
	return total

}


// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )





// func main() {

// 	var safeReportsCompensated int
// 	for _, report := range reports {
// 		if checkReportCompensated(report) {
// 			safeReportsCompensated++
// 		}
// 	}
// 	fmt.Println("Part 1, Safe reports: ", safeReports)
// 	fmt.Println("Part 2, Safe reports compensated: ", safeReportsCompensated)
// }
