package game

// FIX: over-opening
func (g *Game) OpenCell(x, y int) {
	if g.Flagged[y][x] {
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
	if !g.Open[y][x] {
		g.Flagged[y][x] = !g.Flagged[y][x]
	}
}

// TODO: right click opening
