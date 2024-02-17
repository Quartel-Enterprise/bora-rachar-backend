import { BaseEntity } from "@domain/base";
import { Group } from "@domain/group";
import { Expense } from "@domain/expense";
import { GroupParticipant } from "@domain/group-participant";
import { GroupSolicitation } from "@domain/group-solicitation";
import { ExpensePaymentSplit } from "@domain/expense-payment-split";
import { ExpenseCommentary } from "@domain/expense-commentary";

export interface User extends BaseEntity {
    id: string;
    userId: string;
    name: string;
    avatar: string;
    email: string;
    pixKey: string;

    groups?: Group[];
    expenses?: Expense[];
    groupParticipants?: GroupParticipant[];
    groupSolicitations?: GroupSolicitation[];
    expensePaymentSplits?: ExpensePaymentSplit[];
    expenseCommentaries?: ExpenseCommentary[];
}
