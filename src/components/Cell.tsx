import { CELL_SIZE } from "../constants";
import "../App.css";

type CellProps = {
	x: number;
	y: number;
	cellType: number;
	openCell: () => void;
	flagCell: () => void;
};

const Cell = ({ x, y, cellType, openCell, flagCell }: CellProps) => {
	const handleRightClick = (
		e: React.MouseEvent<HTMLDivElement, MouseEvent>
	) => {
		e.preventDefault();
		flagCell();
	};

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
			onContextMenu={(e) => handleRightClick(e)}
		>
			{cellIcon(cellType)}
		</div>
	);
};

const cellIcon = (cellType: number): string => {
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
	return cellIcon;
};

export default Cell;
