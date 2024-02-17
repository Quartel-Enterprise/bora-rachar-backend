import { BaseRepository } from "@domain/base";
import { CreateGroupParticipantDTO, UpdateGroupParticipantDTO } from "./dto";
import { GroupParticipant } from "./entity";

export interface GroupParticipantRepository
    extends Pick<
        BaseRepository<GroupParticipant, CreateGroupParticipantDTO, UpdateGroupParticipantDTO>,
        "findById" | "delete" | "create" | "update"
    > {}
