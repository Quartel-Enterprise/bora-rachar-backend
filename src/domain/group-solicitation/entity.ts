import { BaseEntity } from "@domain/base";
import { Group } from "@domain/group";
import { User } from "@domain/user";

export interface GroupSolicitation extends BaseEntity {
    id: string;
    userId: string;
    groupId: string;

    group?: Group;
    user?: User;
}
