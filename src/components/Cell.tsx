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
	const { icon, background } = cellIcon(cellType);

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
			}}
			onContextMenu={(e) => e.preventDefault()}
			onMouseDown={(e) => setClickDown(e.buttons)}
			onMouseUp={() => sendAction()}
		>
			{icon}
		</div>
	);
};

const cellIcon = (cellType: number): { icon: string; background: string } => {
	let icon: string;
	let background = "#7f8c8d";

	switch (cellType) {
		case -1:
			icon = "";
			background = "white";
			break;
		case 0:
			icon = "";
			break;
		case 1:
		case 2:
		case 3:
		case 4:
		case 5:
		case 6:
		case 7:
		case 8:
			icon = cellType.toString();
			break;
		case 9:
			icon = "💣";
			break;
		case -2:
			icon = "🚩";
			background = "white";
			break;
		default:
			throw new Error(`undefined cell type ${cellType}`);
	}
	return { icon, background };
};

export default Cell;
