import { configureStore } from "@reduxjs/toolkit";
import {
  TypedUseSelectorHook,
  useSelector as rawUseSelector,
  connect,
  ConnectedProps,
} from "react-redux";

import gameDataSlice, { initGame, updateGame } from "./slices/game";

export const store = configureStore({
  reducer: {
    gameData: gameDataSlice,
  },
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export const useSelector: TypedUseSelectorHook<RootState> = rawUseSelector;

const mapState = (state: RootState) => ({
  gameData: state.gameData,
});

const mapDispatch = {
  initGame,
  updateGame,
};

export const connector = connect(mapState, mapDispatch);
export type ReduxProps = ConnectedProps<typeof connector>;
