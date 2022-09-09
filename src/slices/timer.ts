import { createSlice } from "@reduxjs/toolkit";

const timerSlice = createSlice({
  name: "timer",
  initialState: 0,
  reducers: {
    timerIncr: (state: number) => state + 1,
    timerReset: (_) => 0,
  },
});

export const { timerIncr, timerReset } = timerSlice.actions;
export default timerSlice.reducer;
