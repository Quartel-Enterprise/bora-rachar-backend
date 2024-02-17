import { ExpenseCommentary } from "./entity";

export type CreateExpenseCommentaryDTO = Omit<
    ExpenseCommentary,
    "id" | "createdAt" | "updatedAt" | "deletedAt"
>;

export type UpdateExpenseCommentaryDTO = Partial<
    Omit<ExpenseCommentary, "id" | "createdAt" | "updatedAt" | "deletedAt">
>;
