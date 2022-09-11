import { createSlice } from "@reduxjs/toolkit";

const initialState = 0;

const timerSlice = createSlice({
  name: "timer",
  initialState,
  reducers: {
    timerIncr: (state: number) => state + 1,
    timerReset: (_) => initialState,
  },
});

export const { timerIncr, timerReset } = timerSlice.actions;
export default timerSlice.reducer;
