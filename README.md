## 运行 TypeScript 文件

要运行 `cc/a.ts` 文件，请按照以下步骤操作：

1. 确保你已经安装了 Node.js 和 npm。
2. 在项目根目录下运行以下命令来安装 TypeScript 和 Axios：
    ```sh
    npm install typescript axios
    ```
3. 创建一个 `tsconfig.json` 文件来配置 TypeScript 编译选项：
    ```json
    {
        "compilerOptions": {
            "target": "ES6",
            "module": "commonjs",
            "outDir": "./dist",
            "rootDir": "./",
            "strict": true
        },
        "include": [
            "cc/**/*.ts"
        ]
    }
    ```
4. 编译 TypeScript 文件：
    ```sh
    npx tsc
    ```
5. 运行编译后的 JavaScript 文件：
    ```sh
    node dist/cc/a.js
    ```

这样你就可以运行 `cc/a.ts` 文件了。


npm install -g live-server

------
npm install --save-dev webpack webpack-cli webpack-dev-server

创建 Webpack 配置文件 (webpack.config.js)：

const path = require('path');

module.exports = {
    entry: './dist/src/main.js',
    output: {
        filename: 'bundle.js',
        path: path.resolve(__dirname, 'dist'),
        libraryTarget: 'var',
        library: 'App'
    },
    mode: 'development',
    target: 'web'
};

打包代码：

bash
npx webpack
在 HTML 中引入打包后的文件：

html
<script src="./dist/bundle.js"></script>