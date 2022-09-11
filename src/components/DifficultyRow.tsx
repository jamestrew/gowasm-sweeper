import { GameParams, Difficulty } from "../types";
import { useSelector } from "../store";
import { settingsUpdate } from "../slices/settings";
import { useDispatch } from "react-redux";

type BoardParam = "width" | "height" | "mineCount";

type DifficultyRowProps = {
  option: GameParams;
};

const DifficultyRow = ({ option }: DifficultyRowProps) => {
  const settings = useSelector((state) => state.settings);
  const dispatch = useDispatch();

  const difficulties = new Map<Difficulty, string>([
    [Difficulty.Beginner, "Beginner"],
    [Difficulty.Intermediate, "Intermediate"],
    [Difficulty.Expert, "Expert"],
    [Difficulty.Custom, "Custom"],
  ]);

  const paramInput = (paramType: BoardParam) => {
    return (
      <td>
        <input
          className="custom-input"
          type="number"
          name={paramType}
          onChange={(e) =>
            dispatch(
              settingsUpdate({
                ...settings,
                [e.target.name]: parseInt(e.target.value) || settings[e.target.name as BoardParam],
              })
            )
          }
          value={settings[paramType]}
        />
      </td>
    );
  };

  return (
    <>
      <td>
        <input
          type="radio"
          name="difficulty"
          onClick={() => dispatch(settingsUpdate({ ...settings, difficulty: option.difficulty }))}
          defaultChecked={settings.difficulty === option.difficulty}
        />
      </td>
      <td>{difficulties.get(option.difficulty)}</td>
      {option.difficulty === Difficulty.Custom ? (
        <>
          {paramInput("width")}
          {paramInput("height")}
          {paramInput("mineCount")}
        </>
      ) : (
        <>
          <td>{option.width}</td>
          <td>{option.height}</td>
          <td>{option.mineCount}</td>
        </>
      )}
    </>
  );
};

export default DifficultyRow;
