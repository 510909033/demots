"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.assertIsIntMe = assertIsIntMe;
// 类型断言函数，检查一个数值是否为整数
function assertIsIntMe(value) {
    if (!Number.isInteger(value)) {
        throw new Error(`${value} 不是一个整数`);
    }
}
// 示例使用
const num = 5.5;
assertIsIntMe(num); // 编译时不会报错，但运行时会抛出错误
const intNum = num; // 如果 num 不是整数，运行时会抛出错误
