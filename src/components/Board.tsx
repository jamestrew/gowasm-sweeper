import { CELL_SIZE } from "../constants";
import { CellPosition } from "../types";
import Cell from "./Cell";

import "../App.css";
import { useState } from "react";

type BoardProps = {
  board: number[][];
  setGame: (gameStr: string) => void;
};

const Board = ({ board, setGame }: BoardProps) => {
  const [hlPositions, setHlPositions] = useState<CellPosition[]>([]);

  const width = board[0].length;
  const height = board.length;

  const boardWidth = `${CELL_SIZE * width}px`;
  const boardHeight = `${CELL_SIZE * height}px`;

  const isHighlighted = (x: number, y: number): boolean => {
    return hlPositions.findIndex((pos) => pos.x === x && pos.y === y) !== -1;
  };

  return (
    <div className="Board" style={{ width: boardWidth, height: boardHeight }}>
      {board.map((row, i) =>
        row.map((_, j) => (
          <Cell
            x={j}
            y={i}
            key={i * width + j}
            cellType={board[i][j]}
            highlight={isHighlighted(j, i)}
            openCell={() => setGame(window.openCell(j, i))}
            flagCell={() => setGame(window.flagCell(j, i))}
            chordedOpen={() => setGame(window.chordedOpen(j, i))}
            setHighlights={(positions: CellPosition[]) => setHlPositions(positions)}
          />
        ))
      )}
    </div>
  );
};

export default Board;
