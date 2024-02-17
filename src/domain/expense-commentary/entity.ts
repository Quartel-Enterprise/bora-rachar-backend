import { BaseEntity } from "@domain/base";
import { User } from "@domain/user";
import { Expense } from "@domain/expense";

export interface ExpenseCommentary extends BaseEntity {
    id: string;
    userId: string;
    expenseId: string;
    commentary: string;

    user?: User;
    expense?: Expense;
}
