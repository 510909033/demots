"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.demoMap = void 0;
class demoMap {
    constructor() {
        this.m = new Map();
    }
    set(k, v) {
        return this.m.set(k, v);
    }
    get(k) {
        return this.m.get(k);
    }
    getEntity() {
        const keys = this.m.keys();
        for (let key of keys) {
            console.log(`key: ${key}, value: ${this.m.get(key)}`);
        }
        return this.m;
    }
}
exports.demoMap = demoMap;
