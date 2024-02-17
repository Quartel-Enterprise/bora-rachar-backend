import { GroupSolicitation } from "./entity";

export type CreateGroupSolicitationDTO = Omit<
    GroupSolicitation,
    "id" | "createdAt" | "updatedAt" | "deletedAt"
>;

export type UpdateGroupSolicitationDTO = Partial<
    Omit<GroupSolicitation, "id" | "createdAt" | "updatedAt" | "deletedAt">
>;
