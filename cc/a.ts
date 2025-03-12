// import { User } from "./src/pkg/domain/user.js";
// import { User as User2 } from "./src/pkg/domain/user.js";
import { UserRepository } from "./src/pkg/repository/user_repository.js";
import { Cat } from "./src/pkg/usecase/cate_usecase.js";
import { DemoCondition } from "./src/pkg/usecase/condition_usecase.js";
import { demoHttpJsonUsecase } from "./src/pkg/usecase/http_json_usecase.js";
import { DemoKeyof } from "./src/pkg/usecase/keyof_usecase.js";
import { demoMap } from "./src/pkg/usecase/map_usecase.js";
import { MultiHttpJsonUsecase } from "./src/pkg/usecase/multi_http_json_usecase.js";
import { writeFileContent } from "./src/pkg/usecase/require_usecase.js";
import { AgeType, UserUseCase } from "./src/pkg/usecase/user_usecase.js";
import { demoWsClientUsecase } from "./src/pkg/usecase/ws_client_usecase.js";
if (false) {
    const userUseCase = new UserUseCase(new UserRepository("test1", 2, 3));
    console.log(userUseCase.getUserInfo());
    userUseCase.UpdateUser({
        id: 1,
        name: 'test2',
        age: 18
    });
    console.log(userUseCase.getUserInfo());

    const {name1,age1} = userUseCase.demoObj();
    console.log(name1,age1);

    let at:AgeType = null;
    console.log(at);

    const cat:Cat.CatUseCase = new Cat.CatUseCase(new Cat.CatRepository("tom",2));
    console.log(cat.getCatInfo().age);

    // 示例：读取和写入文件
    // const filePath = './example.txt';
    // const content = 'Hello, world!';

    // writeFileContent(filePath, content);
    // const fileContent = readFileContent(filePath);
    // console.log('File content:', fileContent);

    const dm = new demoMap()
    console.log(dm.set("a",1))
    console.log(dm.get("a"))

    let demokeyof = new DemoKeyof()
    demokeyof.demo1();
    demokeyof.demo2();

    new DemoCondition()

    let demoHttpJson = new demoHttpJsonUsecase()
    demoHttpJson.getJsonData();

    let multiHttpJson = new MultiHttpJsonUsecase()
    multiHttpJson.getJsonData();
}
let ws = new demoWsClientUsecase('ws://localhost:8080/ws', 3000);
ws.connect();

console.log("over")
// setTimeout(() => {
//     ws.checkAndReConnect();
// }, 2000);

// 此处休眠 3秒
// (async () => {
//     // 此处休眠 3秒
//     await new Promise(resolve => setTimeout(resolve, 2000));

//  for (let i = 0; i < 10; i++) {
//     ws.sendMessage({name: "name"+ i , time: new Date().toString() });
// }
// })();


// dm.getEntity().forEach((value, key) => {
//     console.log(`Key: ${key}, Value: ${value}`);
// });



class demo1 implements UserRepositoryInterface {
    age: number;
    id: number;
    name: string;
    constructor(name: string, age: number, id?: number) {
        this.name = name;
        this.age = age;
        if (typeof id === 'undefined') {
            this.id = Math.random(); 
        } else {
            this.id = id;
        }
    }
    update(user: UserEntity): Resp {
        this.name = user.name;
        this.age = user.age;
        this.id = user.id;
        return { 
            status: 1,
            message: 'User updated successfully',
            HasError: ()=> true,
        };
    }

    getUserInfo(): UserEntity {
        return {
            id: this.id,
            name: this.name,
            age: this.age
        };
    }
  
 }

 let d:demo1 = new demo1("",1,2)














// // @ts-check
// import axios from 'axios';
// import { User } from './users/user.js';
// import MyStatus from './demo/enum.js';
// import { Direction } from './demo/enum.js';
// import Status from './demo/enum.js';
// import { StringValidator } from "./mod/moda.js";
// // import {numberRegexp, ZipCodeValidator} from './modb/modb.js';
// import * as validator from './modb/modb.js';
// // import { ZipCodeValidator as aaa } from './modb/modb.js';

// console.log((new validator.ZipCodeValidator()).isAcceptable("123"));

// // let z:ZipCodeValidator = new ZipCodeValidator();
// // console.log(z.isAcceptable('123'));

// // numberRegexp.test('123');
// // let s = new aaa();
// // console.log(s.isAcceptable('123'));

// // 打印当前时间函数
// function printCurrentTime(msg: string) {
//     const currentTime = new Date();
//     console.log(`Current time: ${currentTime} -- ${msg}`);
// }

// // sleep 2秒函数
// function sleep(ms: number) {
//     return new Promise(resolve => setTimeout(resolve, ms));
// }

// // 立即睡眠3秒钟函数
// async function sleepSync(seconds: number) {
//     await sleep(seconds * 1000);
// }

// function makeRequest() {
//     printCurrentTime('Starting the request');
//     axios.get('https://jsonplaceholder.typicode.com/todos/1')
//         .then(response => {
//             printCurrentTime('Request completed');
//             console.log(response.data);
//             // 遍历 response.headers
//             // for (const header in response.headers) {
//             //     console.log(`${header}: ${response.headers[header]}`);
//             // }
//         })
//         .catch(error => {
//             printCurrentTime('Error making request');
//             console.error('Error making request:', error);
//         });

//     printCurrentTime('start 2 s');
//     // sleep(2000);
//     sleepSync(4).then(() => {
//         printCurrentTime("now");
//     });
//     printCurrentTime('2秒后结束');
// }

// printCurrentTime('Starting the program');

// // 实例化 User 类
// const user = new User('Alice', 30);
// console.log(user.getUserInfo());

// makeRequest();
// printCurrentTime('Finishing the program');

// // sleep 2秒
// // await new Promise(resolve => setTimeout(resolve, 2000));

// printCurrentTime('over');

// interface Todo2 {
//     title: string;
//     description: string;
// }

// function updateTodo(todo: Todo2, fieldsToUpdate: Partial<Todo2>) {
//     return { ...todo, ...fieldsToUpdate };
// }

// const todo1 = {
//     title: "organize desk",
//     description: "clear clutter",
// };

// const todo2 = updateTodo(todo1, {
//     description: "throw out trash",
// });


// interface Props {
//     a?: number;
//     b?: string;
// }

// const obj: Props = { a: 5 };

// //   const obj2: Required<Props> = { a: 5 };
// const obj3: Partial<Props> = { a: 5 };
// const obj4: Readonly<Props> = { a: 5 };
// //   obj4.a = 3; //无法为“a”赋值，因为它是只读属性

// interface CatInfo {
//     age: number;
//     breed: string;
// }

// type CatName = "miffy" | "boris" | "mordred";

// const cats: Record<CatName, CatInfo> = {
//     miffy: { age: 10, breed: "Persian" },
//     boris: { age: 5, breed: "Maine Coon" },
//     mordred: { age: 16, breed: "British Shorthair" },
//     // "boris": { age: 5, breed: "Maine Coon" },
// };

// interface Todo {
//     title: string;
//     description: string;
//     completed: boolean;
// }

// type TodoPreview = Pick<Todo, "title" | "completed">;

// const todo: TodoPreview = {
//     title: "Clean room",
//     completed: false,
// };


// console.log(MyStatus.ACCEPTED);
// console.log(Direction.Down);
  