import { useEffect, useState } from "react";
import "./App.css";

import { DEFAULT_SETTINGS, CELL_SIZE } from "./constants";
import { GameData } from "./types";
import { boardDimensions, gameObj } from "./utils";
import Cell from './components/Cell'


function App() {
  const [game, setGame] = useState<GameData>();

  useEffect(() => {
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


export default App;
