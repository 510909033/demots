"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Cat = void 0;
var Cat;
(function (Cat) {
    class CatRepository {
        constructor(name, age) {
            this.name = name;
            this.age = age;
        }
    }
    Cat.CatRepository = CatRepository;
    class CatUseCase {
        constructor(catRepository) {
            this.catRepository = catRepository;
        }
        getCatInfo() {
            return {
                name: this.catRepository.name,
                age: this.catRepository.age
            };
        }
    }
    Cat.CatUseCase = CatUseCase;
})(Cat || (exports.Cat = Cat = {}));
