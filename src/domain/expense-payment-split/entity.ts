import { BaseEntity } from "@domain/base";
import { User } from "@domain/user";
import { Expense } from "@domain/expense";

export enum TransactionTypeEnum {
    P = "P",
    B = "B",
}

export type TransactionType = `${TransactionTypeEnum}`;

export interface ExpensePaymentSplit extends BaseEntity {
    id: string;
    userId: string;
    expenseId: string;
    value: number;
    transactionType: TransactionType;
    isDebtSettled: boolean;

    user?: User;
    expense?: Expense;
}
