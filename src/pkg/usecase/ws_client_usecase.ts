// 帮我实现功能类 demoWsClientUsecase
//1. 连接ws服务器
//2. 进行收发json 消息
//3. 断开和异常情况重连
//4. 正确处理各种异常情况

export class demoWsClientUsecase {
	private ws: WebSocket | null = null;
	private url: string;
	private reconnectInterval: number;
	private reconnectTimeout: NodeJS.Timeout | null = null;
    
    private num = 0;

	constructor(url: string, reconnectInterval: number = 5000) {
		this.url = url;
        console.log("demoWsClientUsecase, ws.status: ", this.ws?.readyState);
		this.reconnectInterval = reconnectInterval;
        setInterval(() => {
            console.log("demoWsClientUsecase, ws.status: ", this.ws?.readyState);
        }, 1000);
	}

	public connect() {
        this.reconnectTimeout = null;
        this.ws = new WebSocket(this.url);
    
        this.ws.onopen = () => {
            console.log('WebSocket Connected to WebSocket server');
        };
    
        this.ws.onmessage = (event) => {
            try {
                console.log('WebSocket Received message:', event);
                const message = JSON.parse(event.data);
                this.handleMessage(message);
            } catch (error) {
                console.error('WebSocket Error parsing message:', error);
            }
        };
    
        this.ws.onclose = (event: CloseEvent) => {
            console.log('WebSocket connection closed:', event.code, event.reason);
            this.reconnect();
        };
    
        this.ws.onerror = (error: Event) => {
            this.num++;
            // console.log("WebSocket error:", this.num, error);
            console.log("WebSocket error:", this.num, this.ws?.readyState);

            switch (this.ws?.readyState) {
                case WebSocket.CONNECTING:
                    console.log("WebSocket is connecting...");
                    break;
                case WebSocket.OPEN:
                    console.log("WebSocket is open.");
                    break;
                case WebSocket.CLOSING:
                    console.log("WebSocket is closing...");
                    break;
                case WebSocket.CLOSED:
                    console.log("WebSocket is closed.");
                    break;
             }
            
            // 如果 ws 已经存在且状态不是已经关闭，则尝试关闭它
            // if (this.ws && this.ws.readyState !== WebSocket.CLOSED) {
            //     this.ws.close();
            // }
    
            // 延迟3秒后重连
            this.reconnect();
        };
    }

	private reconnect() {
		if (!this.reconnectTimeout) {
			this.reconnectTimeout = setTimeout(() => {
				console.log('Reconnecting to WebSocket server...');
				this.connect();
			}, this.reconnectInterval);
		}
	}

	private handleMessage(message: any) {
		// 处理接收到的消息
		console.log('WebSocket Received message:', message);
	}

	public sendMessage(message: any) {
		if (this.ws && this.ws.readyState === WebSocket.OPEN) {
			this.ws.send(JSON.stringify(message));
            console.log('WebSocket Sent message:', message);
		} else {
			console.error('WebSocket is not open. Unable to send message.');
		}
	}

	public close() {
		if (this.ws) {
			this.ws.close();
		}
	}
}