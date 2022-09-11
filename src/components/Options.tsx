import { useEffect } from "react";
import { useDispatch } from "react-redux";

import { BOARD_OPTIONS } from "../constants";
import { useSelector } from "../store";
import { GameParams } from "../types";
import DifficultyRow from "./DifficultyRow";

import { settingsUpdate } from "../slices/settings";

type OptionsProps = {
  startGame: () => void;
};

const OptionsPanel = ({ startGame }: OptionsProps) => {
  const settings = useSelector((state) => state.settings);
  const dispatch = useDispatch();

  const isValidCustomBoard = (settings: GameParams): boolean => {
    return settings.width * settings.height > settings.mineCount;
  };

  useEffect(() => {
    if (!isValidCustomBoard(settings)) {
      window.alert("Too many mines for the board dimensions");
      const newSettings = {
        ...settings,
        mineCount: settings.width * settings.height - 1,
      };
      dispatch(settingsUpdate(newSettings));
    }
  }, [settings, dispatch]);

  return (
    <>
      <div className="Options">
        <table id="option-table">
          <thead>
            <tr>
              <th></th>
              <th></th>
              <th>Width</th>
              <th>Height</th>
              <th>Mines</th>
            </tr>
          </thead>
          <tbody>
            {BOARD_OPTIONS.map((option) => (
              <tr key={option.difficulty}>
                <DifficultyRow option={option} />
              </tr>
            ))}
          </tbody>
        </table>
      </div>
      <button onClick={startGame}>New Game</button>
    </>
  );
};

export default OptionsPanel;
