
// 定义 Int 类型为整数类型
export type Int = number & { __int__: void };

// 定义 Float 类型为浮点数类型
export type Float = number & { __float__: void };

// 类型守卫函数，检查一个数值是否为整数
export function assertIsInt(value: number): asserts value is Int {
    if (!Number.isInteger(value)) {
        throw new Error(`${value} 不是一个整数`);
    }
}

// 类型守卫函数，检查一个数值是否为浮点数
export function assertIsFloat(value: number): asserts value is Float {
    if (Number.isInteger(value)) {
        throw new Error(`${value} 不是一个浮点数`);
    }
}

// 示例使用
const num1: number = 5.5;
assertIsFloat(num1); // 编译时不会报错，但运行时会抛出错误
const floatNum: Float = num1; // 如果 num1 不是浮点数，运行时会抛出错误

const num2: number = 5;
assertIsInt(num2); // 编译时不会报错，但运行时会抛出错误
const intNum: Int = num2; // 如果 num2 不是整数，运行时会抛出错误
