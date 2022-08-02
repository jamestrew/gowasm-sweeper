declare global {
  export interface Window {
    Go: any;
    myGolangFunction: (a: number, b: number) => number;
  }
}

export {};
