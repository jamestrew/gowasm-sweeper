import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { DEFAULT_SETTINGS } from "../constants";
import { GameParams } from "../types";

const settingsSlice = createSlice({
  name: "settings",
  initialState: DEFAULT_SETTINGS,
  reducers: {
    settingsUpdate: (_: GameParams, action: PayloadAction<GameParams>) => action.payload,
  },
});

export const { settingsUpdate } = settingsSlice.actions;
export default settingsSlice.reducer;
