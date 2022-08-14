import { GameParams } from "./types";

export const DEFAULT_SETTINGS: GameParams = {
	difficulty: 0,
	width: 4,
	height: 3,
	mineCount: 5,
};

export const BOARD_OPTIONS: GameParams[] = [
	{ difficulty: 0, width: 9, height: 9, mineCount: 10 },
	{ difficulty: 1, width: 16, height: 16, mineCount: 40 },
	{ difficulty: 2, width: 30, height: 16, mineCount: 99 },
	{ difficulty: 3, width: 5, height: 5, mineCount: 5 },
];

export const CELL_SIZE = 30;
