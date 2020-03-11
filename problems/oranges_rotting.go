/**
 * 994. 腐烂的橘子
 * https://leetcode-cn.com/problems/rotting-oranges/
 **/

package problems

func OrangesRotting(grid [][]int) int {
	// check grid & find rotted oranges
	maxRow, maxCol := len(grid), len(grid[0])
	freshOranges, hasRottedOranges := 0, false
	var rottedOranges []map[string]int

	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {
			switch grid[row][col] {
			case 1:
				freshOranges++
			case 2:
				hasRottedOranges = true
				position := map[string]int{"row": row, "col": col}
				rottedOranges = append(rottedOranges, position)
			}
		}
	}

	switch {
	case freshOranges == 0:
		return 0
	case !hasRottedOranges:
		return -1
	}

	days := 0
	for len(rottedOranges) > 0 {
		var nextDaysRottedOranges []map[string]int
		for i := 0; i < len(rottedOranges); i++ {
			row, col := rottedOranges[i]["row"], rottedOranges[i]["col"]
			// up
			if row-1 >= 0 && grid[row-1][col] == 1 {
				freshOranges--
				grid[row-1][col] = 2
				position := map[string]int{"row": row - 1, "col": col}
				nextDaysRottedOranges = append(nextDaysRottedOranges, position)
			}
			// down
			if row+1 < maxRow && grid[row+1][col] == 1 {
				freshOranges--
				grid[row+1][col] = 2
				position := map[string]int{"row": row + 1, "col": col}
				nextDaysRottedOranges = append(nextDaysRottedOranges, position)
			}
			// left
			if col-1 >= 0 && grid[row][col-1] == 1 {
				freshOranges--
				grid[row][col-1] = 2
				position := map[string]int{"row": row, "col": col - 1}
				nextDaysRottedOranges = append(nextDaysRottedOranges, position)
			}
			// right
			if col+1 < maxCol && grid[row][col+1] == 1 {
				freshOranges--
				grid[row][col+1] = 2
				position := map[string]int{"row": row, "col": col + 1}
				nextDaysRottedOranges = append(nextDaysRottedOranges, position)
			}
		}

		rottedOranges = nil
		if len(nextDaysRottedOranges) > 0 {
			rottedOranges = nextDaysRottedOranges
			days++
		}
	}

	if freshOranges > 0 {
		return -1
	}

	return days
}
