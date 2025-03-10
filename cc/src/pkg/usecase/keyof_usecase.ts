// 类demoKeyof ，提供 keyof 使用示例
export class DemoKeyof {
    // 示例1
    demo1() {
        type Person = {
            name: string;
            age: number;
        };
        type P = keyof Person;
        let p: P = 'age';
        
        console.log(p);

        type Arrayish = { [n: number]: unknown };
        type A = keyof Arrayish;
        let a: A = 1;
        // let a1: A = "unknown";

        type Mapish = { [k: string]: boolean|number, arg2:number};
        type M = keyof Mapish;
        let m: M = 'name';
        let m1: M = 1;
    }

    // 示例2
    demo2() {
        // type Person = {
        //     name: string;
        //     age: number;
        // };

        type Mapish = { [k: string]: boolean | string, arg2: string };
        let m :Mapish = {
            name: true,
            age: "str1",
            arg2: "str",
            arg3:"str3",
        };
        for (let k in m) {
            let v = m[k];
            console.log(k, v);
        }
    }
}
