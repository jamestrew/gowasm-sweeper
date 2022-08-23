import "../App.css"

import { truncateStr } from '../utils';
import { Score } from "../types";
import { DUMMY_LEADERBOARD_DATA } from "../constants";

const Leaderboards = () => {
  return (
    <div className="Leaderboards">
      <Leaderboard difficulty="Beginner" scores={DUMMY_LEADERBOARD_DATA} />
      <Leaderboard difficulty="Intermediate" scores={DUMMY_LEADERBOARD_DATA} />
      <Leaderboard difficulty="Expert" scores={DUMMY_LEADERBOARD_DATA} />
    </div>
  );
};

type LeaderboardProps = {
  difficulty: string;
  scores: Score[];
};

const Leaderboard = ({ difficulty, scores }: LeaderboardProps) => {
  return (
    <div className="Leaderboard">
      <h4>{difficulty}</h4>
      <table>
        {scores.map((score, idx) => (
          <tr>
            <td>{`${idx + 1}.`}</td>
            <td>{truncateStr(score.name, 15)}</td>
            <td>{score.time}</td>
          </tr>
        ))}
      </table>
    </div>
  );
};

export default Leaderboards;
