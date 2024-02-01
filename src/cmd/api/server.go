package main

import (
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/cmd/generated-code"
	"net/http"
)

type BoraRacharServer struct{}

// PostScreensLogin (POST /screen/login)
func (_ BoraRacharServer) PostScreensLogin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// GetScreensActivities (GET /screens/activities)
func (_ BoraRacharServer) GetScreensActivities(w http.ResponseWriter, r *http.Request, params generated.GetScreensActivitiesParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// GetScreensContacts (GET /screens/contacts)
func (_ BoraRacharServer) GetScreensContacts(w http.ResponseWriter, r *http.Request, params generated.GetScreensContactsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// PostScreensContacts (POST /screens/contacts)
func (_ BoraRacharServer) PostScreensContacts(w http.ResponseWriter, r *http.Request, params generated.PostScreensContactsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// PostScreensExpenses (POST /screens/expenses)
func (_ BoraRacharServer) PostScreensExpenses(w http.ResponseWriter, r *http.Request, params generated.PostScreensExpensesParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// DeleteScreensExpensesExpenseId (DELETE /screens/expenses/{expenseId})
func (_ BoraRacharServer) DeleteScreensExpensesExpenseId(w http.ResponseWriter, r *http.Request, expenseId string, params generated.DeleteScreensExpensesExpenseIdParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// PutScreensExpensesExpenseId (PUT /screens/expenses/{expenseId})
func (_ BoraRacharServer) PutScreensExpensesExpenseId(w http.ResponseWriter, r *http.Request, expenseId string, params generated.PutScreensExpensesExpenseIdParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// GetScreensGroups (GET /screens/groups)
func (_ BoraRacharServer) GetScreensGroups(w http.ResponseWriter, r *http.Request, params generated.GetScreensGroupsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// PostScreensGroups (POST /screens/groups)
func (_ BoraRacharServer) PostScreensGroups(w http.ResponseWriter, r *http.Request, params generated.PostScreensGroupsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// GetScreensGroupsGroupId (GET /screens/groups/{groupId})
func (_ BoraRacharServer) GetScreensGroupsGroupId(w http.ResponseWriter, r *http.Request, groupId string, params generated.GetScreensGroupsGroupIdParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// PostScreensGroupsGroupIdCommentary (POST /screens/groups/{groupId}/commentary)
func (_ BoraRacharServer) PostScreensGroupsGroupIdCommentary(w http.ResponseWriter, r *http.Request, groupId string) {
	w.WriteHeader(http.StatusNotImplemented)
}
