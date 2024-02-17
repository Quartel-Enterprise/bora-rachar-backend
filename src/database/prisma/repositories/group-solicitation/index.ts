import { prisma } from "@database/prisma/client";
import { GroupSolicitationRepository } from "@domain/group-solicitation";

const groupSolicitationModel = prisma.groupSolicitation;

export function makeGroupSolicitationRepository(
    model = groupSolicitationModel,
): GroupSolicitationRepository {
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
            const groupSolicitationDeleted = await model.delete({ where: { id } });
            return !!groupSolicitationDeleted;
        },
    };
}

export const groupSolicitationRepository = makeGroupSolicitationRepository();
