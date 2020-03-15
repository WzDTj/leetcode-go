/**
 * 695. 岛屿的最大面积
 * https://leetcode-cn.com/problems/max-area-of-island/
 **/

package problems

var directions [][]int
var rows, cols int

func MaxAreaOfIsland(grid [][]int) int {
	directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	rows, cols = len(grid), len(grid[0])

	var max int

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 0 {
				continue
			}

			var area int
			measureIslandArea(grid, directions, row, col, &area)
			if area > max {
				max = area
			}
		}
	}

	return max
}

func measureIslandArea(grid, directions [][]int, row, col int, area *int) {
	*area++
	grid[row][col] = 0

	for _, direction := range directions {
		nextRow := row + direction[0]
		nextCol := col + direction[1]
		if true &&
			nextRow >= 0 &&
			nextRow < rows &&
			nextCol >= 0 &&
			nextCol < cols &&
			grid[nextRow][nextCol] == 1 {
			measureIslandArea(grid, directions, nextRow, nextCol, area)
		}
	}
}
