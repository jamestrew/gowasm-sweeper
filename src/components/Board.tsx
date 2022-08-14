import { CELL_SIZE } from "../constants";
import { GameData } from "../types";
import { gameObj } from "../utils";
import Cell from "./Cell";

import '../App.css'

type BoardProps = {
	board: number[][];
	setGame: React.Dispatch<GameData>;
};

const Board = ({ board, setGame }: BoardProps) => {
	const width = board[0].length;
	const height = board.length;

	const boardWidth = `${CELL_SIZE * width}px`;
	const boardHeight = `${CELL_SIZE * height}px`;

	return (
		<div className="Board" style={{ width: boardWidth, height: boardHeight }}>
			{board.map((row, i) =>
				row.map((_, j) => (
					<Cell
						x={j}
						y={i}
						key={i * width + j}
						cellType={board[i][j]}
						openCell={() => setGame(gameObj(window.openCell(j, i)))}
						flagCell={() => setGame(gameObj(window.flagCell(j, i)))}
						chordedOpen={() => setGame(gameObj(window.chordedOpen(j, i)))}
					/>
				))
			)}
		</div>
	);
};

export default Board;
