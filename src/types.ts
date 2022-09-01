export type GameData = {
  state: State;
  board: number[][];
  flagCount: number;
};

export type GameParams = {
  difficulty: Difficulty,
  width: number;
  height: number;
  mineCount: number;
};

export enum State {
  Unstarted = 0,
  Playing,
  Win,
  Lose,
}

export type CellPosition = {
  x: number;
  y: number;
};

export type Score = { name: string; time: number };

export type LeaderboardsScore = {
  beginnerScore?: Score[];
  intermediateScore?: Score[];
  expertScore?: Score[];
};

export enum Difficulty {
  Beginner = 0,
  Intermediate,
  Expert,
  Custom
}
