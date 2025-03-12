import axios from 'axios';

interface JsonResponse {
    cpid: string;
    cpkey: number;
    api_domain: string;
    admin_domain: string;
    customer_web_socket_domain: string;
    ims_tcp_domain: string;
    ims_web_socket_domain: string;
    logo: string;
    cp_name: string;
    gw_k8s_addr: string;
    k8s_addr: string;
}

interface ApiData {
    data: JsonResponse;
    code: number;
}

interface ErrApiData {
    msg: string;
    client_ip: string;
}

export class MultiHttpJsonUsecase {
    async getJsonData() {
        const url = 'xxx';
        const requests = Array(10).fill(axios.get<ApiData | ErrApiData>(url));

        try {
            const responses = await Promise.all(requests);
            responses.forEach((res, index) =>  {
                if ('data' in res.data) {
                    const data: JsonResponse = res.data.data;
                    console.log(`请求 ${index + 1} 获取json数据并解析成功`, data);
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
                } else {
                    const errorData: ErrApiData = res.data;
                    console.error(`请求 ${index + 1} 获取json数据失败`, errorData.msg);
                    console.error('Client IP:', errorData.client_ip);
                }
            });
        } catch (err) {
            console.error('获取json数据失败', err);
        }
    }
}

// 创建实例并调用方法
// const multiHttpJsonUsecase = new MultiHttpJsonUsecase();
// multiHttpJsonUsecase.getJsonData();
