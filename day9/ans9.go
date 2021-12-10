package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	terrain := terrainBuilder(read("input9.dat"))

	//p1
	fmt.Println(findRiskLevel(terrain))

	//p2
	_, points := findRiskLevel(terrain)
	sizes := []int{}

	for _, p := range points {
		fmt.Println(findBasicSize(p, terrain))
		fmt.Println()
		sizes = append(sizes, findBasicSize(p, terrain))
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	fmt.Println("result", sizes[0]*sizes[1]*sizes[2])

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

func terrainBuilder(lines []string) [][]int {
	var terrain [][]int

	for _, l := range lines {
		var row []int
		for _, c := range string(l) {
			num, _ := strconv.Atoi(string(c))
			row = append(row, num)
		}

		terrain = append(terrain, row)
	}

	return terrain
}

func findRiskLevel(t [][]int) (int, [][]int) {
	risk := 0
	var lowPoints [][]int

	for i := 0; i < len(t); i++ {
		for j := 0; j < len(t[0]); j++ {
			c := 0
			adjacent := 4

			//up
			if i-1 >= 0 && i-1 <= len(t)-1 {
				if t[i][j] < t[i-1][j] {
					c++
				}
			} else {
				adjacent--
			}

			//down
			if i+1 >= 0 && i+1 <= len(t)-1 {
				if t[i][j] < t[i+1][j] {
					c++
				}
			} else {
				adjacent--
			}

			//left
			if j-1 >= 0 && j-1 <= len(t[0])-1 {
				if t[i][j] < t[i][j-1] {
					c++
				}
			} else {
				adjacent--
			}

			//right
			if j+1 >= 0 && j+1 <= len(t[0])-1 {
				if t[i][j] < t[i][j+1] {
					c++
				}
			} else {
				adjacent--
			}

			if c == adjacent {
				//fmt.Println(i, j, t[i][j])
				lowPoints = append(lowPoints, []int{i, j})
				risk += (1 + t[i][j])
			}

		}
	}

	return risk, lowPoints
}

func findBasicSize(p []int, t [][]int) int {
	//dfs
	var stack [][]int
	visited := make(map[string]bool)
	stack = append(stack, p)

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		i := current[0]
		j := current[1]

		key := strconv.Itoa(i) + strconv.Itoa(j)
		visited[key] = true

		fmt.Println("current", i, j)

		next := [][]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}

		fmt.Println("next", next)

		for _, n := range next {
			a := n[0]
			b := n[1]

			nextKey := strconv.Itoa(a) + strconv.Itoa(b)
			_, hasKey := visited[nextKey]

			if (a >= 0 && a <= len(t)-1) && (b >= 0 && b <= len(t[0])-1) {
				if t[a][b] != 9 && !hasKey {
					fmt.Println(a, b, "is part of the basin")
					stack = append(stack, n)
				}
			}

		}
	}

	fmt.Println("visited", visited)
	return len(visited)
}
