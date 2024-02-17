export interface BaseRepository<
    EntityType,
    CreateDTOType = Partial<EntityType>,
    UpdateDTOType = Partial<EntityType>,
> {
    findAll: () => Promise<EntityType[]>;
    findById: (id: string) => Promise<EntityType>;
    create: (createDTO: CreateDTOType) => Promise<EntityType>;
    update: (id: string, updateDTO: UpdateDTOType) => Promise<EntityType>;
    delete: (id: string) => Promise<boolean>;
}
