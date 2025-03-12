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

	constructor(url: string, reconnectInterval: number = 5000) {
		this.url = url;
		this.reconnectInterval = reconnectInterval;
		this.connect();
	}

	private connect() {
		this.ws = new WebSocket(this.url);

		this.ws.onopen = () => {
			console.log('Connected to WebSocket server');
			if (this.reconnectTimeout) {
				clearTimeout(this.reconnectTimeout);
				this.reconnectTimeout = null;
			}
		};

		this.ws.onmessage = (event) => {
			try {
				const message = JSON.parse(event.data);
				this.handleMessage(message);
			} catch (error) {
				console.error('Error parsing message:', error);
			}
		};

		this.ws.onclose = (event) => {
			console.log('WebSocket connection closed:', event);
			this.reconnect();
		};

		this.ws.onerror = (error) => {
			console.error('WebSocket error:', error);
			this.ws?.close();
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
		console.log('Received message:', message);
	}

	public sendMessage(message: any) {
		if (this.ws && this.ws.readyState === WebSocket.OPEN) {
			this.ws.send(JSON.stringify(message));
            console.log('Sent message:', message);
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