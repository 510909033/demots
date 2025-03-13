import { Worker, isMainThread, parentPort, workerData } from 'worker_threads';

if (isMainThread) {
    // 主线程逻辑
    console.log('主线程启动');

    // 创建工作线程
    const worker = new Worker(__filename, {
        workerData: { start: 1, end: 10 }
    });

    // 监听工作线程的消息
    worker.on('message', (result) => {
        console.log(`工作线程返回结果: ${result}`);
    });

    // 监听工作线程的错误
    worker.on('error', (error) => {
        console.error(`工作线程错误: ${error}`);
    });

    // 监听工作线程的退出
    worker.on('exit', (code) => {
        if (code !== 0) {
            console.error(`工作线程退出，退出码: ${code}`);
        } else {
            console.log('工作线程正常退出');
        }
    });
} else {
    // 工作线程逻辑
    const { start, end } = workerData;
    let sum = 0;

    for (let i = start; i <= end; i++) {
        sum += i;
    }

    // 暂停2秒

    setTimeout(() => {
       // 向主线程发送结果，确保 parentPort 不为 null
    if (parentPort) {
        parentPort.postMessage(sum);
        parentPort.postMessage(sum);
    } else {
        console.error('当前不在工作线程中，无法发送消息');
    }
        
    }, 3000);
    

    setTimeout(() => {
   
         
     }, 3000);
}