package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var initState []int
	for _, l := range read("input6.dat") {
		states := strings.Split(l, ",")

		for _, state := range states {
			intState, _ := strconv.Atoi(state)
			initState = append(initState, intState)
		}
	}

	//fmt.Println(initState)

	/*

		DAY := 80

		for i := 0; i < DAY; i++ {
			for index, val := range initState {
				if val == 0 {
					initState[index] = 6
					initState = append(initState, 8)
				} else {
					initState[index]--
				}
			}

			//fmt.Println(i+1, initState)
		}

		fmt.Println(len(initState))
	*/

	dataSet := createParallelDataSet(8, initState)
	fmt.Println(dataSet, len(dataSet))

	for _, state := range dataSet {
		go countFish(state, 256)
	}

	time.Sleep(time.Second * 100)
	fmt.Println("done")

	fmt.Println("wtf?")

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

func createParallelDataSet(cpuCore int, dataSet []int) [][]int {
	var result [][]int
	var single []int
	for i, _ := range dataSet {
		if i == 0 {
			single = append(single, dataSet[i])
		} else if i%(len(dataSet)/cpuCore+1) != 0 {
			single = append(single, dataSet[i])
		} else {
			result = append(result, single)
			single = []int{}
		}
	}

	result = append(result, single)

	return result
}

func countFish(state []int, day int) int {
	for i := 0; i < day; i++ {
		for index, val := range state {
			if val == 0 {
				state[index] = 6
				state = append(state, 8)
			} else {
				state[index]--
			}
		}
	}

	fmt.Println(len(state))
	return len(state)
}
