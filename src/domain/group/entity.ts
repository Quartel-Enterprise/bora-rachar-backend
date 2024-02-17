import { BaseEntity } from "@domain/base";
import { User } from "@domain/user";
import { Expense } from "@domain/expense";
import { GroupParticipant } from "@domain/group-participant";
import { GroupSolicitation } from "@domain/group-solicitation";

export interface Group extends BaseEntity {
    id: string;
    name: string;
    avatar: string;
    accessCode: string;
    createdBy: string;

    user?: User;
    expenses?: Expense[];
    groupParticipants?: GroupParticipant[];
    groupSolicitations?: GroupSolicitation[];
}
