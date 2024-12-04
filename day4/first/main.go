package main

import "fmt"

func countXMAS(grid [][]rune) int {
	directions := [][2]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}

	word := "XMAS"
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	fmt.Println("Grid dimensions:", rows, "x", cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, d := range directions {
				x, y := i, j
				match := true

				for k := 0; k < len(word); k++ {
					if x < 0 || x >= rows || y < 0 || y >= cols || grid[x][y] != rune(word[k]) {
						match = false
						break
					}
					x += d[0]
					y += d[1]
				}

				if match {
					count++
				}
			}
		}
	}
	return count
}

func buildGrid(input string) [][]rune {
	var grid [][]rune
	rows := []rune(input)

	row := []rune{}
	for _, char := range rows {
		if char == '\n' {
			grid = append(grid, row)
			row = []rune{}
		} else {
			row = append(row, char)
		}
	}

	if len(row) > 0 {
		grid = append(grid, row)
	}

	fmt.Println("Grid built with", len(grid), "rows")
	return grid
}

func main() {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	grid := buildGrid(input)
	result := countXMAS(grid)
	fmt.Println("Total occurrences of XMAS:", result)

}
