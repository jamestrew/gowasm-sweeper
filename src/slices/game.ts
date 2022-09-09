import { createSlice, PayloadAction } from "@reduxjs/toolkit";

import { CellPosition, GameData, GameParams } from "../types";
import { DEFAULT_GAME } from "../constants";

const gameDataSlice = createSlice({
  name: "gameData",
  initialState: DEFAULT_GAME,
  reducers: {
    gameInit: (state: GameData, action: PayloadAction<GameParams>) => {
      state = JSON.parse(window.newGame(action.payload));
      return state;
    },
    gameUpdate: (state: GameData, action: PayloadAction<GameData>) => { // deprecate?
      state = action.payload;
      return state;
    },
    gameOpenCell: (state: GameData, action: PayloadAction<CellPosition>) => {
      const { x, y } = action.payload;
      state = JSON.parse(window.openCell(x, y));
      return state;
    },
    gameFlagCell: (state: GameData, action: PayloadAction<CellPosition>) => {
      const { x, y } = action.payload;
      state = JSON.parse(window.flagCell(x, y));
      return state;
    },
    gameChordedOpen: (state: GameData, action: PayloadAction<CellPosition>) => {
      const { x, y } = action.payload;
      state = JSON.parse(window.chordedOpen(x, y));
      return state;
    },
  },
});

export const { gameInit, gameUpdate, gameOpenCell, gameFlagCell, gameChordedOpen } =
  gameDataSlice.actions;
export default gameDataSlice.reducer;
