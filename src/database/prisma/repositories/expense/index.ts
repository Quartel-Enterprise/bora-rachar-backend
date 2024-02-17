import { prisma } from "@database/prisma/client";
import { ExpenseRepository } from "@domain/expense";

const expenseModel = prisma.expense;

export function makeExpenseRepository(model = expenseModel): ExpenseRepository {
    return {
        async create(data) {
            return model.create({ data });
        },
        async findById(id) {
            return model.findUniqueOrThrow({
                where: {
                    id,
                },
            });
        },
        async update(id, data) {
            return model.update({ where: { id }, data });
        },
        async delete(id) {
            const expenseDeleted = await model.delete({ where: { id } });
            return !!expenseDeleted;
        },
    };
}

export const expenseRepository = makeExpenseRepository();
