// @ts-check
import axios from 'axios';
// 打印当前时间函数
function printCurrentTime(msg:string) {
    const currentTime = new Date();
    console.log(`Current time: ${currentTime} -- ${msg}`);
}

function makeRequest() {
    printCurrentTime('Starting the request');
    axios.get('https://jsonplaceholder.typicode.com/todos/1')
        .then(response => {
            printCurrentTime('Request completed');
            console.log(response.data);
            // 遍历 response.headers
            for (const header in response.headers) {
                console.log(`${header}: ${response.headers[header]}`);
            }
        })
        .catch(error => {
            printCurrentTime('Error making request');
            console.error('Error making request:', error);
        });

    printCurrentTime('start 2 s')
    await new Promise(resolve => setTimeout(resolve, 2000));
    printCurrentTime('2秒后结束');

}

printCurrentTime('Starting the program');
makeRequest();
printCurrentTime('Finishing the program');

// sleep 2秒
// await new Promise(resolve => setTimeout(resolve, 2000));

printCurrentTime('2秒结束')