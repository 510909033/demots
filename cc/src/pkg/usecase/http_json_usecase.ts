import axios from 'axios'; // 更新导入语句

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

// 获取 JsonResponse 中 cpkey 的类型
type CpkeyType = JsonResponse['cpkey'];

export class demoHttpJsonUsecase {
    getJsonData() {
        // 获取json数据并解析
        axios.get<ApiData | ErrApiData>('xxx')
            .then((res) => {
                // 检查 res.data 是否包含 'data' 属性
                if ('data' in res.data) {
                    const data: JsonResponse = res.data.data;
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

                     // 使用 CpkeyType 确保类型一致
                     const cpkey: CpkeyType = data.cpkey;
                     console.log('CPKey:', cpkey);
                     
                    // typeof
                    if (typeof data.cpkey == "number") {
                        console.log('CPKey is a number');
                    } else {
                        console.log('CPKey is not a number');
                        throw new Error('CPKey is not a number');
                    }

                   
                    
                } else {
                    const errorData: ErrApiData = res.data;
                    console.error('获取json数据失败', errorData.msg);
                    console.error('Client IP:', errorData.client_ip);
                }
            })
            .catch((err) => {
                console.error('获取json数据失败', err);
            });
    }
}