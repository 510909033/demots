// 定义 IntMe 类型为整数类型
export type IntMe = number & { __int__: void };

// 类型断言函数，检查一个数值是否为整数
export function assertIsIntMe(value: number): asserts value is IntMe {
    if (!Number.isInteger(value)) {
        throw new Error(`${value} 不是一个整数`);
    }
}

// 示例使用
const num: number = 5.5;
assertIsIntMe(num); // 编译时不会报错，但运行时会抛出错误
const intNum: IntMe = num; // 如果 num 不是整数，运行时会抛出错误
