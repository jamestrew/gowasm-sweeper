package game

// FIX: over-opening
func (g *Game) OpenCell(x, y int) {
	if g.Flagged[y][x] || g.State != Playing {
		return
	}
	if g.Mines[y][x] == 9 {
		g.State = Lose
	}

	if g.Mines[y][x] == 0 {
		g.OpenBlankCells(x, y)
	}
	g.Open[y][x] = true
	g.OpenBlankNeighbors()
	if g.CheckWin() {
		g.State = Win
	}
}

func (g *Game) ToggleFlag(x, y int) {
	if !g.Open[y][x] && g.State == Playing {
		g.Flagged[y][x] = !g.Flagged[y][x]
	}
}

// TODO: right click opening
func (g *Game) ChordedOpen(x, y int) {
	if !g.Open[y][x] || g.State != Playing || !g.flagCountMatchesCell(x, y) {
		return
	}

	for _, pos := range g.cellNeighbors(x, y) {
		g.OpenCell(pos.X, pos.Y)
	}
}

func (g *Game) flagCountMatchesCell(x, y int) bool {
	flagCount := 0
	for _, pos := range g.cellNeighbors(x, y) {
		if g.Flagged[pos.Y][pos.X] {
			flagCount++
		}
	}
	return flagCount == g.Mines[y][x]
}
