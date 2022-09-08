import { createSlice, PayloadAction } from "@reduxjs/toolkit";

import { GameData, GameParams } from "../types";
import { DEFAULT_GAME } from "../constants";

const gameDataSlice = createSlice({
  name: "gameData",
  initialState: DEFAULT_GAME,
  reducers: {
    initGame: (state: GameData, action: PayloadAction<GameParams>) => {
      state = JSON.parse(window.newGame(action.payload));
      return state;
    },
    updateGame: (state: GameData, action: PayloadAction<GameData>) => {
      state = action.payload
      return state;
    }
  },
});

export const { initGame, updateGame } = gameDataSlice.actions;
export default gameDataSlice.reducer;
