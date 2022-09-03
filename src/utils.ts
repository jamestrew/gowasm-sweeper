import { supabase } from "./supabaseClient";
import { Difficulty, GameData, LeaderboardsScore, Score, Scores } from "./types";

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
    .order("time", { ascending: true })
    .limit(10);

  if (error) return [];

  return data.map((row: any) => ({
    name: row.name,
    time: row.time,
  }));
};

const difficultyScores = (difficultyData: Score[]): Scores => {
  return {
    times: difficultyData,
    recordCutOff: difficultyData[difficultyData.length - 1].time
  }
}

export const fetchLeaderboard = async (): Promise<LeaderboardsScore> => {
  let ret: LeaderboardsScore = {};
  await Promise.all([
    fetchDifficultyScores(Difficulty.Beginner),
    fetchDifficultyScores(Difficulty.Intermediate),
    fetchDifficultyScores(Difficulty.Expert),
  ]).then((scores) => {
    ret = {
      beginner: difficultyScores(scores[Difficulty.Beginner]),
      intermediate: difficultyScores(scores[Difficulty.Intermediate]),
      expert: difficultyScores(scores[Difficulty.Expert]),
    };
  });

  return ret;
};
