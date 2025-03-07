// @ts-check
import axios from 'axios';
// 打印当前时间函数
function printCurrentTime(msg) {
    const currentTime = new Date();
    console.log(`Current time: ${currentTime} -- ${msg}`);
}
function makeRequest() {
    axios.get('https://jsonplaceholder.typicode.com/todos/1')
        .then(response => {
        console.log(response.data);
        // 遍历 response.headers
        for (const header in response.headers) {
            console.log(`${header}: ${response.headers[header]}`);
        }
    })
        .catch(error => {
        console.error('Error making request:', error);
    });
}
printCurrentTime('Starting the program');
makeRequest();
printCurrentTime('Finishing the program');
// sleep 2秒
await new Promise(resolve => setTimeout(resolve, 2000));
printCurrentTime('2秒结束');
