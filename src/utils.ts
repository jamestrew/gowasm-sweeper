import { GameData } from "./types";

export const gameObj = (gameDataStr: string): GameData => {
	let ret: GameData = { state: -1, board: [[-1]], flagCount: -1 };
	try {
		ret = JSON.parse(gameDataStr);
	} catch {
		console.error(gameDataStr);
	}
	return ret;
};
