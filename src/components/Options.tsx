import { useEffect, useState } from "react";

import { DEFAULT_SETTINGS, BOARD_OPTIONS } from "../constants";
import { GameParams } from "../types";
import DifficultyRow from './DifficultyRow'

type OptionsProps = {
	onNewGame: (settings: GameParams) => void;
};


const OptionsPanel = ({ onNewGame }: OptionsProps) => {
	const [settings, setSettings] = useState<GameParams>(DEFAULT_SETTINGS);

	const isValidCustomBoard = (settings: GameParams): boolean => {
		return settings.width * settings.height > settings.mineCount;
	};

	useEffect(() => {
		if (!isValidCustomBoard(settings)) {
			window.alert("Too many mines for the board dimensions");
		}
	}, [settings]);

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
			<button onClick={() => onNewGame(settings)}>New Game</button>
		</>
	);
};

export default OptionsPanel;
