import { useEffect } from "react";

import "./App.css";
import Board from "./components/Board";
import OptionsPanel from "./components/Options";
import Scoreboard from "./components/Scoreboard";
import Leaderboards from "./components/Leaderboards";

import { gameInit } from "./slices/game";
import { useSelector } from "./store";
import { useDispatch } from "react-redux";

function App() {
  const { game, settings } = useSelector((state) => ({
    game: state.gameData,
    settings: state.settings,
  }));
  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(gameInit(settings));
  }, [settings, dispatch]);

  return (
    <div className="App">
      <div className="game">
        <Scoreboard game={game} restartGame={() => dispatch(gameInit(settings))} />
        <Board game={game} />
        <OptionsPanel startGame={() => dispatch(gameInit(settings))} />
        <Leaderboards />
      </div>
    </div>
  );
}

export default App;
