package game

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
}

func (g *Game) FlagCell(x, y int) {
	if !g.Open[y][x] {
		g.Flagged[y][x] = true
	}
}

// TODO: right click opening