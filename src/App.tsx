import { useEffect, useRef, useState } from "react";
import { useCookies } from "react-cookie";

import "./App.css";
import { DEFAULT_SETTINGS } from "./constants";
import { GameParams, LeaderboardsScore, State } from "./types";
import { fetchLeaderboard } from "./utils";
import Board from "./components/Board";
import OptionsPanel from "./components/Options";
import Scoreboard from "./components/Scoreboard";
import Leaderboards from "./components/Leaderboards";

import { connector, ReduxProps } from "./store";

function App({ game, gameInit }: ReduxProps ) {
  const [settings, setSettings] = useState<GameParams>(DEFAULT_SETTINGS);
  const [scores, setScores] = useState<LeaderboardsScore>();
  const [cookies, setCookies] = useCookies();

  const prevState = useRef<State>();

  useEffect(() => {
    gameInit(settings)
    fetchLeaderboard().then((data) => setScores(data));
  }, [settings, gameInit]);


  useEffect(() => {
    if (prevState.current === State.Playing && game?.state === State.Win) {
      if (!cookies?.name) {
        setCookies("name", window.prompt("Enter name to save your score"), {
          maxAge: 15,
          sameSite: "lax",
        });
      } else {
        console.log(`You won ${cookies.name} @ ${new Date()}`);
      }
    }
    prevState.current = game?.state;
  }, [game, cookies, setCookies]);


  return (
    <div className="App">
      <div className="game">
        <Scoreboard
          state={game.state}
          flagCount={game.flagCount}
          restartGame={() => gameInit(settings)}
        />
        <Board />
        <OptionsPanel
          settings={settings}
          setSettings={setSettings}
          startGame={() => gameInit(settings)}
        />
        <Leaderboards
          beginner={scores?.beginner}
          intermediate={scores?.intermediate}
          expert={scores?.expert}
        />
      </div>
    </div>
  );
}

export default connector(App);
