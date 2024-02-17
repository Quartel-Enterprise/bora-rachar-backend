import { User } from "./entity";

export type CreateUserDTO = Omit<User, "id" | "createdAt" | "updatedAt" | "deletedAt">;

export type UpdateUserDTO = Partial<Omit<User, "id" | "createdAt" | "updatedAt" | "deletedAt">>;
