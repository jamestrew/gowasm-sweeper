import "../App.css";

import { truncateStr } from "../utils";
import { LeaderboardsScore, Score } from "../types";
import { DUMMY_LEADERBOARD_DATA } from "../constants";

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
              <td>{`${idx + 1}.`}</td>
              <td>{truncateStr(score.name, 15)}</td>
              <td>{score.time}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default Leaderboards;
