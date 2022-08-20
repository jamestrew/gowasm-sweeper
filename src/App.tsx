import { useCallback, useEffect, useState } from "react";
import "./App.css";

import { DEFAULT_SETTINGS } from "./constants";
import { GameParams, State } from "./types";
import { useGame } from './hooks'
import Board from "./components/Board";
import OptionsPanel from "./components/Options";
import Scoreboard from "./components/Scoreboard";


function App() {
  const [settings, setSettings] = useState<GameParams>(DEFAULT_SETTINGS);
	const [game, setGame] = useGame();

	const startGame = useCallback((settings: GameParams) => {
    setGame(window.newGame(settings));
	}, [setGame]);


	useEffect(() => {
		startGame(DEFAULT_SETTINGS);
	}, [startGame]);

	return (
		<div className="App">
			<div className="game">
				<Scoreboard
					state={game?.state || State.Unstarted}
					flagCount={game?.flagCount ?? 999}
					restartGame={() => startGame(settings)}
				/>
				{game?.board && <Board board={game.board} setGame={setGame} />}
				<OptionsPanel
					settings={settings}
					setSettings={setSettings}
					startGame={() => startGame(settings)}
				/>
			</div>
		</div>
	);
}

export default App;
