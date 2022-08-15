import { useEffect, useState } from "react";
import "../App.css";

import { State } from "../types";

type ScoreboardProps = {
	startTime: Date;
	state: State;
	flagCount: number;
};

const Scoreboard = ({ startTime, state, flagCount }: ScoreboardProps) => {
	const [seconds, setSeconds] = useState(0);

	// FIX: timer temporarily pauses if inputs given
	useEffect(() => {
		if (state === State.Playing) {
			const myInterval = setInterval(() => {
				setSeconds(
					Math.round((new Date().getTime() - startTime.getTime()) / 1000)
				);
			}, 1000);
			return () => clearInterval(myInterval);
		}
	});

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
	return <div className="Counter">{value.toString()}</div>;
};

type MinesweeperGuyProps = {
	state: State;
};

// TODO: onClick restart game with current settings
const MinesweeperGuy = ({ state }: MinesweeperGuyProps) => {
	const emoji = new Map<State, string>([
		[State.Playing, "😇"],
		[State.Lose, "💀"],
		[State.Win, "😎"],
	]);

	return <div className="MinesweeperGuy">{emoji.get(state)}</div>;
};

export default Scoreboard;
