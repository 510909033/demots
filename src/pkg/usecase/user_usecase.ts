export type AgeType = number|null
export class UserUseCase implements UserUseCaseInterface {
  private userRepository: UserRepositoryInterface;

  constructor(userRepository: UserRepositoryInterface) {
    this.userRepository = userRepository;
  }
     UpdateUser(user: UserEntity): Resp {
           const updatedUser =  this.userRepository.update(user);
           return updatedUser
    }

  getUserInfo(): string {
    return this.userRepository.name
  }

  demoObj():{name1:string, age1:number} {
    return {
      name1: this.userRepository.name,
      age1: this.userRepository.age
    }
  }

}