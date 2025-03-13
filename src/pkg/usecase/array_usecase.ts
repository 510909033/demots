
export class demoArrayUseCase {
  arr: Array<string>;
  constructor() {
    this.arr = new Array<string>();
  }
  getArr(): Array<string> {
    return this.arr;
  }
  setArr(arr: Array<string>): void {
    this.arr = arr;
  }
  getArrLength(): number {
    return this.arr.length;
  }
  addArr(str: string): void{
    
    
    this.arr.push(str);
  }
}