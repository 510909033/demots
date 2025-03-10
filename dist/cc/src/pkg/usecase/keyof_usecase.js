"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.DemoKeyof = void 0;
// 类demoKeyof ，提供 keyof 使用示例
class DemoKeyof {
    // 示例1
    demo1() {
        let p = 'age';
        console.log(p);
        let a = 1;
        let m = 'name';
        let m1 = 1;
    }
    // 示例2
    demo2() {
        // type Person = {
        //     name: string;
        //     age: number;
        // };
        let m = {
            name: true,
            age: "str1",
            arg2: "str",
            arg3: "str3",
        };
        for (let k in m) {
            let v = m[k];
            console.log(k, v);
        }
    }
}
exports.DemoKeyof = DemoKeyof;
