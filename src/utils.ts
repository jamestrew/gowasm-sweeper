import { GameData } from "./types";

export const gameObj = (gameDataStr: string): GameData => {
	let ret: GameData = { state: -1, board: [[-1]] };
	try {
		const parsed = JSON.parse(gameDataStr);
		ret = { state: parsed.State, board: parsed.Board };
	} catch {
		console.error(gameDataStr);
	}
	return ret;
};
