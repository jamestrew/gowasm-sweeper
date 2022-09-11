import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { LeaderboardsScore } from "../types";

const leaderboardsSlice = createSlice({
  name: "leaderboards",
  initialState: {} as LeaderboardsScore,
  reducers: {
    leaderboardsInit: (_: LeaderboardsScore, action: PayloadAction<LeaderboardsScore>) =>
      action.payload,
  },
});

export const { leaderboardsInit } = leaderboardsSlice.actions;
export default leaderboardsSlice.reducer;
