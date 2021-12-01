package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//p1
	lines := read("input1.dat")

	var intArray []int

	for _, value := range lines {
		i, _ := strconv.Atoi(value)
		intArray = append(intArray, i)
	}

	counter := 0

	for i := 1; i < len(intArray); i++ {
		if intArray[i] > intArray[i-1] {
			counter++
		}
	}

	fmt.Println(counter)

	//p2

	var threeSumWindow []int

	for i := 0; i < len(intArray)-2; i++ {
		threeSumWindow = append(threeSumWindow, intArray[i]+intArray[i+1]+intArray[i+2])
	}

	partTwoCounter := 0

	for i := 1; i < len(threeSumWindow); i++ {
		if threeSumWindow[i] > threeSumWindow[i-1] {
			partTwoCounter++
		}
	}

	fmt.Println(partTwoCounter)

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
