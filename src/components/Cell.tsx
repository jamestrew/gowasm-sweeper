import { CELL_SIZE } from "../constants";
import "../App.css";
import { useState } from "react";

type CellProps = {
	x: number;
	y: number;
	cellType: number;
	openCell: () => void;
	flagCell: () => void;
	chordedOpen: () => void;
};

const Cell = ({
	x,
	y,
	cellType,
	openCell,
	flagCell,
	chordedOpen,
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
	};

	return (
		<div
			className="Cell"
			style={{
				width: `${CELL_SIZE - 6}px`,
				height: `${CELL_SIZE - 6}px`,
				gridColumnStart: x + 1,
				gridRowStart: y + 1,
				background: background,
				fontWeight: "525",
				color: color,
			}}
			onContextMenu={(e) => e.preventDefault()}
			onMouseDown={(e) => setClickDown(e.buttons)}
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
	"0": { icon: "", background: "#c0c0c0", color: "" },
	"1": { icon: "1", background: "#c0c0c0", color: "#0000ff" },
	"2": { icon: "2", background: "#c0c0c0", color: "#008200" },
	"3": { icon: "3", background: "#c0c0c0", color: "#ff0000" },
	"4": { icon: "4", background: "#c0c0c0", color: "#000084" },
	"5": { icon: "5", background: "#c0c0c0", color: "#840000" },
	"6": { icon: "6", background: "#c0c0c0", color: "#008284" },
	"7": { icon: "7", background: "#c0c0c0", color: "#840084" },
	"8": { icon: "8", background: "#c0c0c0", color: "#757575" },
	"9": { icon: "💣", background: "#c0c0c0", color: "" },
};

export default Cell;
