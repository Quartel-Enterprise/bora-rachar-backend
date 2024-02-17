import { BaseRepository } from "@domain/base";
import { CreateExpenseDTO, UpdateExpenseDTO } from "./dto";
import { Expense } from "./entity";

export interface ExpenseRepository
    extends Pick<
        BaseRepository<Expense, CreateExpenseDTO, UpdateExpenseDTO>,
        "findById" | "delete" | "create" | "update"
    > {}
