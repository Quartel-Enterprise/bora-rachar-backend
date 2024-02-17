import { ExpenseCommentary } from "./entity";

export type CreateExpenseCommentaryDTO = Omit<
    ExpenseCommentary,
    "id" | "expense" | "user" | "createdAt" | "updatedAt" | "deletedAt"
>;

export type UpdateExpenseCommentaryDTO = Partial<
    Omit<ExpenseCommentary, "id" | "expense" | "user" | "createdAt" | "updatedAt" | "deletedAt">
>;
