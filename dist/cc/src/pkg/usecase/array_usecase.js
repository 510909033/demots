"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.demoArrayUseCase = void 0;
class demoArrayUseCase {
    constructor() {
        this.arr = new Array();
    }
    getArr() {
        return this.arr;
    }
    setArr(arr) {
        this.arr = arr;
    }
    getArrLength() {
        return this.arr.length;
    }
    addArr(str) {
        this.arr.push(str);
    }
}
exports.demoArrayUseCase = demoArrayUseCase;
