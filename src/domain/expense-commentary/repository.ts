import { BaseRepository } from "@domain/base";
import { CreateExpenseCommentaryDTO, UpdateExpenseCommentaryDTO } from "./dto";
import { ExpenseCommentary } from "./entity";

export interface ExpenseCommentaryRepository
    extends Pick<
        BaseRepository<ExpenseCommentary, CreateExpenseCommentaryDTO, UpdateExpenseCommentaryDTO>,
        "findById" | "delete" | "create" | "update"
    > {}
