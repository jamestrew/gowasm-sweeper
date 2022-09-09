import { configureStore } from "@reduxjs/toolkit";
import {
  TypedUseSelectorHook,
  useSelector as rawUseSelector,
  connect,
  ConnectedProps,
} from "react-redux";

import gameDataSlice, { initGame, updateGame } from "./slices/game";
import timerSlice, { timerIncr, timerReset } from "./slices/timer";

export const store = configureStore({
  reducer: {
    gameData: gameDataSlice,
    timer: timerSlice,
  },
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export const useSelector: TypedUseSelectorHook<RootState> = rawUseSelector;

const mapState = (state: RootState) => ({
  gameData: state.gameData,
  timerCount: state.timer,
});

const mapDispatch = {
  initGame,
  updateGame,
  timerIncr,
  timerReset,
};

export const connector = connect(mapState, mapDispatch);
export type ReduxProps = ConnectedProps<typeof connector>;
