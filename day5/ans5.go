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
	allPoints := linesBuilder(read("input5.dat"))
	//fmt.Println(points)

	//p1
	var horiVertPoints []map[string]int
	var diagPoints []map[string]int
	mapLength, mapWidth := 0, 0

	for _, line := range allPoints {
		if line["x1"] > mapLength {
			mapLength = line["x1"]
		}

		if line["x2"] > mapLength {
			mapLength = line["x2"]
		}

		if line["y1"] > mapWidth {
			mapWidth = line["y1"]
		}

		if line["y2"] > mapWidth {
			mapWidth = line["y2"]
		}

		if line["x1"] == line["x2"] || line["y1"] == line["y2"] {
			horiVertPoints = append(horiVertPoints, line)
		}

		//if line["x1"]-line["x2"] != 0 && line["y1"]-line["y2"] != 0 {
		//	fmt.Println(int(math.Abs((line["y1"])-(line["x1"]))) / int(math.Abs(float64(line["y2"])-float64(line["x2"]))))
		//}
		if line["x1"]-line["x2"] != 0 {
			if abs(line["y1"]-line["y2"])/abs(line["x1"]-line["x2"]) == 1 {
				diagPoints = append(diagPoints, line)
			}
		}
	}

	//fmt.Println(mapLength, mapWidth, horiVertPoints)

	var terrainMap [][]int

	for i := 0; i < mapWidth+1; i++ {
		var row []int
		for j := 0; j < mapLength+1; j++ {
			row = append(row, 0)
		}
		terrainMap = append(terrainMap, row)
	}

	fmt.Println("map config -> len, wid ", len(terrainMap[0]), len(terrainMap))

	for _, p := range horiVertPoints {
		if p["x1"] == p["x2"] {
			if p["y1"] < p["y2"] {
				for i := 0; i < p["y2"]-p["y1"]+1; i++ {
					//fmt.Println(p["x1"], p["y1"]+i, p)
					terrainMap[p["y1"]+i][p["x1"]]++
				}
			} else {
				for i := 0; i < p["y1"]-p["y2"]+1; i++ {
					terrainMap[p["y2"]+i][p["x1"]]++
				}
			}
		} else {
			if p["x1"] < p["x2"] {
				for i := 0; i < p["x2"]-p["x1"]+1; i++ {
					terrainMap[p["y1"]][p["x1"]+i]++
				}
			} else {
				for i := 0; i < p["x1"]-p["x2"]+1; i++ {
					terrainMap[p["y1"]][p["x2"]+i]++
				}
			}
		}
	}

	dangerPointCounter := 0

	for i := 0; i < len(terrainMap); i++ {
		for j := 0; j < len(terrainMap[0]); j++ {
			if terrainMap[i][j] >= 2 {
				dangerPointCounter++
			}
		}
	}

	fmt.Println(dangerPointCounter)

	//p2

	fmt.Println(diagPoints)

	for _, p := range diagPoints {
		if p["x2"] > p["x1"] && p["y2"] > p["y1"] {
			for i := 0; i < p["x2"]-p["x1"]+1; i++ {
				terrainMap[p["y1"]+i][p["x1"]+i]++
				fmt.Println(p["x1"]+i, p["y1"]+i, "one")
			}
		} else if p["x1"] > p["x2"] && p["y1"] > p["y2"] {
			for i := 0; i < p["x1"]-p["x2"]+1; i++ {
				terrainMap[p["y2"]+i][p["x2"]+i]++
				fmt.Println(p["x2"]+i, p["y2"]+i, "two")
			}
		} else if p["x1"] > p["x2"] && p["y1"] < p["y2"] {
			for i := 0; i < p["x1"]-p["x2"]+1; i++ {
				terrainMap[p["y1"]+i][p["x1"]-i]++
				fmt.Println(p["x1"]-i, p["y1"]+i, "three")
			}
		} else if p["x1"] < p["x2"] && p["y1"] > p["y2"] {
			for i := 0; i < p["x2"]-p["x1"]+1; i++ {
				terrainMap[p["y1"]-i][p["x1"]+i]++
				fmt.Println(p["x1"]+i, p["y1"]-i, "four")
			}
		}
	}

	dangerPointCounterP2 := 0

	for i := 0; i < len(terrainMap); i++ {
		for j := 0; j < len(terrainMap[0]); j++ {
			if terrainMap[i][j] >= 2 {
				dangerPointCounterP2++
			}
		}
	}

	fmt.Println(dangerPointCounterP2)

	/*
		for _, row := range terrainMap {
			fmt.Println(row)
		}
	*/

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

func linesBuilder(lines []string) []map[string]int {
	var points []map[string]int

	for _, l := range lines {
		twoPoints := strings.Split(l, " -> ")

		lineMap := make(map[string]int)
		start := strings.Split(twoPoints[0], ",")
		end := strings.Split(twoPoints[1], ",")

		x1, _ := strconv.Atoi(start[0])
		y1, _ := strconv.Atoi(start[1])
		x2, _ := strconv.Atoi(end[0])
		y2, _ := strconv.Atoi(end[1])

		lineMap["x1"] = x1
		lineMap["x2"] = x2
		lineMap["y1"] = y1
		lineMap["y2"] = y2

		points = append(points, lineMap)
	}

	return points
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
