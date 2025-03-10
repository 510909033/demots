"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.assertIsInt = assertIsInt;
exports.assertIsFloat = assertIsFloat;
// 类型守卫函数，检查一个数值是否为整数
function assertIsInt(value) {
    if (!Number.isInteger(value)) {
        throw new Error(`${value} 不是一个整数`);
    }
}
// 类型守卫函数，检查一个数值是否为浮点数
function assertIsFloat(value) {
    if (Number.isInteger(value)) {
        throw new Error(`${value} 不是一个浮点数`);
    }
}
// 示例使用
const num1 = 5.5;
assertIsFloat(num1); // 编译时不会报错，但运行时会抛出错误
const floatNum = num1; // 如果 num1 不是浮点数，运行时会抛出错误
const num2 = 5;
assertIsInt(num2); // 编译时不会报错，但运行时会抛出错误
const intNum = num2; // 如果 num2 不是整数，运行时会抛出错误
