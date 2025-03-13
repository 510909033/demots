import { Worker, isMainThread, parentPort, workerData } from 'worker_threads';

if (isMainThread) {
    // 主线程逻辑
    console.log('主线程启动');

    const numThreads = 3;
    const promises = [];

    for (let i = 0; i < numThreads; i++) {
        promises.push(new Promise((resolve, reject) => {
            const worker = new Worker(__filename, {
                workerData: { start: i * 10 + 1, end: (i + 1) * 10 }
            });

            worker.on('message', (result) => {
                console.log(`工作线程 ${i} 返回结果: ${result}`);
                resolve(result);
            });

            worker.on('error', (error) => {
                console.error(`工作线程 ${i} 错误: ${error}`);
                reject(error);
            });

            worker.on('exit', (code) => {
                if (code !== 0) {
                    console.error(`工作线程 ${i} 退出，退出码: ${code}`);
                } else {
                    console.log(`工作线程 ${i} 正常退出`);
                }
            });
        }));
    }

    // 等待所有工作线程完成
    Promise.all(promises).then((results) => {
        console.log('所有工作线程已完成');
        console.log('结果:', results);
    }).catch((error) => {
        console.error('工作线程执行出错:', error);
    });
} else {
    // 工作线程逻辑
    const { start, end } = workerData;
    let sum = 0;

    for (let i = start; i <= end; i++) {
        sum += i;
    }

    // 向主线程发送结果，确保 parentPort 不为 null
    if (parentPort) {
        parentPort.postMessage(sum);
    } else {
        console.error('当前不在工作线程中，无法发送消息');
    }
}