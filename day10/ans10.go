package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	/*
		//p1
		score := 0
		for _, l := range read("input10.dat") {
			fmt.Println(l)
			score, _ = checkSyntax(l, score)
		}

		fmt.Println(score)
	*/

	//p2
	score := 0
	var compScores []int
	for _, l := range read("input10.dat") {
		_, cs := checkSyntax(l, score)
		if cs != 0 {
			compScores = append(compScores, cs)
		}
	}

	sort.Slice(compScores, func(i, j int) bool {
		return compScores[i] > compScores[j]
	})

	fmt.Println(compScores[len(compScores)/2])

}

func read(path string) []string {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func checkSyntax(input string, score int) (int, int) {
	var stack []string

	scores := make(map[string]int)
	scores[")"] = 3
	scores["]"] = 57
	scores["}"] = 1197
	scores[">"] = 25137

	corrupted := false

	for _, c := range input {
		if string(c) == "[" || string(c) == "{" || string(c) == "(" || string(c) == "<" {
			stack = append(stack, string(c))
		} else if string(c) == "]" || string(c) == "}" || string(c) == ")" || string(c) == ">" {
			current := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			switch current {
			case "[":
				if string(c) != "]" {
					fmt.Println("current", current, "expected ], but found ", string(c))
					score += scores[string(c)]
					corrupted = true
				}
			case "{":
				if string(c) != "}" {
					fmt.Println("current", current, "expected }, but found ", string(c))
					score += scores[string(c)]
					corrupted = true
				}
			case "(":
				if string(c) != ")" {
					fmt.Println("current", current, "expected ), but found ", string(c))
					score += scores[string(c)]
					corrupted = true
				}
			case "<":
				if string(c) != ">" {
					fmt.Println("current", current, "expected >, but found ", string(c))
					score += scores[string(c)]
					corrupted = true
				}
			}

		}

	}

	closeScores := make(map[string]int)
	closeScores["("] = 1
	closeScores["["] = 2
	closeScores["{"] = 3
	closeScores["<"] = 4

	autoCompleteScore := 0

	if len(stack) > 0 && !corrupted {
		fmt.Println("still have unprocessed symbols left in the stack, got missing symbols")
		for len(stack) > 0 {
			current := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			autoCompleteScore = (autoCompleteScore * 5) + closeScores[current]
		}
	}

	return score, autoCompleteScore
}
