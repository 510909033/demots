"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.UserRepository = void 0;
class UserRepository {
    constructor(name, age, id) {
        this._phone = "";
        this.name = name;
        this.age = age;
        if (typeof id === 'undefined') {
            this.id = Math.random();
        }
        else {
            this.id = id;
        }
    }
    update(user) {
        this.name = user.name;
        this.age = user.age;
        this.id = user.id;
        return {
            status: 0,
            message: 'User updated successfully',
            HasError: () => true,
        };
    }
    getUserInfo() {
        return `Name: ${this.name}, Age: ${this.age}`;
    }
    // 设置名称
    setName(name) {
        this.name = name;
    }
    setPhone(phone) {
        this._phone = phone;
    }
}
exports.UserRepository = UserRepository;
