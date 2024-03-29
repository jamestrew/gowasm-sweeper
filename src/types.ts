export type GameData = {
  state: State;
  board: number[][];
  flagCount: number;
};

export type GameParams = {
  difficulty: Difficulty;
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

export type Scores = {
  times: Score[];
  recordCutOff: number;
};

export enum Difficulty {
  Beginner = 0,
  Intermediate,
  Expert,
  Custom,
}

export type LeaderboardsScore = {
  [Difficulty.Beginner]?: Scores;
  [Difficulty.Intermediate]?: Scores;
  [Difficulty.Expert]?: Scores;
  [Difficulty.Custom]?: null;
};
