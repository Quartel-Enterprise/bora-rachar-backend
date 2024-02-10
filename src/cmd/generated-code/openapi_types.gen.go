// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package generated

import (
	"time"
)

const (
	BasicAuthScopes = "BasicAuth.Scopes"
)

// Defines values for ActivitiesScreenActivitiesType.
const (
	EXECUTEDPAYMENT      ActivitiesScreenActivitiesType = "EXECUTED_PAYMENT"
	EXPENSEPARTICIPATION ActivitiesScreenActivitiesType = "EXPENSE_PARTICIPATION"
	NEWGROUPPARTICIPANT  ActivitiesScreenActivitiesType = "NEW_GROUP_PARTICIPANT"
	RECEIVEDPAYMENT      ActivitiesScreenActivitiesType = "RECEIVED_PAYMENT"
)

// Defines values for ExpenseCategory.
const (
	ENTERTAINMENT  ExpenseCategory = "ENTERTAINMENT"
	FOOD           ExpenseCategory = "FOOD"
	TRANSPORTATION ExpenseCategory = "TRANSPORTATION"
)

// ActivitiesScreenActivities defines model for ActivitiesScreenActivities.
type ActivitiesScreenActivities struct {
	Date       *time.Time                      `json:"date,omitempty"`
	GroupId    *string                         `json:"groupId,omitempty"`
	GroupName  *string                         `json:"groupName,omitempty"`
	GroupPhoto *string                         `json:"groupPhoto,omitempty"`
	Type       *ActivitiesScreenActivitiesType `json:"type,omitempty"`
	UserId     *string                         `json:"userId,omitempty"`
	UserName   *string                         `json:"userName,omitempty"`
	UserPhoto  *string                         `json:"userPhoto,omitempty"`
	Value      *float32                        `json:"value,omitempty"`
}

// ActivitiesScreenActivitiesType defines model for ActivitiesScreenActivities.Type.
type ActivitiesScreenActivitiesType string

// ActivitiesScreenResponse defines model for ActivitiesScreenResponse.
type ActivitiesScreenResponse struct {
	Activities *[]ActivitiesScreenActivities `json:"activities,omitempty"`
	Balance    *SummaryBalance               `json:"balance,omitempty"`
	Meta       *Pagination                   `json:"meta,omitempty"`
}

// ContactsScreenRequestBody defines model for ContactsScreenRequestBody.
type ContactsScreenRequestBody struct {
	Contacts *[]struct {
		Email    *string `json:"email,omitempty"`
		Telefone *string `json:"telefone,omitempty"`
	} `json:"contacts,omitempty"`
}

// ContactsScreenResponse defines model for ContactsScreenResponse.
type ContactsScreenResponse struct {
	Contacts *[]struct {
		Email    *string  `json:"email,omitempty"`
		Id       *string  `json:"id,omitempty"`
		Name     *string  `json:"name,omitempty"`
		Telefone *string  `json:"telefone,omitempty"`
		Value    *float32 `json:"value,omitempty"`
	} `json:"contacts,omitempty"`
}

// ExpenseCategory defines model for ExpenseCategory.
type ExpenseCategory string

// ExpenseCommentRequestBody defines model for ExpenseCommentRequestBody.
type ExpenseCommentRequestBody struct {
	Commentary string     `json:"commentary"`
	Date       *time.Time `json:"date,omitempty"`
	UserId     string     `json:"userId"`
}

// ExpenseScreenRequestBody defines model for ExpenseScreenRequestBody.
type ExpenseScreenRequestBody struct {
	Category     *ExpenseCategory `json:"category,omitempty"`
	Description  *string          `json:"description,omitempty"`
	Name         string           `json:"name"`
	Participants []struct {
		UserId *string `json:"userId,omitempty"`
		Value  *string `json:"value,omitempty"`
	} `json:"participants"`
	Payers []struct {
		UserId *string `json:"userId,omitempty"`
		Value  *string `json:"value,omitempty"`
	} `json:"payers"`
	Value float32 `json:"value"`
}

// GroupDetailsScreenBalance defines model for GroupDetailsScreenBalance.
type GroupDetailsScreenBalance struct {
	AmountToPay *struct {
		PayTo *[]struct {
			UserId    *string  `json:"userId,omitempty"`
			UserName  *string  `json:"userName,omitempty"`
			UserPhoto *string  `json:"userPhoto,omitempty"`
			Value     *float32 `json:"value,omitempty"`
		} `json:"payTo,omitempty"`
		Value *float32 `json:"value,omitempty"`
	} `json:"amountToPay,omitempty"`
	AmountToReceive *struct {
		ReceiveFrom *[]struct {
			UserId    *string  `json:"userId,omitempty"`
			UserName  *string  `json:"userName,omitempty"`
			UserPhoto *string  `json:"userPhoto,omitempty"`
			Value     *float32 `json:"value,omitempty"`
		} `json:"receiveFrom,omitempty"`
		Value *float32 `json:"value,omitempty"`
	} `json:"amountToReceive,omitempty"`
}

