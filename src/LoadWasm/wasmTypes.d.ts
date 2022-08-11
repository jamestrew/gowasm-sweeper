type board = string;

declare global {
	export interface Window {
		Go: any;
		newGame: (difficulty: number) => board;
		openCell: (x: number, y: number) => board;
	}
}

export { };
