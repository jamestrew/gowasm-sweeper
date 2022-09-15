import { useEffect, useRef } from "react";
import { useDispatch } from "react-redux";
import { useCookies } from "react-cookie";
import "../App.css";

import { fetchLeaderboard, truncateStr, saveNewScore } from "../utils";
import { Difficulty, LeaderboardsScore, Scores, State } from "../types";
import { useSelector } from "../store";
import { leaderboardsInit } from "../slices/leaderboards";
import { COOKIE_SETTINGS } from "../constants";

const newHighScore = (
  time: number,
  difficulty: Difficulty,
  leaderboards: LeaderboardsScore
): boolean => {
  if (difficulty === Difficulty.Custom) return false;
  if (time >= (leaderboards[difficulty]?.recordCutOff ?? 0)) return false;
  return true;
};

const getPlayerName = (cookies: { [x: string]: string }): string => {
  if (cookies?.name) return cookies.name;
  return window.prompt("Enter name to save your score")?.trim() ?? "UNKNOWN PLAYER";
};

const gameJustWon = (currState: State, prevState?: State): boolean => {
  return prevState === State.Playing && currState === State.Win;
};

const Leaderboards = () => {
  const [cookies, setCookies] = useCookies();
  const prevState = useRef<State>();

  const { leaderboards, game, settings, time } = useSelector((state) => ({
    leaderboards: state.leaderboards,
    game: state.gameData,
    settings: state.settings,
    time: state.timer,
  }));
  const dispatch = useDispatch();

  useEffect(() => {
    if (
      gameJustWon(game.state, prevState.current) &&
      newHighScore(time, settings.difficulty, leaderboards)
    ) {
      const playerName = getPlayerName(cookies);
      setCookies("name", playerName, COOKIE_SETTINGS);
      saveNewScore(playerName, settings.difficulty, time);
    }
    prevState.current = game.state;
  }, [game.state, settings.difficulty, time, cookies, leaderboards, setCookies]);

  useEffect(() => {
    fetchLeaderboard().then((scores) => dispatch(leaderboardsInit(scores)));
  }, [dispatch]);

  return (
    <div className="Leaderboards">
      <Leaderboard difficulty="Beginner" scores={leaderboards[Difficulty.Beginner]} />
      <Leaderboard difficulty="Intermediate" scores={leaderboards[Difficulty.Intermediate]} />
      <Leaderboard difficulty="Expert" scores={leaderboards[Difficulty.Expert]} />
    </div>
  );
};

type LeaderboardProps = {
  difficulty: string;
  scores?: Scores;
};

const Leaderboard = ({ difficulty, scores }: LeaderboardProps) => {
  // fix style when there's no scores
  // fix style for name lengths
  return (
    <div className="Leaderboard">
      <h4>{difficulty}</h4>
      <table>
        <tbody>
          {scores?.times
            .filter((_, idx) => idx < 5)
            .map((score, idx) => (
              <tr key={idx}>
                <td style={{ width: "1em", textAlign: "left" }}>{`${idx + 1}.`}</td>
                <td style={{ width: "8em" }}>{truncateStr(score.name, 15)}</td>
                <td style={{ width: "2em", textAlign: "right" }}>{score.time}</td>
              </tr>
            ))}
        </tbody>
      </table>
    </div>
  );
};

export default Leaderboards;
