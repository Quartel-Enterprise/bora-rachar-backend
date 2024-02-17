import { ExpensePaymentSplit } from "./entity";

export type CreateExpensePaymentSplitDTO = Omit<
    ExpensePaymentSplit,
    "id" | "expense" | "user" | "createdAt" | "updatedAt" | "deletedAt"
>;

export type UpdateExpensePaymentSplitDTO = Partial<
    Omit<ExpensePaymentSplit, "id" | "expense" | "user" | "createdAt" | "updatedAt" | "deletedAt">
>;
