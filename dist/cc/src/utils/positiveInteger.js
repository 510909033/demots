"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.isPositiveInteger = isPositiveInteger;
// 类型守卫函数，检查一个数值是否为正整数
function isPositiveInteger(value) {
    return Number.isInteger(value) && value > 0;
}
// 示例使用
const num = 5;
if (isPositiveInteger(num)) {
    console.log(`${num} 是一个正整数`);
}
else {
    console.log(`${num} 不是一个正整数`);
}
