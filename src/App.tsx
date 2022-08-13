import { useEffect, useState } from "react";
import "./App.css";

import { DEFAULT_SETTINGS } from "./constants";
import { GameData } from "./types";
import { gameObj } from "./utils";
import Board from "./components/Board";

function App() {
	const [game, setGame] = useState<GameData>();

	useEffect(() => {
		setGame(gameObj(window.newGame(DEFAULT_SETTINGS)));
	}, []);

	return (
		<div className="App">
			<button
				onClick={() =>
					setGame(
						gameObj(window.newGame({ ...DEFAULT_SETTINGS, difficulty: 3 }))
					)
				}
			>
				New Game
			</button>
			{game?.board && <Board board={game.board} setGame={setGame} />}
		</div>
	);
}

export default App;
