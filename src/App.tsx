import React, { useState } from "react";
import "./App.css";

const CELL_SIZE = 30;

function App() {
	const [board, setBoard] = useState<number[][]>(
		JSON.parse(
			window.newGame(
				DEFAULT_SETTINGS.difficulty,
				DEFAULT_SETTINGS.width,
				DEFAULT_SETTINGS.height,
				DEFAULT_SETTINGS.mineCount
			)
		)
	);

	const width = board[0]?.length ?? 0;
	const height = board.length;

	const boardWidth = `${CELL_SIZE * width}px`;
	const boardHeight = `${CELL_SIZE * height}px`;

	return (
		<div className="App">
			<div className="game">
				<Options
					onNewGame={(settings) =>
						setBoard(
							JSON.parse(
								window.newGame(
									settings.difficulty,
									settings.width,
									settings.height,
									settings.mineCount
								)
							)
						)
					}
				/>
				<div
					className="Board"
					style={{ width: boardWidth, height: boardHeight }}
				>
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

type OptionsProps = {
	onNewGame: (settings: GameSettings) => void;
};

type GameSettings = {
	difficulty: number;
	width: number;
	height: number;
	mineCount: number;
};

const DEFAULT_SETTINGS = {
	difficulty: 0,
	width: 9,
	height: 9,
	mineCount: 10,
};

enum BoardParam {
	Width,
	Height,
	MineCount,
}

const Options = ({ onNewGame }: OptionsProps) => {
	const [settings, setSettings] = useState<GameSettings>(DEFAULT_SETTINGS);

	const isValidCustomBoard = (settings: GameSettings): boolean => {
		return settings.width * settings.height > settings.mineCount;
	};

	const handleInput = (
		event: React.ChangeEvent<HTMLInputElement>,
		paramType: BoardParam
	): void => {
		const param = parseInt(event.target.value);
		if (isNaN(param)) return;

		const newSettings = { ...settings };
		switch (paramType) {
			case BoardParam.Width:
				newSettings.width = param;
				break;
			case BoardParam.Height:
				newSettings.height = param;
				break;
			case BoardParam.MineCount:
				newSettings.mineCount = param;
				break;
		}

		if (isValidCustomBoard(newSettings)) {
			setSettings(newSettings);
		} else {
			window.alert("too many mines for board dimensions");
		}
	};

	// TODO: refactor this to use an array mapping or something less html-y
	return (
		<div>
			<table>
				<thead>
					<tr>
						<th></th>
						<th>Width</th>
						<th>Height</th>
						<th>Mines</th>
					</tr>
				</thead>
				<tbody>
					<tr>
						<td>
							<input
								type="radio"
								name="difficulty"
								id="beginner"
								onClick={() => setSettings({ ...settings, difficulty: 0 })}
							/>
						</td>
						<td>Beginner</td>
						<td>9</td>
						<td>9</td>
						<td>10</td>
					</tr>
					<tr>
						<td>
							<input
								type="radio"
								name="difficulty"
								id="intermediate"
								onClick={() => setSettings({ ...settings, difficulty: 1 })}
							/>
						</td>
						<td>Intermediate</td>
						<td>16</td>
						<td>16</td>
						<td>40</td>
					</tr>
					<tr>
						<td>
							<input
								type="radio"
								name="difficulty"
								id="expert"
								onClick={() => setSettings({ ...settings, difficulty: 2 })}
							/>
						</td>
						<td>Expert</td>
						<td>30</td>
						<td>16</td>
						<td>99</td>
					</tr>
					<tr>
						<td>
							<input
								type="radio"
								name="difficulty"
								id="custom"
								onClick={() => setSettings({ ...settings, difficulty: 3 })}
							/>
						</td>
						<td>Custom</td>
						<td>
							<input
								type="text"
								onChange={(e) => handleInput(e, BoardParam.Width)}
								value={settings.width}
							/>
						</td>
						<td>
							<input
								type="text"
								onChange={(e) => handleInput(e, BoardParam.Height)}
								value={settings.height}
							/>
						</td>
						<td>
							<input
								type="text"
								onChange={(e) => handleInput(e, BoardParam.MineCount)}
								value={settings.mineCount}
							/>
						</td>
					</tr>
				</tbody>
			</table>
			<button onClick={() => onNewGame(settings)}>New Game</button>
		</div>
	);
};

export default App;
