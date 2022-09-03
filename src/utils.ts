import { supabase } from "./supabaseClient";
import { Difficulty, GameData, LeaderboardsScore, Score } from "./types";

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

const fetchDifficultyScores = async (
  difficulty: Difficulty
): Promise<Score[]> => {
  const { data, error } = await supabase
    .from("leaderboard")
    .select(`name, time, difficulties!inner(*)`)
    .eq("difficulties.id", difficulty)
    .order("time", { ascending: true });

  if (error) return [];

  return data.map((row: any) => ({
    name: row.name,
    time: row.time,
  }));
};

export const fetchLeaderboard = async (): Promise<LeaderboardsScore> => {
  let ret: LeaderboardsScore = {}
  await Promise.all([
    fetchDifficultyScores(Difficulty.Beginner),
    fetchDifficultyScores(Difficulty.Intermediate),
    fetchDifficultyScores(Difficulty.Expert),
  ]).then((scores) => {
    ret = {
      beginner: scores[Difficulty.Beginner],
      intermediate: scores[Difficulty.Intermediate],
      expert: scores[Difficulty.Expert],
    };
  });

  return ret;
};
