import { Expense } from "./entity";

export type CreateExpenseDTO = Omit<
    Expense,
    | "id"
    | "user"
    | "group"
    | "expensePaymentSplits"
    | "expenseCommentaries"
    | "createdAt"
    | "updatedAt"
    | "deletedAt"
>;

export type UpdateExpenseDTO = Partial<
    Omit<
        Expense,
        | "id"
        | "user"
        | "group"
        | "expensePaymentSplits"
        | "expenseCommentaries"
        | "createdAt"
        | "updatedAt"
        | "deletedAt"
    >
>;
