import { Group } from "./entity";

export type CreateGroupDTO = Omit<Group, "id" | "createdAt" | "updatedAt" | "deletedAt">;

export type UpdateGroupDTO = Partial<Omit<Group, "id" | "createdAt" | "updatedAt" | "deletedAt">>;
