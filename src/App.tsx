import { useState } from "react";
import "./App.css";

const CELL_SIZE = 30;

function App() {
	const [board, setBoard] = useState<number[][]>([]);

	const width = board[0]?.length ?? 0;
	const height = board.length;

	const boardWidth = `${CELL_SIZE * width}px`;
	const boardHeight = `${CELL_SIZE * height}px`;

	return (
		<div className="App">
			<button onClick={() => setBoard(JSON.parse(window.newGame(0)))}>
				New Game
			</button>
			<div className="Board" style={{ width: boardWidth, height: boardHeight }}>
				{board.map((row, i) =>
					row.map((_, j) => (
						<Cell
							x={j}
							y={i}
							key={i * width + j}
							board={board}
							openCell={() => setBoard(JSON.parse(window.openCell(j, i)))}
						/>
					))
				)}
			</div>
		</div>
	);
}

type CellProps = {
	x: number;
	y: number;
	board: number[][];
	openCell: () => void;
};

const Cell = ({ x, y, board, openCell }: CellProps) => {
	return (
		<div
			className="Cell"
			style={{
				width: `${CELL_SIZE - 6}px`,
				height: `${CELL_SIZE - 6}px`,
				gridColumnStart: x + 1,
				gridRowStart: y + 1,
				background: "white",
			}}
			onClick={openCell}
		>
			{`${board[y][x]}`}
		</div>
	);
};

export default App;
