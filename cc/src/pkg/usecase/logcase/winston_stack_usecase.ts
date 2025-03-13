import winston from 'winston';


export class DemoLogWinstonStack {
    private logger: winston.Logger;

    constructor() {
        this.logger =  winston.createLogger({
            level: 'info',
            format: winston.format.combine(
                winston.format.timestamp(),
                
            //     winston.format((info) => {
            //         const stack = new Error().stack;
            //         console.log('stack', stack);
            //         if (stack) {
            //             const lines = stack.split('\n');
            //             if (lines.length > 2) {
            //                 const caller = lines[2].trim();
            //                 const match = caller.match(/at\s+.*\((.*):(\d+):(\d+)\)/);
            //                 if (match) {
            //                     info.file = match[1];
            //                     info.line = match[2];
            //                 }
            //             }
            //         }
            //         return info;
            //     })(),
            //     winston.format.printf(({ timestamp, level, message, file, line }) => {
            //         return `${timestamp} ${level}: ${message} [${file}:${line}]`;
            //     }),
                winston.format.json(),
            ),
            transports: [
                new winston.transports.Console(),
                // new winston.transports.File({ filename: 'tsapp.log' })
            ]
        });
    }

    public info(message: string) {
        this.logger.info(message);
    }

    public warn(message: string) {
        this.logger.warn(message);
    }

    public calculateDurationInMicroseconds(start: bigint): number {
        const end = process.hrtime.bigint();
        const durationInNanoseconds = end - start;
        const durationInMicroseconds = Number(durationInNanoseconds) / 1000;
        return durationInMicroseconds;
    }

    public error(message: string) {
        const startTime = process.hrtime.bigint();

        const stack = new Error().stack;

        let durationInMicroseconds = this.calculateDurationInMicroseconds(startTime);
        console.info(`耗时：newError().stack:  ${durationInMicroseconds} microseconds`);
        
        let file = 'unknown';
        let line = 'unknown';
        if (stack) {
            const lines = stack.split('\n');
            if (lines.length > 2) {
                const caller = lines[2].trim();
                const match = caller.match(/at\s+.*\((.*):(\d+):(\d+)\)/);
                if (match) {
                    file = match[1];
                    line = match[2];
                }
            }
        }

        durationInMicroseconds = this.calculateDurationInMicroseconds(startTime);
        console.info(`耗时：解析stack:  ${durationInMicroseconds} microseconds`);
        
        this.logger.error({
            message:message,
            file:file,
            line:line,
            obj:{
                "name":"testname",
                "age":20,
            },
        });

        durationInMicroseconds = this.calculateDurationInMicroseconds(startTime);
        console.info(`耗时：执行err记录日志:  ${durationInMicroseconds} microseconds`);
    }
}