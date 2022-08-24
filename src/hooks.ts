import { useState, useCallback } from "react";
import { GameData } from "./types";

export const useGame = (
  gameData: GameData
): [GameData, (gameStr: string) => void] => {
  const [game, setGame] = useState(gameData ?? null);

  const parseAndSetGame = useCallback((gameStr: string) => {
    setGame(JSON.parse(gameStr));
  }, []);

  return [game, parseAndSetGame];
};