// GroupDetailsScreenExpense defines model for GroupDetailsScreenExpense.
type GroupDetailsScreenExpense struct {
	Category   *ExpenseCategory `json:"category,omitempty"`
	Commentary *[]struct {
		Commentary *string    `json:"commentary,omitempty"`
		Date       *time.Time `json:"date,omitempty"`
		UserId     *string    `json:"userId,omitempty"`
		UserName   *string    `json:"userName,omitempty"`
		UserPhoto  *string    `json:"userPhoto,omitempty"`
	} `json:"commentary,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	CreatedBy *struct {
		UserId    *string `json:"userId,omitempty"`
		UserName  *string `json:"userName,omitempty"`
		UserPhoto *string `json:"userPhoto,omitempty"`
	} `json:"createdBy,omitempty"`
	Description *string `json:"description,omitempty"`
	Id          *string `json:"id,omitempty"`
	Involved    *[]struct {
		Balance      *float32 `json:"balance,omitempty"`
		IsDebtSetted *bool    `json:"isDebtSetted,omitempty"`
		UserId       *string  `json:"userId,omitempty"`
		UserName     *string  `json:"userName,omitempty"`
		UserPhoto    *string  `json:"userPhoto,omitempty"`
	} `json:"involved,omitempty"`
	Photo *string `json:"photo,omitempty"`
}

// GroupDetailsScreenResponse defines model for GroupDetailsScreenResponse.
type GroupDetailsScreenResponse struct {
	Expenses  *[]GroupDetailsScreenExpense `json:"expenses,omitempty"`
	GroupInfo *struct {
		Id           *string `json:"id,omitempty"`
		Name         *string `json:"name,omitempty"`
		Participants *[]struct {
			Balance   *GroupDetailsScreenBalance `json:"balance,omitempty"`
			UserId    *string                    `json:"userId,omitempty"`
			UserName  *string                    `json:"userName,omitempty"`
			UserPhoto *string                    `json:"userPhoto,omitempty"`
		} `json:"participants,omitempty"`
		Photo *string `json:"photo,omitempty"`
	} `json:"groupInfo,omitempty"`
	Meta *Pagination `json:"meta,omitempty"`
}

// GroupsScreenGroup defines model for GroupsScreenGroup.
type GroupsScreenGroup struct {
	AmountToPay     *float32                   `json:"amountToPay,omitempty"`
	AmountToReceive *float32                   `json:"amountToReceive,omitempty"`
	GroupId         *string                    `json:"groupId,omitempty"`
	GroupName       *string                    `json:"groupName,omitempty"`
	GroupPhoto      *string                    `json:"groupPhoto,omitempty"`
	Participants    *[]GroupsScreenParticipant `json:"participants,omitempty"`
}

// GroupsScreenParticipant defines model for GroupsScreenParticipant.
type GroupsScreenParticipant struct {
	UserId    *string `json:"userId,omitempty"`
	UserName  *string `json:"userName,omitempty"`
	UserPhoto *string `json:"userPhoto,omitempty"`
}

// GroupsScreenParticipantRequestBody defines model for GroupsScreenParticipantRequestBody.
type GroupsScreenParticipantRequestBody struct {
	IsAdmin *bool   `json:"isAdmin,omitempty"`
	UserId  *string `json:"userId,omitempty"`
}

// GroupsScreenRequestBody defines model for GroupsScreenRequestBody.
type GroupsScreenRequestBody struct {
	Name         string                               `json:"name"`
	Participants []GroupsScreenParticipantRequestBody `json:"participants"`
	Photo        *string                              `json:"photo,omitempty"`
}

// GroupsScreenResponse defines model for GroupsScreenResponse.
type GroupsScreenResponse struct {
	Balance *SummaryBalance      `json:"balance,omitempty"`
	Groups  *[]GroupsScreenGroup `json:"groups,omitempty"`
}

// InfoLoginResponse defines model for InfoLoginResponse.
type InfoLoginResponse struct {
	AccessToken  *string `json:"access_token,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty"`
	UserId       *string `json:"userId,omitempty"`
}

// LoginRequestBody defines model for LoginRequestBody.
type LoginRequestBody struct {
	// Code access token that grants access to that API
	Code *string `json:"code,omitempty"`
}

// Pagination defines model for Pagination.
type Pagination struct {
	CurrentPage  *float32 `json:"currentPage,omitempty"`
	TotalPages   *float32 `json:"totalPages,omitempty"`
	TotalRecords *float32 `json:"totalRecords,omitempty"`
}

// SummaryBalance defines model for SummaryBalance.
type SummaryBalance struct {
	AmountToPay     *float32 `json:"amountToPay,omitempty"`
	AmountToReceive *float32 `json:"amountToReceive,omitempty"`
}

// UserRequestBody defines model for UserRequestBody.
type UserRequestBody struct {
	Avatar *string `json:"avatar,omitempty"`
	Email  *string `json:"email,omitempty"`
	Name   *string `json:"name,omitempty"`
	PixKey *string `json:"pixKey,omitempty"`
}

// UserResponse defines model for UserResponse.
type UserResponse struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Email     *string    `json:"email,omitempty"`
	Id        *int       `json:"id,omitempty"`
	Name      *string    `json:"name,omitempty"`
	PixKey    *string    `json:"pixKey,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// Limit defines model for limit.
type Limit = int

// Page defines model for page.
type Page = int

// UserIdHeader defines model for user-id-header.
type UserIdHeader = string

// GetScreensActivitiesParams defines parameters for GetScreensActivities.
type GetScreensActivitiesParams struct {
	Page   Page         `form:"page" json:"page"`
	Limit  Limit        `form:"limit" json:"limit"`
	UserId UserIdHeader `json:"userId"`
}

// GetScreensContactsParams defines parameters for GetScreensContacts.
type GetScreensContactsParams struct {
	UserId       UserIdHeader `json:"userId"`
	AccessCode   string       `json:"access-code"`
	RefreshToken string       `json:"refresh-token"`
}

// PostScreensContactsParams defines parameters for PostScreensContacts.
type PostScreensContactsParams struct {
	UserId       UserIdHeader `json:"userId"`
	AccessCode   string       `json:"access-code"`
	RefreshToken string       `json:"refresh-token"`
}

// PostScreensExpensesParams defines parameters for PostScreensExpenses.
type PostScreensExpensesParams struct {
	UserId UserIdHeader `json:"userId"`
}

// DeleteScreensExpensesExpenseIdParams defines parameters for DeleteScreensExpensesExpenseId.
type DeleteScreensExpensesExpenseIdParams struct {
	UserId UserIdHeader `json:"userId"`
}

// PutScreensExpensesExpenseIdParams defines parameters for PutScreensExpensesExpenseId.
type PutScreensExpensesExpenseIdParams struct {
	UserId UserIdHeader `json:"userId"`
}

// GetScreensGroupsParams defines parameters for GetScreensGroups.
type GetScreensGroupsParams struct {
	UserId UserIdHeader `json:"userId"`
}

// PostScreensGroupsParams defines parameters for PostScreensGroups.
type PostScreensGroupsParams struct {
	UserId UserIdHeader `json:"userId"`
}

// GetScreensGroupsGroupIdParams defines parameters for GetScreensGroupsGroupId.
type GetScreensGroupsGroupIdParams struct {
	Page   Page         `form:"page" json:"page"`
	Limit  Limit        `form:"limit" json:"limit"`
	UserId UserIdHeader `json:"userId"`
}

// GetUsersParams defines parameters for GetUsers.
type GetUsersParams struct {
	UserId *string `form:"userId,omitempty" json:"userId,omitempty"`
	Email  *string `form:"email,omitempty" json:"email,omitempty"`
}

// PostScreensContactsJSONRequestBody defines body for PostScreensContacts for application/json ContentType.
type PostScreensContactsJSONRequestBody = ContactsScreenRequestBody

// PostScreensExpensesJSONRequestBody defines body for PostScreensExpenses for application/json ContentType.
type PostScreensExpensesJSONRequestBody = ExpenseScreenRequestBody

// PutScreensExpensesExpenseIdJSONRequestBody defines body for PutScreensExpensesExpenseId for application/json ContentType.
type PutScreensExpensesExpenseIdJSONRequestBody = ExpenseScreenRequestBody

// PostScreensGroupsJSONRequestBody defines body for PostScreensGroups for application/json ContentType.
type PostScreensGroupsJSONRequestBody = GroupsScreenRequestBody

// PostScreensGroupsExpenseIdCommentaryJSONRequestBody defines body for PostScreensGroupsExpenseIdCommentary for application/json ContentType.
type PostScreensGroupsExpenseIdCommentaryJSONRequestBody = ExpenseCommentRequestBody

// PostScreensLoginJSONRequestBody defines body for PostScreensLogin for application/json ContentType.
type PostScreensLoginJSONRequestBody = LoginRequestBody

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody = UserRequestBody

// PutUsersUserIdJSONRequestBody defines body for PutUsersUserId for application/json ContentType.
type PutUsersUserIdJSONRequestBody = UserRequestBody
