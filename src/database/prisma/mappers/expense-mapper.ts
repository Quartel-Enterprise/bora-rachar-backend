import { Expense } from "@domain/expense";
import { Expense as PrismaExpense } from "@prisma/client";

function toDomain(prismaExpense: PrismaExpense | undefined): Expense | undefined {
    if (!prismaExpense) {
        return undefined;
    }

    return {
        ...prismaExpense,
        value: prismaExpense.value.toNumber(),
    };
}

export const expenseMapper = {
    toDomain,
};
