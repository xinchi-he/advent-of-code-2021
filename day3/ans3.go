package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := matrixBuilder(read("input3.dat"))

	//p1
	gammarRate := ""

	for i := 0; i < len(input[0]); i++ {
		zeroCounter, oneCounter := 0, 0

		for j := 0; j < len(input); j++ {
			if input[j][i] == 0 {
				zeroCounter++
			} else if input[j][i] == 1 {
				oneCounter++
			}
		}

		if zeroCounter > oneCounter {
			gammarRate += "0"
		} else {
			gammarRate += "1"
		}
	}

	epsilonRate := ""

	for _, c := range gammarRate {
		if string(c) == "0" {
			epsilonRate += "1"
		} else {
			epsilonRate += "0"
		}
	}

	gammarRateDecimal, _ := strconv.ParseInt(gammarRate, 2, 64)
	epsilonRateDecimal, _ := strconv.ParseInt(epsilonRate, 2, 64)
	res := gammarRateDecimal * epsilonRateDecimal

	//fmt.Println(gammarRate)
	//fmt.Println(epsilonRate)
	fmt.Println(res)

	//p2
	oxygenMatrix, co2Matrix := input, input

	oxyPtr := 0

	for len(oxygenMatrix) > 1 {
		zeroCounter, oneCounter := 0, 0

		for i := 0; i < len(oxygenMatrix); i++ {
			if oxygenMatrix[i][oxyPtr] == 0 {
				zeroCounter++
			} else if oxygenMatrix[i][oxyPtr] == 1 {
				oneCounter++
			}
		}

		var rowsToRetain []int

		if zeroCounter > oneCounter {
			for i := 0; i < len(oxygenMatrix); i++ {
				if oxygenMatrix[i][oxyPtr] == 0 {
					rowsToRetain = append(rowsToRetain, i)
				}
			}
		} else if oneCounter > zeroCounter {
			for i := 0; i < len(oxygenMatrix); i++ {
				if oxygenMatrix[i][oxyPtr] == 1 {
					rowsToRetain = append(rowsToRetain, i)
				}
			}
		} else if zeroCounter == oneCounter {
			for i := 0; i < len(oxygenMatrix); i++ {
				if oxygenMatrix[i][oxyPtr] == 1 {
					rowsToRetain = append(rowsToRetain, i)
				}
			}
		}

		dict := make(map[int][]int)

		for i := 0; i < len(oxygenMatrix); i++ {
			dict[i] = oxygenMatrix[i]
		}

		var temp [][]int

		for _, row := range rowsToRetain {
			temp = append(temp, dict[row])
		}

		oxygenMatrix = temp
		//fmt.Println(oxygenMatrix)

		//fmt.Println(zeroCounter, oneCounter)
		//fmt.Println(len(oxygenMatrix))
		//fmt.Println(rowsToRetain)

		oxyPtr++
	}

	co2Ptr := 0

	for len(co2Matrix) > 1 {
		zeroCounter, oneCounter := 0, 0

		for i := 0; i < len(co2Matrix); i++ {
			if co2Matrix[i][co2Ptr] == 0 {
				zeroCounter++
			} else if co2Matrix[i][co2Ptr] == 1 {
				oneCounter++
			}
		}

		var rowsToRetain []int

		if zeroCounter > oneCounter {
			for i := 0; i < len(co2Matrix); i++ {
				if co2Matrix[i][co2Ptr] == 1 {
					rowsToRetain = append(rowsToRetain, i)
				}
			}
		} else if zeroCounter < oneCounter {
			for i := 0; i < len(co2Matrix); i++ {
				if co2Matrix[i][co2Ptr] == 0 {
					rowsToRetain = append(rowsToRetain, i)
				}
			}
		} else if zeroCounter == oneCounter {
			for i := 0; i < len(co2Matrix); i++ {
				if co2Matrix[i][co2Ptr] == 0 {
					rowsToRetain = append(rowsToRetain, i)
				}
			}
		}

		dict := make(map[int][]int)

		for i := 0; i < len(co2Matrix); i++ {
			dict[i] = co2Matrix[i]
		}

		var temp [][]int

		for _, row := range rowsToRetain {
			temp = append(temp, dict[row])
		}

		co2Matrix = temp
		fmt.Println(co2Matrix)
		fmt.Println(rowsToRetain)
		fmt.Println(co2Ptr)
		fmt.Println(zeroCounter, oneCounter)

		co2Ptr++
	}

	fmt.Println(oxygenMatrix)
	fmt.Println(co2Matrix)

	oxygenStr, co2Str := "", ""

	for _, d := range oxygenMatrix[0] {
		oxygenStr += strconv.Itoa(d)
	}

	for _, d := range co2Matrix[0] {
		co2Str += strconv.Itoa(d)
	}

	oxygenDecimal, _ := strconv.ParseInt(oxygenStr, 2, 64)
	co2Decimal, _ := strconv.ParseInt(co2Str, 2, 64)
	fmt.Println(oxygenDecimal * co2Decimal)

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

func matrixBuilder(lines []string) [][]int {
	var matrix [][]int

	for _, l := range lines {
		var singleRow []int
		for _, c := range l {
			digit, _ := strconv.Atoi(string(c))
			singleRow = append(singleRow, digit)
		}
		matrix = append(matrix, singleRow)
	}
	return matrix
}
