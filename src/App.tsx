import { useEffect, useState } from "react";
import "./App.css";

import { DEFAULT_SETTINGS } from "./constants";
import { GameData, State } from "./types";
import { gameObj } from "./utils";
import Board from "./components/Board";
import OptionsPanel from "./components/Options";

function App() {
	const [game, setGame] = useState<GameData>();

	useEffect(() => {
		setGame(gameObj(window.newGame(DEFAULT_SETTINGS)));
	}, []);

	switch (game?.state) {
		case State.Win:
			console.log("Winner!");
			break;
		case State.Lose:
			console.log("Loser!");
			break;
	}

	return (
		<div className="App">
			<div className="game">
				{game?.board && <Board board={game.board} setGame={setGame} />}
				<OptionsPanel
					onNewGame={(gameParams) =>
						setGame(gameObj(window.newGame(gameParams)))
					}
				/>
			</div>
		</div>
	);
}

export default App;
