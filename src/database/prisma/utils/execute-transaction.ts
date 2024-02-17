import { Prisma } from "@prisma/client";
import { prisma } from "../client";

export async function executeTransaction<T>(
    callback: (tc: Prisma.TransactionClient) => Promise<T>,
    timeout = 50000,
) {
    return await prisma.$transaction(
        async (transactionClient) => {
            return await callback(transactionClient);
        },
        { timeout },
    );
}

export type ExecuteTransaction = typeof executeTransaction;
