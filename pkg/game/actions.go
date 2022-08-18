package game

// FIX: over-opening
func (g *Game) OpenCell(x, y int) {
	if g.Flagged[y][x] || !g.playable() {
		return
	}

	if g.State == Unstarted {
		g.GenerateCleanMines(Pos{x, y})
		g.CalcAllNeighbors()
		g.State = Playing
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
		g.flagAllMines()
	}
}

func (g *Game) ToggleFlag(x, y int) {
	if !g.Open[y][x] && g.playable() {
		g.Flagged[y][x] = !g.Flagged[y][x]
	}
}

// TODO: right click opening
func (g *Game) ChordedOpen(x, y int) {
	if !g.Open[y][x] || g.State != Playing || !g.flagCountMatchesCell(x, y) {
		return
	}

	for _, pos := range g.CellNeighbors(x, y) {
		g.OpenCell(pos.X, pos.Y)
	}
}
