import { prisma } from "@database/prisma/client";
import { GroupParticipantRepository } from "@domain/group-participant";

const groupParticipantModel = prisma.groupParticipant;

export function makeGroupParticipantRepository(
    model = groupParticipantModel,
): GroupParticipantRepository {
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
            const groupParticipantDeleted = await model.delete({ where: { id } });
            return !!groupParticipantDeleted;
        },
    };
}

export const groupParticipantRepository = makeGroupParticipantRepository();
