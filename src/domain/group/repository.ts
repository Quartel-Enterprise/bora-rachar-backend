import { BaseRepository } from "@domain/base";
import { CreateGroupDTO, UpdateGroupDTO } from "./dto";
import { Group } from "./entity";

export interface GroupRepository
    extends Pick<
        BaseRepository<Group, CreateGroupDTO, UpdateGroupDTO>,
        "findById" | "delete" | "create" | "update"
    > {}
