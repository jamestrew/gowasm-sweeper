import "../App.css";

import { truncateStr } from "../utils";
import { LeaderboardsScore, Score } from "../types";

const Leaderboards = ({
  beginnerScore,
  intermediateScore,
  expertScore,
}: LeaderboardsScore) => {
  return (
    <div className="Leaderboards">
      <Leaderboard difficulty="Beginner" scores={beginnerScore} />
      <Leaderboard difficulty="Intermediate" scores={intermediateScore} />
      <Leaderboard difficulty="Expert" scores={expertScore} />
    </div>
  );
};

type LeaderboardProps = {
  difficulty: string;
  scores?: Score[];
};

const Leaderboard = ({ difficulty, scores }: LeaderboardProps) => {
  // fix style when there's no scores
  // fix style for name lengths
  return (
    <div className="Leaderboard">
      <h4>{difficulty}</h4>
      <table>
        <tbody>
          {scores?.map((score, idx) => (
            <tr key={idx}>
              <td style={{ width: "1em", textAlign: "left" }}>{`${
                idx + 1
              }.`}</td>
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
