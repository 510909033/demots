export var Cat;
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
})(Cat || (Cat = {}));
