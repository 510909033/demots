export class UserUseCase {
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
