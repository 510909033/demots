
// const fs = require('fs');
// const path = require('path');
import * as fs from 'fs';
import * as path from 'path';

// 读取文件内容的函数
function readFileContent(filePath: string): string {
    try {
        const absolutePath = path.resolve(__dirname, filePath);
        const data = fs.readFileSync(absolutePath, 'utf8');
        return data;
    } catch (err) {
        console.error('Error reading file:', err);
        return '';
    }
}

// 写入文件内容的函数
function writeFileContent(filePath: string, content: string): void {
    try {
        const absolutePath = path.resolve(__dirname, filePath);
        console.log('absolutePath',absolutePath)
        fs.writeFileSync(absolutePath, content, 'utf8');
        console.log('File written successfully');
    } catch (err) {
        console.error('Error writing file:', err);
    }
}

// // 示例：读取和写入文件
// const filePath = './example.txt';
// const content = 'Hello, world!';

// writeFileContent(filePath, content);
// const fileContent = readFileContent(filePath);
// console.log('File content:', fileContent);

// 导出 
export {
    readFileContent,
    writeFileContent
}