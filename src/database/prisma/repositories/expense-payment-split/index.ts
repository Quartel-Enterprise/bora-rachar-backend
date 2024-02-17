import { prisma } from "@database/prisma/client";
import { ExpensePaymentSplitRepository } from "@domain/expense-payment-split";

const expensePaymentSplitModel = prisma.expensePaymentSplit;

export function makeExpensePaymentSplitRepository(
    model = expensePaymentSplitModel,
): ExpensePaymentSplitRepository {
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
            const expensePaymentSplitDeleted = await model.delete({ where: { id } });
            return !!expensePaymentSplitDeleted;
        },
    };
}

export const expensePaymentSplitRepository = makeExpensePaymentSplitRepository();
