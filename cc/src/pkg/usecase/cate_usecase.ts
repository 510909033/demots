
export namespace Cat {
    
     interface CatEntity {
        name: string;
        age: number;
    }
    export class CatRepository {
        name: string;
        age: number;
        constructor(name: string, age: number) {
            this.name = name;
            this.age = age;
        }
    }
    export class CatUseCase {
        catRepository: CatRepository;
        constructor(catRepository: CatRepository) {
            this.catRepository = catRepository;
        }
        getCatInfo(): CatEntity {
            return {
                name: this.catRepository.name,
                age: this.catRepository.age
            };
        }
    }
}
