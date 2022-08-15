import { useEffect, useState } from "react";
import "./App.css";

import { DEFAULT_SETTINGS } from "./constants";
import { GameData, GameParams, State } from "./types";
import { gameObj } from "./utils";
import Board from "./components/Board";
import OptionsPanel from "./components/Options";
import Scoreboard from "./components/Scoreboard";

function App() {
	const [game, setGame] = useState<GameData>();
	const [startTime, setStartTime] = useState<Date>(new Date());

	const startGame = (settings: GameParams) => {
		setStartTime(new Date());
		setGame(gameObj(window.newGame(settings)));
	};

	useEffect(() => {
		startGame(DEFAULT_SETTINGS);
	}, []);

	return (
		<div className="App">
			<div className="game">
				<Scoreboard
					startTime={startTime}
					state={game?.state || State.Playing}
					flagCount={10}
				/>
				{game?.board && <Board board={game.board} setGame={setGame} />}
				<OptionsPanel
					onNewGame={(gameParams: GameParams) => startGame(gameParams)}
				/>
			</div>
		</div>
	);
}

export default App;
