package reveal_minesweeper

var d = map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9"}

func RevealMinesweeper(b [][]string, row int, col int) [][]string {
	if b[row][col] == "M" {
		b[row][col] = "X"
		return b
	}
	dive(b, row, col)
	return b
}

func dive(b [][]string, x, y int) [][]string {
	if x < 0 || x >= len(b) || y < 0 || y >= len(b[0]) || b[x][y] != "H" {
		return b
	}
	var c int
	c += hasm(b, x-1, y-1)
	c += hasm(b, x-1, y)
	c += hasm(b, x-1, y+1)
	c += hasm(b, x, y-1)
	c += hasm(b, x, y+1)
	c += hasm(b, x+1, y-1)
	c += hasm(b, x+1, y)
	c += hasm(b, x+1, y+1)
	b[x][y] = d[c]
	b = dive(b, x-1, y-1)
	b = dive(b, x-1, y)
	b = dive(b, x-1, y+1)
	b = dive(b, x, y-1)
	b = dive(b, x, y+1)
	b = dive(b, x+1, y-1)
	b = dive(b, x+1, y)
	b = dive(b, x+1, y+1)
	return b
}

func hasm(b [][]string, x, y int) int {
	if x >= 0 && x < len(b) && y >= 0 && y < len(b[0]) {
		if b[x][y] == "M" {
			return 1
		}
	}
	return 0
}
