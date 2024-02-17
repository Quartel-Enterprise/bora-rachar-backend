import {
    makeUserRepository,
    makeGroupRepository,
    makeExpenseRepository,
    makeGroupParticipantRepository,
    makeGroupSolicitationRepository,
    makeExpensePaymentSplitRepository,
    makeExpenseCommentaryRepository,
} from "../repositories";

export const RepositoriesMap = {
    user: makeUserRepository,
    group: makeGroupRepository,
    expense: makeExpenseRepository,
    groupParticipant: makeGroupParticipantRepository,
    groupSolicitation: makeGroupSolicitationRepository,
    expensePaymentSplit: makeExpensePaymentSplitRepository,
    expenseCommentary: makeExpenseCommentaryRepository,
} as const;
