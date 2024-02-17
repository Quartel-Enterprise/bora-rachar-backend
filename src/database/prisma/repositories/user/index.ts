import { prisma } from "@database/prisma/client";
import { UserRepository } from "@domain/user";

const userModel = prisma.user;

export function makeUserRepository(model = userModel): UserRepository {
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
            const userDeleted = await model.delete({ where: { id } });
            return !!userDeleted;
        },
    };
}

export const userRepository = makeUserRepository();
