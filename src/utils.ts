import { GameData } from "./types";

export const gameObj = (gameDataStr: string): GameData => {
  let ret: GameData = { state: -1, board: [[-1]], flagCount: -1 };
  try {
    ret = JSON.parse(gameDataStr);
  } catch {
    console.error(gameDataStr);
  }
  return ret;
};

export const truncateStr = (str: string, length: number): string => {
  if (str.length <= length) return str;

  let newStr = str.substring(0, length - 3) + "...";
  return newStr;
};
