import { prisma } from "@database/prisma/client";
import { GroupRepository } from "@domain/group";

const groupModel = prisma.group;

export function makeGroupRepository(model = groupModel): GroupRepository {
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
            const groupDeleted = await model.delete({ where: { id } });
            return !!groupDeleted;
        },
    };
}

export const groupRepository = makeGroupRepository();
