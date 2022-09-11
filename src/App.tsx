import { useEffect, useRef, useState } from "react";
import { useCookies } from "react-cookie";

import "./App.css";
import { DEFAULT_SETTINGS } from "./constants";
import { GameParams, State } from "./types";
import Board from "./components/Board";
import OptionsPanel from "./components/Options";
import Scoreboard from "./components/Scoreboard";
import Leaderboards from "./components/Leaderboards";

import { gameInit } from './slices/game'
import { useSelector } from "./store";
import { useDispatch } from "react-redux";

function App() {
  const [settings, setSettings] = useState<GameParams>(DEFAULT_SETTINGS);
  const [cookies, setCookies] = useCookies();
  const game = useSelector((state) => state.gameData)
  const dispatch = useDispatch()


  const prevState = useRef<State>();

  useEffect(() => {
    dispatch(gameInit(settings))
  }, [settings, dispatch]);


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
          restartGame={() => dispatch(gameInit(settings))}
        />
        <Board />
        <OptionsPanel
          settings={settings}
          setSettings={setSettings}
          startGame={() => dispatch(gameInit(settings))}
        />
        <Leaderboards />
      </div>
    </div>
  );
}

export default App;
