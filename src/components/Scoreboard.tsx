import { useEffect, useState } from "react";
import "../App.css";

import { State } from "../types";

type ScoreboardProps = {
	state: State;
	flagCount: number;
};

const Scoreboard = ({ state, flagCount }: ScoreboardProps) => {
	const [seconds, setSeconds] = useState(0);

	useEffect(() => {
		if (state === State.Playing) {
			const myInterval = setInterval(() => {
				setSeconds((prev) => prev + 1);
			}, 1000);
			return () => clearInterval(myInterval);
		}

		if (state === State.Unstarted) {
			setSeconds(0);
		}
	}, [state]);

	return (
		<div className="Scoreboard">
			<Counter value={seconds} />
			<MinesweeperGuy state={state} />
			<Counter value={flagCount} />
		</div>
	);
};

type CounterProps = {
	value: number;
};

const Counter = ({ value }: CounterProps) => {
	return <div className="Counter" style={{ paddingRight: "1px"}}>{value.toString()}</div>;
};

type MinesweeperGuyProps = {
	state: State;
};

// TODO: onClick restart game with current settings
const MinesweeperGuy = ({ state }: MinesweeperGuyProps) => {
	const emoji = new Map<State, string>([
		[State.Unstarted, "😇"],
		[State.Playing, "😇"],
		[State.Lose, "💀"],
		[State.Win, "😎"],
	]);

	return <div className="MinesweeperGuy">{emoji.get(state)}</div>;
};

export default Scoreboard;
