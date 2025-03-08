export class User {
    constructor(name, age, id) {
        this.name = name;
        this.age = age;
        if (typeof id === 'undefined') {
            this.id = Math.random();
        }
        else {
            this.id = id;
        }
    }
    getUserInfo() {
        return `Name: ${this.name}, Age: ${this.age}`;
    }
}
