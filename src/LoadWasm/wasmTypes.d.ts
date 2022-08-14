import { GameParams } from '../types'

declare global {
	export interface Window {
		Go: any;
		newGame: (gameParams: GameParams) => string;
		openCell: (x: number, y: number) => string;
		flagCell: (x: number, y: number) => string;
		chordedOpen: (x: number, y: number) => string;
		getState: () => string;
	}
}

export { };
