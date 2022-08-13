import { useState } from "react";

import "./App.css";
import { DEFAULT_SETTINGS } from "./constants";
import { GameState } from './types'
// import Options from "./components/Options";

const CELL_SIZE = 30;

function App() {
  const [game, setGame] = useState<GameState>(
    JSON.parse(
      window.newGame(
        DEFAULT_SETTINGS.difficulty,
        DEFAULT_SETTINGS.width,
        DEFAULT_SETTINGS.height,
        DEFAULT_SETTINGS.mineCount
      )
    )
  );

	const width = board[0]?.length ?? 0;
	const height = board.length;

  const width = game.Board[0]?.length ?? 0;
  const height = game.Board.length;

  const boardWidth = `${CELL_SIZE * width}px`;
  const boardHeight = `${CELL_SIZE * height}px`;

  return (
    <div className="App">
      <div className="game">
        <div
          className="Board"
          style={{ width: boardWidth, height: boardHeight }}
        >
          {game.Board.map((row, i) =>
            row.map((_, j) => (
              <Cell
                x={j}
                y={i}
                key={i * width + j}
                cellType={game.Board[j][i]}
                openCell={() => setGame(JSON.parse(window.openCell(j, i)))}
              />
            ))
          )}
        </div>
      </div>
      {/* FIX: styling causing issues with clicking cells */}
      {/* <Options */}
      {/*   onNewGame={(settings) => */}
      {/*     setBoard( */}
      {/*       JSON.parse( */}
      {/*         window.newGame( */}
      {/*           settings.difficulty, */}
      {/*           settings.width, */}
      {/*           settings.height, */}
      {/*           settings.mineCount */}
      {/*         ) */}
      {/*       ) */}
      {/*     ) */}
      {/*   } */}
      {/* /> */}
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
  return (
    <div
      className="Cell"
      style={{
        width: `${CELL_SIZE - 6}px`,
        height: `${CELL_SIZE - 6}px`,
        gridColumnStart: x + 1,
        gridRowStart: y + 1,
        background: "white",
      }}
      onClick={openCell}
    >
      {`${cellType}`}
    </div>
  );
};

export default App;
