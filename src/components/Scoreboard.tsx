import { useEffect } from "react";
import "../App.css";

import { connector, ReduxProps } from "../store";
import { State } from "../types";

type ScoreboardProps = {
  state: State;
  flagCount: number;
  restartGame: () => void;
} & ReduxProps;

const Scoreboard = ({
  restartGame,
  gameData,
  timerCount,
  timerIncr,
  timerReset,
}: ScoreboardProps) => {
  useEffect(() => {
    if (gameData.state === State.Playing && timerCount <= 999) {
      const myInterval = setInterval(() => {
        timerIncr();
      }, 1000);
      return () => clearInterval(myInterval);
    }

    if (gameData.state === State.Unstarted) {
      timerReset();
    }
  }, [gameData, timerCount, timerIncr, timerReset]);

  return (
    <div className="Scoreboard">
      <Counter value={timerCount} />
      <MinesweeperGuy state={gameData.state} restartGame={restartGame} />
      <Counter value={gameData.flagCount} />
    </div>
  );
};

type CounterProps = {
  value: number;
};

const Counter = ({ value }: CounterProps) => {
  return (
    <div className="Counter" style={{ paddingRight: "1px" }}>
      {value.toString()}
    </div>
  );
};

type MinesweeperGuyProps = {
  state: State;
  restartGame: () => void;
};

// TODO: onClick restart game with current settings
const MinesweeperGuy = ({ state, restartGame }: MinesweeperGuyProps) => {
  const emoji = new Map<State, string>([
    [State.Unstarted, "😇"],
    [State.Playing, "😇"],
    [State.Lose, "💀"],
    [State.Win, "😎"],
  ]);

  return (
    <div className="MinesweeperGuy" onClick={restartGame}>
      {emoji.get(state)}
    </div>
  );
};

export default connector(Scoreboard);
