import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { LeaderboardsScore } from "../types";

const leaderboardsSlice = createSlice({
  name: "leaderboards",
  initialState: {} as LeaderboardsScore,
  reducers: {
    leaderboardsInit: (state: LeaderboardsScore, action: PayloadAction<LeaderboardsScore>) => {
      state = action.payload;
      return state;
    },
  },
});

export const { leaderboardsInit } = leaderboardsSlice.actions;
export default leaderboardsSlice.reducer;
