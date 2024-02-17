import { BaseRepository } from "@domain/base";
import { CreateUserDTO, UpdateUserDTO } from "./dto";
import { User } from "./entity";

export interface UserRepository
    extends Pick<
        BaseRepository<User, CreateUserDTO, UpdateUserDTO>,
        "findById" | "delete" | "create" | "update"
    > {}
