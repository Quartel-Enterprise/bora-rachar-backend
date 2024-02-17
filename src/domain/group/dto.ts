import { Group } from "./entity";

export type CreateGroupDTO = Omit<
    Group,
    | "id"
    | "user"
    | "expenses"
    | "groupParticipants"
    | "groupSolicitations"
    | "createdAt"
    | "updatedAt"
    | "deletedAt"
>;

export type UpdateGroupDTO = Partial<
    Omit<
        Group,
        | "id"
        | "user"
        | "expenses"
        | "groupParticipants"
        | "groupSolicitations"
        | "createdAt"
        | "updatedAt"
        | "deletedAt"
    >
>;
