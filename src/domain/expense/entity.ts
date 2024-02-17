import { BaseEntity } from "@domain/base";
import { User } from "@domain/user";
import { Group } from "@domain/group";
import { ExpensePaymentSplit } from "@domain/expense-payment-split";
import { ExpenseCommentary } from "@domain/expense-commentary";

export interface Expense extends BaseEntity {
    id: string;
    groupId: string;
    title: string;
    description: string;
    category: string;
    avatar: string;
    value: number;
    expenseDate: Date;
    createdBy: string;

    user?: User;
    group?: Group;
    expensePaymentSplits?: ExpensePaymentSplit[];
    expenseCommentaries?: ExpenseCommentary[];
}
