import { GroupParticipant } from "./entity";

export type CreateGroupParticipantDTO = Omit<
    GroupParticipant,
    "id" | "group" | "user" | "createdAt" | "updatedAt" | "deletedAt"
>;

export type UpdateGroupParticipantDTO = Partial<
    Omit<GroupParticipant, "id" | "group" | "user" | "createdAt" | "updatedAt" | "deletedAt">
>;
