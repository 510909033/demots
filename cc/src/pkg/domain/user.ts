interface UserRepositoryInterface {
    id: number;
    name: string;
    age: number;
    update(user: UserEntity): Resp;
}

interface UserEntity {
    id: number;
    name: string;
    age: number;
}

interface UserUseCaseInterface {
    UpdateUser(user: UserEntity): Resp;
}
