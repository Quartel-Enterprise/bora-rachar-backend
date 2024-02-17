import { BaseEntity } from "@domain/base";
import { Group } from "@domain/group";
import { User } from "@domain/user";

export interface GroupParticipant extends BaseEntity {
    id: string;
    userId: string;
    groupId: string;
    isAdmin: boolean;
    admissionDate: Date;
    removedAt?: Date;

    group?: Group;
    user?: User;
}
