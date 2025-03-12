"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.demoHttpJsonUsecase = void 0;
const axios_1 = __importDefault(require("axios")); // 更新导入语句
class demoHttpJsonUsecase {
    getJsonData() {
        // 获取json数据并解析
        axios_1.default.get('https://cn-admin-test.ruixueyun.com/api/v1/publicadminapi/app/cpinfo?t=893c6794-0b34-4400-a2a2-49cd4e76d741')
            .then((res) => {
            // 检查 res.data 是否包含 'data' 属性
            if ('data' in res.data) {
                const data = res.data.data;
                console.log('获取json数据并解析成功', data);
                console.log('CPID:', data.cpid);
                console.log('CPKey:', data.cpkey);
                console.log('API Domain:', data.api_domain);
                console.log('Admin Domain:', data.admin_domain);
                console.log('Customer WebSocket Domain:', data.customer_web_socket_domain);
                console.log('IMS TCP Domain:', data.ims_tcp_domain);
                console.log('IMS WebSocket Domain:', data.ims_web_socket_domain);
                console.log('Logo:', data.logo);
                console.log('CP Name:', data.cp_name);
                console.log('GW K8S Addr:', data.gw_k8s_addr);
                console.log('K8S Addr:', data.k8s_addr);
                // typeof
                if (typeof data.cpkey == "string") {
                    console.log('CPKey is a string');
                }
                else {
                    console.log('CPKey is not a string');
                }
            }
            else {
                const errorData = res.data;
                console.error('获取json数据失败', errorData.msg);
                console.error('Client IP:', errorData.client_ip);
            }
        })
            .catch((err) => {
            console.error('获取json数据失败', err);
        });
    }
}
exports.demoHttpJsonUsecase = demoHttpJsonUsecase;
