import "./App.css";

const CELL_SIZE = 30;

function App() {
	const width = 10;
	const height = 10;

	const boardWidth = `${CELL_SIZE * width}px`;
	const boardHeight = `${CELL_SIZE * height}px`;
	const board: number[][] = Array(10).fill(Array(10).fill(0));

	return (
		<div className="App">
			<div className="Board" style={{ width: boardWidth, height: boardHeight }}>
				{board.map((row, i) =>
					row.map((_, j) => <Cell x={j} y={i} key={i * 10 + j} />)
				)}
			</div>
		</div>
	);
}

type CellProps = {
	x: number;
	y: number;
};

const Cell = ({ x, y }: CellProps) => {
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
			onClick={() => console.log(`clicked on ${x},${y}`)}
		>
			{`${x},${y}`}
		</div>
	);
};

export default App;
