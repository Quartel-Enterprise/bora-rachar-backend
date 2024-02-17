import { BaseRepository } from "@domain/base";
import { CreateExpensePaymentSplitDTO, UpdateExpensePaymentSplitDTO } from "./dto";
import { ExpensePaymentSplit } from "./entity";

export interface ExpensePaymentSplitRepository
    extends Pick<
        BaseRepository<
            ExpensePaymentSplit,
            CreateExpensePaymentSplitDTO,
            UpdateExpensePaymentSplitDTO
        >,
        "findById" | "delete" | "create" | "update"
    > {}
