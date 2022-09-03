import { CELL_SIZE } from "../constants";
import "../App.css";
import { useState } from "react";
import { CellPosition } from "../types";

const OPEN_COLOR = "#c0c0c0";

type CellProps = {
  x: number;
  y: number;
  cellType: number;
  highlight: boolean;
  openCell: () => void;
  flagCell: () => void;
  chordedOpen: () => void;
  setHighlights: (positions: CellPosition[]) => void;
};

const Cell = ({
  x,
  y,
  cellType,
  highlight,
  openCell,
  flagCell,
  chordedOpen,
  setHighlights,
}: CellProps) => {
  const [clickDown, setClickDown] = useState<number>();
  const { icon, background, color } = cellIcon[cellType];

  const sendAction = () => {
    switch (clickDown) {
      case 1:
        openCell();
        break;
      case 2:
        flagCell();
        break;
      case 3:
        chordedOpen();
        break;
    }
    setHighlights([]);
  };

  const handleMouseDown = (buttons: number) => {
    setClickDown(buttons);
    if (buttons === 1) setHighlights([{ x, y }]);
    else if (buttons === 3) setHighlights(JSON.parse(window.cellNeighbors(x, y)));
  };

  return (
    <div
      className="Cell"
      style={{
        width: `${CELL_SIZE - 6}px`,
        height: `${CELL_SIZE - 6}px`,
        gridColumnStart: x + 1,
        gridRowStart: y + 1,
        background: highlight && cellType !== -2 ? OPEN_COLOR : background,
        fontWeight: "525",
        color: color,
      }}
      onContextMenu={(e) => e.preventDefault()}
      onMouseDown={(e) => handleMouseDown(e.buttons)}
      onMouseUp={() => sendAction()}
    >
      {icon}
    </div>
  );
};

const cellIcon: {
  [cellType: string]: { icon: string; background: string; color: string };
} = {
  "-2": { icon: "🚩", background: "white", color: "" },
  "-1": { icon: "", background: "white", color: "" },
  "0": { icon: "", background: OPEN_COLOR, color: "" },
  "1": { icon: "1", background: OPEN_COLOR, color: "#0000ff" },
  "2": { icon: "2", background: OPEN_COLOR, color: "#008200" },
  "3": { icon: "3", background: OPEN_COLOR, color: "#ff0000" },
  "4": { icon: "4", background: OPEN_COLOR, color: "#000084" },
  "5": { icon: "5", background: OPEN_COLOR, color: "#840000" },
  "6": { icon: "6", background: OPEN_COLOR, color: "#008284" },
  "7": { icon: "7", background: OPEN_COLOR, color: "#840084" },
  "8": { icon: "8", background: OPEN_COLOR, color: "#757575" },
  "9": { icon: "💣", background: OPEN_COLOR, color: "" },
};

export default Cell;
