export class UserRepository implements UserRepositoryInterface {
    name: string;
    age: number;
    id: number;
    private _phone: string = "";


    constructor(name: string, age: number, id?: number) {
        this.name = name;
        this.age = age;
        if (typeof id === 'undefined') {
            this.id = Math.random(); 
        } else {
            this.id = id;
        }
    }
    update(user: UserEntity): Resp {
        this.name = user.name;
        this.age = user.age;
        this.id = user.id;
        return {
            status: 0,
            message: 'User updated successfully',
            HasError: ()=> true,
        }
    }

    getUserInfo(): string {
        return `Name: ${this.name}, Age: ${this.age}`;
    }

    // 设置名称
        setName(name: string): void {
            this.name = name;
        }
        
    setPhone(phone: string): void {
        this._phone = phone;
    }
    
}
