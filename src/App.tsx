import { useCallback, useEffect, useRef, useState } from "react";
import { useCookies } from "react-cookie";

import "./App.css";
import { DEFAULT_GAME, DEFAULT_SETTINGS } from "./constants";
import { GameParams, LeaderboardsScore, State } from "./types";
import { useGame } from "./hooks";
import { fetchLeaderboard } from './utils'
import Board from "./components/Board";
import OptionsPanel from "./components/Options";
import Scoreboard from "./components/Scoreboard";
import Leaderboards from "./components/Leaderboards";


function App() {
  const [settings, setSettings] = useState<GameParams>(DEFAULT_SETTINGS);
  const [game, setGame] = useGame(DEFAULT_GAME);
  const [scores, setScores] = useState<LeaderboardsScore>()
  const [cookies, setCookies] = useCookies();

  const prevState = useRef<State>();

  const startGame = useCallback(
    (settings: GameParams) => {
      setGame(window.newGame(settings));
    },
    [setGame]
  );

  useEffect(() => {
    startGame(DEFAULT_SETTINGS);
    fetchLeaderboard().then(data => setScores(data));
  }, [startGame]);

  useEffect(() => {
    if (prevState.current === State.Playing && game.state === State.Win) {
      if (!cookies?.name) {
        setCookies("name", window.prompt("Enter name to save your score"), {
          maxAge: 15,
          sameSite: "lax",
        });
      } else {
        console.log(`You won ${cookies.name} @ ${new Date()}`);
      }
    }
    prevState.current = game.state;
  }, [game, cookies, setCookies]);

  return (
    <div className="App">
      <div className="game">
        <Scoreboard
          state={game?.state || State.Unstarted}
          flagCount={game?.flagCount ?? 999}
          restartGame={() => startGame(settings)}
        />
        <Board board={game.board} setGame={setGame} />
        <OptionsPanel
          settings={settings}
          setSettings={setSettings}
          startGame={() => startGame(settings)}
        />
        <Leaderboards
          beginnerScore={scores?.beginnerScore}
          intermediateScore={scores?.intermediateScore}
          expertScore={scores?.expertScore}
          />
      </div>
    </div>
  );
}

export default App;
