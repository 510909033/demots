"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.UserUseCase = void 0;
class UserUseCase {
    constructor(userRepository) {
        this.userRepository = userRepository;
    }
    UpdateUser(user) {
        const updatedUser = this.userRepository.update(user);
        return updatedUser;
    }
    getUserInfo() {
        return this.userRepository.name;
    }
    demoObj() {
        return {
            name1: this.userRepository.name,
            age1: this.userRepository.age
        };
    }
}
exports.UserUseCase = UserUseCase;
