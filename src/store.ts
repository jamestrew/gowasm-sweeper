import { configureStore } from "@reduxjs/toolkit";
import {
  TypedUseSelectorHook,
  useSelector as rawUseSelector,
  connect,
  ConnectedProps,
} from "react-redux";

import gameDataSlice, {
  gameInit,
  gameOpenCell,
  gameFlagCell,
  gameChordedOpen,
} from "./slices/game";
import timerSlice, { timerIncr, timerReset } from "./slices/timer";
import leaderboardsSlice, { leaderboardsInit } from "./slices/leaderboards";
import settingsSlice from "./slices/settings";

export const store = configureStore({
  reducer: {
    gameData: gameDataSlice,
    timer: timerSlice,
    leaderboards: leaderboardsSlice,
    settings: settingsSlice,
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
  gameOpenCell,
  gameFlagCell,
  gameChordedOpen,
  timerIncr,
  timerReset,
  leaderboardsInit,
};

export const connector = connect(mapState, mapDispatch);
export type ReduxProps = ConnectedProps<typeof connector>;
