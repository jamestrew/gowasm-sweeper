import { GameData } from "./types";

export const gameObj = (gameDataStr: string): GameData => {
	let ret: GameData = { state: -1, board: [[-1]], flagCount: -1 };
	try {
		const parsed = JSON.parse(gameDataStr);
		ret = {
			state: parsed.State,
			board: parsed.Board,
			flagCount: parsed.FlagCount,
		};
	} catch {
		console.error(gameDataStr);
	}
	return ret;
};
