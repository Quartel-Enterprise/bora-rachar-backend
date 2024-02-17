import { prisma } from "@database/prisma/client";
import { ExpenseCommentaryRepository } from "@domain/expense-commentary";

const expenseCommentaryModel = prisma.expenseCommentary;

export function makeExpenseCommentaryRepository(
    model = expenseCommentaryModel,
): ExpenseCommentaryRepository {
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
            const expenseCommentaryDeleted = await model.delete({ where: { id } });
            return !!expenseCommentaryDeleted;
        },
    };
}

export const expenseCommentaryRepository = makeExpenseCommentaryRepository();
