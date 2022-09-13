import { useEffect, useRef } from "react";
import { useDispatch } from "react-redux";
import { useCookies } from "react-cookie";
import "../App.css";

import { fetchLeaderboard, truncateStr, saveNewScore } from "../utils";
import { Difficulty, Scores, State } from "../types";
import { useSelector } from "../store";
import { leaderboardsInit } from "../slices/leaderboards";

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
    const newHighScore = (): boolean => {
      if (settings.difficulty === Difficulty.Custom) return false;
      if (time >= (leaderboards[settings.difficulty]?.recordCutOff ?? 0)) return false;
      return true;
    };

    const playerName = (): string => {
      if (cookies?.name) return cookies.name;
      const playerName = window.prompt("Enter name to save your score")?.trim() ?? "UNKNOWN PLAYER";
      setCookies("name", playerName, {
        maxAge: 15,
        sameSite: "lax",
      });
      return playerName;
    };

    if (prevState.current === State.Playing && game.state === State.Win) {
      if (newHighScore()) {
        saveNewScore(playerName(), settings.difficulty, time);
      }
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
