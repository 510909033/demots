// @ts-check
import axios from 'axios';
// 打印当前时间函数
function printCurrentTime(msg) {
    const currentTime = new Date();
    console.log(`Current time: ${currentTime} -- ${msg}`);
}
// sleep 2秒函数
function sleep(ms) {
    new Promise(resolve => setTimeout(resolve, ms));
}
// 立即睡眠3秒钟函数
async function sleepSync(seconds) {
    await sleep(seconds);
}
function makeRequest() {
    printCurrentTime('Starting the request');
    axios.get('https://jsonplaceholder.typicode.com/todos/1')
        .then(response => {
        printCurrentTime('Request completed');
        console.log(response.data);
        // 遍历 response.headers
        // for (const header in response.headers) {
        //     console.log(`${header}: ${response.headers[header]}`);
        // }
    })
        .catch(error => {
        printCurrentTime('Error making request');
        console.error('Error making request:', error);
    });
    printCurrentTime('start 2 s');
    //  sleep(2000);
    sleepSync(4).then(() => {
        printCurrentTime("now");
    });
    printCurrentTime('2秒后结束');
}
printCurrentTime('Starting the program');
makeRequest();
printCurrentTime('Finishing the program');
// sleep 2秒
// await new Promise(resolve => setTimeout(resolve, 2000));
printCurrentTime('over');
