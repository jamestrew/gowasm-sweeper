import { useEffect } from 'react'
import "../App.css";

import { fetchLeaderboard, truncateStr } from "../utils";
import { Scores } from "../types";
import { connector, ReduxProps } from "../store";

const Leaderboards = ({leaderboards, leaderboardsInit}: ReduxProps) => {
  useEffect(() => {
    fetchLeaderboard().then(scores => leaderboardsInit(scores))
  }, [leaderboards, leaderboardsInit])

  return (
    <div className="Leaderboards">
      <Leaderboard difficulty="Beginner" scores={leaderboards.beginner} />
      <Leaderboard difficulty="Intermediate" scores={leaderboards.intermediate} />
      <Leaderboard difficulty="Expert" scores={leaderboards.expert} />
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

export default connector(Leaderboards);
