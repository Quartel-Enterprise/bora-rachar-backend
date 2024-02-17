import { Prisma } from "@prisma/client";
import { RepositoriesMap } from "./repositories-map";

type Models = keyof Omit<
    Prisma.TransactionClient,
    "$executeRaw" | "$executeRawUnsafe" | "$queryRaw" | "$queryRawUnsafe"
>;
type Repositories = typeof RepositoriesMap;
type RepositoriesKeys = keyof Repositories;
// eslint-disable-next-line @typescript-eslint/no-explicit-any
type MapToTransactionModel<T extends any[]> = {
    [key in T[number] as `${key}Repository`]: ReturnType<Repositories[key]>;
};

export function getRepositories<T extends Models[]>(
    transactionClient: Prisma.TransactionClient,
    ...models: T
): MapToTransactionModel<RepositoriesKeys[]> {
    const transactionRepositories = {} as MapToTransactionModel<RepositoriesKeys[]>;

    for (const model of models) {
        const modelObj = transactionClient[model];
        const repositoryKey = `${String(model)}Repository`;
        const repositoryConstructor = RepositoriesMap[model];

        if (!repositoryConstructor) {
            console.error(`getRepositories: Unmapped model "${String(model)}" is being used`);
            continue;
        }

        Object.assign(transactionRepositories, {
            [repositoryKey]: repositoryConstructor(modelObj),
        });
    }

    return transactionRepositories;
}

export type GetRepositories = typeof getRepositories;
