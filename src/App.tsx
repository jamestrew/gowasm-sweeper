import { useEffect, useState } from "react";
import "./App.css";
import { DEFAULT_SETTINGS } from "./constants";
import { GameState } from './types'
// import Options from "./components/Options";

import { DEFAULT_SETTINGS } from "./constants";
import { GameData } from "./types";
import { boardDimensions, gameObj } from "./utils";
// import Options from "./components/Options";

const CELL_SIZE = 30;

function App() {
  const [game, setGame] = useState<GameData>();

  useEffect(() => {
    console.log("App rendered");
    setGame(gameObj(window.newGame(DEFAULT_SETTINGS)));
  }, []);

  const { width, height } = boardDimensions(game);
  const board = game?.board;

  const boardWidth = `${CELL_SIZE * width}px`;
  const boardHeight = `${CELL_SIZE * height}px`;

  return (
    <div className="App">
      <button
        onClick={() =>
          setGame(
            gameObj(window.newGame({ ...DEFAULT_SETTINGS, difficulty: 3 }))
          )
        }
      >
        New Game
      </button>
      {board && (
        <div
          className="Board"
          style={{ width: boardWidth, height: boardHeight }}
        >
          {board.map((row, i) =>
            row.map((_, j) => (
              <Cell
                x={j}
                y={i}
                key={i * width + j}
                cellType={board[i][j]}
                openCell={() => setGame(gameObj(window.openCell(j, i)))}
              />
            ))
          )}
        </div>
      )}
    </div>
  );
}

type CellProps = {
  x: number;
  y: number;
  cellType: number;
  openCell: () => void;
};

const Cell = ({ x, y, cellType, openCell }: CellProps) => {
  let cellIcon: string;

  switch (cellType) {
    case -1:
    case 0:
      cellIcon = "";
      break;
    case 1:
    case 2:
    case 3:
    case 4:
    case 5:
    case 6:
    case 7:
    case 8:
      cellIcon = cellType.toString();
      break;
    case 9:
      cellIcon = "💣";
      break;
    case -2:
      cellIcon = "🚩";
      break;
    default:
      throw new Error(`undefined cell type ${cellType}`);
  }

  return (
    <div
      className="Cell"
      style={{
        width: `${CELL_SIZE - 6}px`,
        height: `${CELL_SIZE - 6}px`,
        gridColumnStart: x + 1,
        gridRowStart: y + 1,
        background: `${cellType === -1 ? "white" : " #7f8c8d "}`,
      }}
      onClick={openCell}
    >
      {cellIcon}
    </div>
  );
};

export default App;
