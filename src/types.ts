export type GameData = {
  state: State;
  board: number[][];
	flagCount: number;
};

export type GameParams = {
  difficulty: number;
  width: number;
  height: number;
  mineCount: number;
};

export enum State {
	Unstarted = 0,
	Playing,
	Win,
	Lose
}
