import { GroupSolicitation } from "./entity";

export type CreateGroupSolicitationDTO = Omit<
    GroupSolicitation,
    "id" | "group" | "user" | "createdAt" | "updatedAt" | "deletedAt"
>;

export type UpdateGroupSolicitationDTO = Partial<
    Omit<GroupSolicitation, "id" | "group" | "user" | "createdAt" | "updatedAt" | "deletedAt">
>;
