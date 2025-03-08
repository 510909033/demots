export class User {
    constructor(name, age) {
        this.name = name;
        this.age = age;
    }
    getUserInfo() {
        return `Name: ${this.name}, Age: ${this.age}`;
    }
}
