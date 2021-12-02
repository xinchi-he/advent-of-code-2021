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
	lines := read("input2.dat")

	//p1
	distance, depth := 0, 0

	for _, l := range lines {
		splits := strings.Split(l, " ")
		command, unit := splits[0], splits[1]
		fmt.Println(command)
		fmt.Println(unit)
		distance, depth = operations(command, unit, distance, depth)
	}

	fmt.Println(distance)
	fmt.Println(depth)

	fmt.Println(distance * depth)

	//p2
	distanceAim, depthAim, aim := 0, 0, 0

	for _, l := range lines {
		splits := strings.Split(l, " ")
		command, unit := splits[0], splits[1]
		distanceAim, depthAim, aim = operationsAim(command, unit, distanceAim, depthAim, aim)
	}

	fmt.Println(distanceAim * depthAim)

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

func operations(comand, unit string, distance, depth int) (int, int) {
	unit_int, _ := strconv.Atoi(unit)
	switch comand {
	case "forward":
		distance += unit_int
	case "up":
		depth -= unit_int
	case "down":
		depth += unit_int
	}
	return distance, depth
}

func operationsAim(command, unit string, distance, depth, aim int) (int, int, int) {
	unit_int, _ := strconv.Atoi(unit)
	switch command {
	case "down":
		aim += unit_int
	case "up":
		aim -= unit_int
	case "forward":
		distance += unit_int
		depth += aim * unit_int
	}
	return distance, depth, aim
}
