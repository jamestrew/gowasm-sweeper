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

export const boardDimensions = (game?: GameData): { width: number, height: number } => {
	if (game?.board) {
		return { width: game.board[0].length, height: game.board.length };
	}
  return { width: 0, height: 0};
};
