import * as readline from 'readline';

export class demoShellInputUsecase {
	private rl: readline.Interface;

	constructor() {
		this.rl = readline.createInterface({
			input: process.stdin,
			output: process.stdout
		});
	}

	public start() {
		this.rl.question('请输入字符串: ', (input: string) => {
			this.handleInput(input);
			this.rl.close();
		});
	}

	private handleInput(input: string) {
		console.log('您输入的字符串是:', input);
	}

    // 输入后不退出， 可以继续输入
    
    public start2() {
        this.rl.question('请输入字符串: ', (input: string) => {
            this.handleInput(input);
            this.start2();
        });
    }
    public async start3() {
        while (true) {
            const input: string = await this.askQuestion('请输入字符串: ');
            this.handleInput(input);
        }
    }

   
    private askQuestion(query: string): Promise<string> {
        return new Promise((resolve) => {
            this.rl.question(query, resolve);
        });
    }
}



