"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.DemoCondition = void 0;
//   type Example1 = Dog extends Animal ? number : string;
//   type Example2 = RegExp extends Animal ? number : string;
class DemoCondition {
    demo1() {
        let e1 = 1;
    }
}
exports.DemoCondition = DemoCondition;
function createLabel(idOrName) {
    if (typeof idOrName === "string") {
        return { name: idOrName };
    }
    return { id: idOrName };
}
//   type IntegerOrString<T extends number | string> = T extends number ? Integer<T> : T;
//   let i2:IntegerOrString<number> = 1.2;
//   console.log("-----------",i2);
