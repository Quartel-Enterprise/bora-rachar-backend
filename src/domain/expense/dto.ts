import { Expense } from "./entity";

export type CreateExpenseDTO = Omit<Expense, "id" | "createdAt" | "updatedAt" | "deletedAt">;

export type UpdateExpenseDTO = Partial<
    Omit<Expense, "id" | "createdAt" | "updatedAt" | "deletedAt">
>;
