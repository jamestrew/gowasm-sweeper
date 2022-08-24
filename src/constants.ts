import { GameData, GameParams, Score, State } from "./types";

export const DEFAULT_SETTINGS: GameParams = {
  difficulty: 0,
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
  { difficulty: 0, width: 9, height: 9, mineCount: 10 },
  { difficulty: 1, width: 16, height: 16, mineCount: 40 },
  { difficulty: 2, width: 30, height: 16, mineCount: 99 },
  { difficulty: 3, width: 5, height: 5, mineCount: 5 },
];

export const CELL_SIZE = 30;

export const DUMMY_LEADERBOARD_DATA: Score[] = [
  { name: "john", time: 54 },
  { name: "jason", time: 55 },
  { name: "MUNZE KONZA", time: 56 },
  { name: "TWITCH.TV/SOME_STREAMING_GUY", time: 57 },
  { name: "TWITCH.TV/SOME_STREAMING_GUY", time: 58 },
  { name: "TWITCH.TV/SOME_STREAMING_GUY", time: 58 },
  { name: "TWITCH.TV/SOME_STREAMING_GUY", time: 59 },
];

