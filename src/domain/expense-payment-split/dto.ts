import { ExpensePaymentSplit } from "./entity";

export type CreateExpensePaymentSplitDTO = Omit<
    ExpensePaymentSplit,
    "id" | "createdAt" | "updatedAt" | "deletedAt"
>;

export type UpdateExpensePaymentSplitDTO = Partial<
    Omit<ExpensePaymentSplit, "id" | "createdAt" | "updatedAt" | "deletedAt">
>;
