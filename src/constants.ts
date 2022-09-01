import { GameData, GameParams, State, Difficulty } from "./types";

export const DEFAULT_SETTINGS: GameParams = {
  difficulty: Difficulty.Beginner,
  width: 4,
  height: 3,
  mineCount: 5,
};

export const DEFAULT_GAME: GameData = {
  state: State.Unstarted,
  board: [[]],
  flagCount: 0,
};

export const BOARD_OPTIONS: GameParams[] = [
  { difficulty: Difficulty.Beginner, width: 9, height: 9, mineCount: 10 },
  { difficulty: Difficulty.Intermediate, width: 16, height: 16, mineCount: 40 },
  { difficulty: Difficulty.Expert, width: 30, height: 16, mineCount: 99 },
  { difficulty: Difficulty.Custom, width: 5, height: 5, mineCount: 5 },
];

export const CELL_SIZE = 30;

