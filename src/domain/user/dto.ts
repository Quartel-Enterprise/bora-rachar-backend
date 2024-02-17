import { User } from "./entity";

export type CreateUserDTO = Omit<
    User,
    | "id"
    | "groups"
    | "expenses"
    | "groupParticipants"
    | "groupSolicitations"
    | "expensePaymentSplits"
    | "expenseCommentaries"
    | "createdAt"
    | "updatedAt"
    | "deletedAt"
>;

export type UpdateUserDTO = Partial<
    Omit<
        User,
        | "id"
        | "groups"
        | "expenses"
        | "groupParticipants"
        | "groupSolicitations"
        | "expensePaymentSplits"
        | "expenseCommentaries"
        | "createdAt"
        | "updatedAt"
        | "deletedAt"
    >
>;
