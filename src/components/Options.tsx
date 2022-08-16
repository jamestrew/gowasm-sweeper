import { useEffect } from "react";

import { BOARD_OPTIONS } from "../constants";
import { GameParams } from "../types";
import DifficultyRow from "./DifficultyRow";

type OptionsProps = {
	settings: GameParams;
	setSettings: React.Dispatch<GameParams>;
	startGame: () => void;
};

const OptionsPanel = ({ settings, setSettings, startGame }: OptionsProps) => {
	const isValidCustomBoard = (settings: GameParams): boolean => {
		return settings.width * settings.height > settings.mineCount;
	};

	useEffect(() => {
		if (!isValidCustomBoard(settings)) {
			window.alert("Too many mines for the board dimensions");
			const newSettings = {
				...settings,
				mineCount: settings.width * settings.height - 1,
			};
			setSettings(newSettings);
		}
	}, [settings, setSettings]);

	return (
		<>
			<div className="Options">
				<table id="option-table">
					<thead>
						<tr>
							<th></th>
							<th></th>
							<th>Width</th>
							<th>Height</th>
							<th>Mines</th>
						</tr>
					</thead>
					<tbody>
						{BOARD_OPTIONS.map((option, idx) => (
							<tr key={idx}>
								<DifficultyRow
									option={option}
									settings={settings}
									setSettings={setSettings}
								/>
							</tr>
						))}
					</tbody>
				</table>
			</div>
			<button onClick={startGame}>New Game</button>
		</>
	);
};

export default OptionsPanel;
