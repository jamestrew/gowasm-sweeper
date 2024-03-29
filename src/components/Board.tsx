import { CELL_SIZE } from "../constants";
import { CellPosition, GameData } from "../types";
import Cell from "./Cell";

import "../App.css";
import { useState } from "react";

type BoardProps = {
  game: GameData;
};

const Board = ({ game }: BoardProps) => {
  const [hlPositions, setHlPositions] = useState<CellPosition[]>([]);

  const width = game.board[0].length;
  const height = game.board.length;

  const boardWidth = `${CELL_SIZE * width}px`;
  const boardHeight = `${CELL_SIZE * height}px`;

  const isHighlighted = (x: number, y: number): boolean => {
    return hlPositions.findIndex((pos) => pos.x === x && pos.y === y) !== -1;
  };

  return (
    <div className="Board" style={{ width: boardWidth, height: boardHeight }}>
      {game.board.map((row, i) =>
        row.map((_, j) => (
          <Cell
            x={j}
            y={i}
            key={i * width + j}
            cellType={game.board[i][j]}
            highlight={isHighlighted(j, i)}
            setHighlights={(positions: CellPosition[]) => setHlPositions(positions)}
          />
        ))
      )}
    </div>
  );
};

export default Board;
