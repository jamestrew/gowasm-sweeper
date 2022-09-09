import { configureStore } from "@reduxjs/toolkit";
import {
  TypedUseSelectorHook,
  useSelector as rawUseSelector,
  connect,
  ConnectedProps,
} from "react-redux";

import gameDataSlice, {
  gameInit,
  gameUpdate,
  gameOpenCell,
  gameFlagCell,
  gameChordedOpen,
} from "./slices/game";
import timerSlice, { timerIncr, timerReset } from "./slices/timer";
import leaderboardsSlice, { leaderboardsInit } from "./slices/leaderboards";

export const store = configureStore({
  reducer: {
    gameData: gameDataSlice,
    timer: timerSlice,
    leaderboards: leaderboardsSlice,
  },
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export const useSelector: TypedUseSelectorHook<RootState> = rawUseSelector;

const mapState = (state: RootState) => ({
  game: state.gameData,
  timerCount: state.timer,
  leaderboards: state.leaderboards,
});

const mapDispatch = {
  gameInit,
  gameUpdate,
  gameOpenCell,
  gameFlagCell,
  gameChordedOpen,
  timerIncr,
  timerReset,
  leaderboardsInit,
};

export const connector = connect(mapState, mapDispatch);
export type ReduxProps = ConnectedProps<typeof connector>;
