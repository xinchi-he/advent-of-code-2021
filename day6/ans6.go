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
	var initState []int
	for _, l := range read("input6.dat") {
		states := strings.Split(l, ",")

		for _, state := range states {
			intState, _ := strconv.Atoi(state)
			initState = append(initState, intState)
		}
	}

	fmt.Println(countFish(initState, 256))

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

func countFish(states []int, day int) int64 {
	/*

		day		0, 1, 2, 3, 4, 5, 6, 7, 8
		init	   1, 1, 2, 1
		day1	1, 1, 2, 1
		day2    1, 2, 1           1,    1
		day3    2, 1,          1, 1, 1, 1

		algo:  left shift every day, if day[0] != 0, shift first, and then append day[0] to the end,
				also increase day[6] by day[0]

	*/
	fishes := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for _, fish := range states {
		fishes[fish]++
	}

	fmt.Println(fishes)

	for i := 0; i < day; i++ {
		if fishes[0] == 0 {
			fishes = fishes[1:]
			fishes = append(fishes, 0)
		} else {
			top := fishes[0]
			fishes = fishes[1:]
			fishes = append(fishes, top)
			fishes[6] += top
		}
	}

	var res int64
	res = 0

	for _, f := range fishes {
		res += int64(f)
	}

	return res
}
