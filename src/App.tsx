import { useCallback, useEffect, useRef, useState } from "react";
import { supabase } from "./supabaseClient";
import "./App.css";

import { DEFAULT_GAME, DEFAULT_SETTINGS } from "./constants";
import { GameParams, State } from "./types";
import { useGame } from "./hooks";
import Board from "./components/Board";
import OptionsPanel from "./components/Options";
import Scoreboard from "./components/Scoreboard";
import Leaderboards from "./components/Leaderboards";
import { useCookies } from "react-cookie";

const fetchData = async () => {
  let { data, error, status } = await supabase
    .from("difficulties")
    .select(`id, description`);

  console.log({ data, error, status });
};

function App() {
  const [settings, setSettings] = useState<GameParams>(DEFAULT_SETTINGS);
  const [game, setGame] = useGame(DEFAULT_GAME);
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
    fetchData();
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
        <Leaderboards />
      </div>
    </div>
  );
}

export default App;
