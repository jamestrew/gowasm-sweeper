import { GameParams } from "../types";

type BoardParam = "width" | "height" | "mineCount";

type DifficultyRowProps = {
	option: GameParams;
	settings: GameParams;
	setSettings: (settings: GameParams) => void;
};

const DifficultyRow = ({
	option,
	settings,
	setSettings,
}: DifficultyRowProps) => {
	const difficulties: { [code: number]: string } = {
		0: "Beginner",
		1: "Intermediate",
		2: "Expert",
		3: "Custom",
	};

	const paramInput = (paramType: BoardParam) => {
		return (
			<td>
				<input
					className="custom-input"
					type="number"
					name={paramType}
					onChange={(e) =>
						setSettings({
							...settings,
							[e.target.name]:
								parseInt(e.target.value) ||
								settings[e.target.name as BoardParam],
						})
					}
					value={settings[paramType]}
				/>
			</td>
		);
	};

	return (
		<>
			<td>
				<input
					type="radio"
					name="difficulty"
					onClick={() =>
						setSettings({ ...settings, difficulty: option.difficulty })
					}
					defaultChecked={settings.difficulty === option.difficulty}
				/>
			</td>
			<td>{difficulties[option.difficulty]}</td>
			{option.difficulty === 3 ? (
				<>
					{paramInput("width")}
					{paramInput("height")}
					{paramInput("mineCount")}
				</>
			) : (
				<>
					<td>{option.width}</td>
					<td>{option.height}</td>
					<td>{option.mineCount}</td>
				</>
			)}
		</>
	);
};

export default DifficultyRow;
