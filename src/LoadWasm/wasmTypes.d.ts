import { GameParams } from '../types'

declare global {
	export interface Window {
		Go: any;
		newGame: (gameParams: GameParams) => string;
		openCell: (x: number, y: number) => string;
	}
}

export { };
