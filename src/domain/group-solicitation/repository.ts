import { BaseRepository } from "@domain/base";
import { CreateGroupSolicitationDTO, UpdateGroupSolicitationDTO } from "./dto";
import { GroupSolicitation } from "./entity";

export interface GroupSolicitationRepository
    extends Pick<
        BaseRepository<GroupSolicitation, CreateGroupSolicitationDTO, UpdateGroupSolicitationDTO>,
        "findById" | "delete" | "create" | "update"
    > {}
