import { GroupParticipant } from "./entity";

export type CreateGroupParticipantDTO = Omit<
    GroupParticipant,
    "id" | "createdAt" | "updatedAt" | "deletedAt"
>;

export type UpdateGroupParticipantDTO = Partial<
    Omit<GroupParticipant, "id" | "createdAt" | "updatedAt" | "deletedAt">
>;
