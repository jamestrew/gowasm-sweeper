import { supabase } from "./supabaseClient";
import { GameData, LeaderboardsScore } from "./types";

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

export const fetchData = async (): Promise<LeaderboardsScore> => {
  let { data, error, status } = await supabase
    .from("leaderboard")
    .select(`name, time, difficulties!inner(*)`)
    .eq("difficulties.description", "beginner")
    .order('time', { ascending: true });

  return {
    beginnerScore: data?.map((row: any) => ({
      name: row.name,
      time: row.time,
    })),
    intermediateScore: [{ name: "jimbo", time: 96 }],
    expertScore: [{ name: "timmy", time: 300 }],
  };
};
