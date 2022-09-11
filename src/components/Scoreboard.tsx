import { useEffect } from "react";
import { useDispatch } from "react-redux";

import "../App.css";
import { State } from "../types";

import { timerIncr, timerReset } from '../slices/timer'
import { useSelector } from "../store";

type ScoreboardProps = {
  state: State;
  flagCount: number;
  restartGame: () => void;
};

const Scoreboard = ({
  state,
  flagCount,
  restartGame,
}: ScoreboardProps) => {
  const timerCount = useSelector(state => state.timer)
  const dispatch = useDispatch()

  useEffect(() => {
    if (state === State.Playing && timerCount <= 999) {
      const myInterval = setInterval(() => {
        dispatch(timerIncr())
      }, 1000);
      return () => clearInterval(myInterval);
    }

    if (state === State.Unstarted) {
      dispatch(timerReset())
    }
  }, [state, timerCount, dispatch ]);

  return (
    <div className="Scoreboard">
      <Counter value={timerCount} />
      <MinesweeperGuy state={state} restartGame={restartGame} />
      <Counter value={flagCount} />
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

export default Scoreboard;
