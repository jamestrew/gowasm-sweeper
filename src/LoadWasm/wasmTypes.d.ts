type board = string;

declare global {
	export interface Window {
		Go: any;
		newGame: (
			difficulty: number,
			width: number,
			height: number,
			mines: number
		) => board;
		openCell: (x: number, y: number) => board;
	}
}

export { };
