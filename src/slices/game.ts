import { createSlice, PayloadAction } from "@reduxjs/toolkit";

import { CellPosition, GameData, GameParams } from "../types";
import { DEFAULT_GAME } from "../constants";

const gameDataSlice = createSlice({
  name: "gameData",
  initialState: DEFAULT_GAME,
  reducers: {
    gameInit: (_: GameData, action: PayloadAction<GameParams>) =>
      JSON.parse(window.newGame(action.payload)),
    gameOpenCell: (_: GameData, action: PayloadAction<CellPosition>) => {
      const { x, y } = action.payload;
      return JSON.parse(window.openCell(x, y));
    },
    gameFlagCell: (_: GameData, action: PayloadAction<CellPosition>) => {
      const { x, y } = action.payload;
      return JSON.parse(window.flagCell(x, y));
    },
    gameChordedOpen: (_: GameData, action: PayloadAction<CellPosition>) => {
      const { x, y } = action.payload;
      return JSON.parse(window.chordedOpen(x, y));
    },
  },
});

export const { gameInit, gameOpenCell, gameFlagCell, gameChordedOpen } = gameDataSlice.actions;
export default gameDataSlice.reducer;
