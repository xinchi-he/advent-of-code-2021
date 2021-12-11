package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var positions []int

	for _, l := range read("input7.dat") {
		nums := strings.Split(l, ",")
		for _, num := range nums {
			n, _ := strconv.Atoi(num)
			positions = append(positions, n)
		}
	}

	//p1
	fmt.Println(leastFuel(positions))

	//p2
	fmt.Println(leastFuelFactorial(positions))

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

func leastFuel(positions []int) int {
	p := make(map[int]int)

	for _, pos := range positions {
		_, hasKey := p[pos]

		if hasKey {
			p[pos]++
		} else {
			p[pos] = 1
		}
	}

	//fmt.Println(p)

	var fuel []int

	for position := range p {
		currentFuel := 0
		for key, value := range p {
			if key > position {
				currentFuel += (key - position) * value
			} else {
				currentFuel += (position - key) * value
			}

			//fmt.Println("move from", key, " to ", position, " current fuel: ", currentFuel)
		}

		fuel = append(fuel, currentFuel)
	}

	sort.Slice(fuel, func(i, j int) bool {
		return fuel[i] < fuel[j]
	})

	//fmt.Println(fuel)

	return fuel[0]
}

func leastFuelFactorial(positions []int) int64 {
	p := make(map[int]int)
	factDict := make(map[int]int64)

	for _, pos := range positions {
		_, hasKey := p[pos]

		if hasKey {
			p[pos]++
		} else {
			p[pos] = 1
		}
	}

	//fmt.Println(p)

	var fuel []int64

	max := -1

	for key := range p {
		if key > max {
			max = key
		}
	}

	for position := 0; position <= max; position++ {
		currentFuel := int64(0)
		for key, value := range p {
			if key > position {
				currentFuel += factorial(key-position, factDict) * int64(value)
			} else {
				currentFuel += factorial(position-key, factDict) * int64(value)
			}

			//fmt.Println("move from", key, " to ", position, " current fuel: ", currentFuel)
		}

		fuel = append(fuel, currentFuel)
	}

	sort.Slice(fuel, func(i, j int) bool {
		return fuel[i] < fuel[j]
	})

	//fmt.Println(fuel)

	return fuel[0]
}

func factorial(n int, d map[int]int64) int64 {
	_, hasKey := d[n]

	if hasKey {
		return d[n]
	} else {
		f := int64(0)

		for i := 1; i <= n; i++ {
			f += int64(i)
		}

		d[n] = f
		return f
	}
}
