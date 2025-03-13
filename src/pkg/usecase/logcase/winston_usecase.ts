import winston from 'winston';

export class  DemoLogWinston {
    private logger: winston.Logger;
    
    constructor() {
        this.logger =   winston.createLogger({
            level: 'info',
            // format:winston.format.json(),
            format: winston.format.combine(
                winston.format.timestamp(),
                winston.format.json(),
                // winston.format.printf(({ timestamp, level, message }) => {
                //     return `${timestamp} ${level}: ${message}`;
                // })
            ),
            transports: [
                new winston.transports.Console(),
                new winston.transports.File({ filename: 'tsapp.log' })
            ]
        })
    }
    

    public info(message: string) {
        this.logger.info(message);
    }
    public warn(message: string) {
        this.logger.warn(message);
    }

    public demo() {
        const logger = winston.createLogger({
            level: 'info',
            format: winston.format.combine(
                winston.format.timestamp(),
                winston.format.printf(({ timestamp, level, message }) => {
                    return `${timestamp} ${level}: ${message}`;
                })
            ),
            transports: [
                new winston.transports.Console(),
                // new winston.transports.File({ filename: 'app.log' })
            ]
        })

        logger.log({
            level: 'info',
            message: 'Hello, Winston!',
            metadata: {
                userId: 123,
                userName: 'John Doe'
            }
        });
        
    }

}