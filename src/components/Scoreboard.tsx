import { useEffect } from "react";
import { useDispatch } from "react-redux";

import "../App.css";
import { GameData, State } from "../types";

import { timerIncr, timerReset } from '../slices/timer'
import { useSelector } from "../store";

type ScoreboardProps = {
  game: GameData,
  restartGame: () => void;
};

const Scoreboard = ({
  game,
  restartGame,
}: ScoreboardProps) => {
  const timerCount = useSelector(state => state.timer)
  const dispatch = useDispatch()

  useEffect(() => {
    if (game.state === State.Playing && timerCount <= 999) {
      const myInterval = setInterval(() => {
        dispatch(timerIncr())
      }, 1000);
      return () => clearInterval(myInterval);
    }

    if (game.state === State.Unstarted) {
      dispatch(timerReset())
    }
  }, [game.state, timerCount, dispatch ]);

  return (
    <div className="Scoreboard">
      <Counter value={timerCount} />
      <MinesweeperGuy state={game.state} restartGame={restartGame} />
      <Counter value={game.flagCount} />
    </div>
  );
};

type CounterProps = {
  value: number;
};

const Counter = ({ value }: CounterProps) => {
  return (
    <div className="Counter">
      {value.toString()}
    </div>
  );
};

type MinesweeperGuyProps = {
  state: State;
  restartGame: () => void;
};

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
