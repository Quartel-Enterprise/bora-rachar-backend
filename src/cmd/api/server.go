package main

import (
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/cmd/generated-code"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/flow"
	"net/http"
)

type BoraRacharServer struct{}

// DeleteCommentary (DELETE /commentaries/{commentaryId})
func (_ BoraRacharServer) DeleteCommentary(w http.ResponseWriter, r *http.Request, commentaryId string) {
	flow.DeleteCommentary(w, r, commentaryId, MYSLQ)
}

// UpdateCommentary (PUT /commentaries/{commentaryId})
func (_ BoraRacharServer) UpdateCommentary(w http.ResponseWriter, r *http.Request, commentaryId string) {
	flow.UpdateCommentary(w, r, commentaryId, MYSLQ)
}

// PostScreensLogin (POST /screen/login)
func (_ BoraRacharServer) PostScreensLogin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// GetScreensActivities (GET /screens/activities)
func (_ BoraRacharServer) GetScreensActivities(w http.ResponseWriter, r *http.Request, params swagger.GetScreensActivitiesParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// GetScreensContacts (GET /screens/contacts)
func (_ BoraRacharServer) GetScreensContacts(w http.ResponseWriter, r *http.Request, params swagger.GetScreensContactsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// PostScreensContacts (POST /screens/contacts)
func (_ BoraRacharServer) PostScreensContacts(w http.ResponseWriter, r *http.Request, params swagger.PostScreensContactsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// CreatExpenseScreen (POST /screens/expenses)
func (_ BoraRacharServer) CreatExpenseScreen(w http.ResponseWriter, r *http.Request, params swagger.CreatExpenseScreenParams) {
	flow.CreatExpenseScreen(w, r, params, MYSLQ)
}

// DeleteScreensExpensesExpenseId (DELETE /screens/expenses/{expenseId})
func (_ BoraRacharServer) DeleteScreensExpensesExpenseId(w http.ResponseWriter, r *http.Request, expenseId string, params swagger.DeleteScreensExpensesExpenseIdParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// PutScreensExpensesExpenseId (PUT /screens/expenses/{expenseId})
func (_ BoraRacharServer) PutScreensExpensesExpenseId(w http.ResponseWriter, r *http.Request, expenseId string, params swagger.PutScreensExpensesExpenseIdParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// AddCommentary (POST /screens/expenses/{expenseId}/commentary)
func (_ BoraRacharServer) AddCommentary(w http.ResponseWriter, r *http.Request, expenseId string) {
	flow.AddCommentary(w, r, expenseId, MYSLQ)
}

// DeleteGroup (DELETE /groups/{groupId})
func (_ BoraRacharServer) DeleteGroup(w http.ResponseWriter, r *http.Request, groupId string) {
	flow.DeleteGroup(w, r, groupId, MYSLQ)
}

// UpdateGroup (PUT /groups/{groupId})
func (_ BoraRacharServer) UpdateGroup(w http.ResponseWriter, r *http.Request, groupId string) {
	flow.UpdateGroup(w, r, groupId, MYSLQ)
}

// GetGroupListScreen (GET /screens/groups)
func (_ BoraRacharServer) GetGroupListScreen(w http.ResponseWriter, r *http.Request, params swagger.GetGroupListScreenParams) {
	flow.GetGroupListScreen(w, r, params, MYSLQ)
}

// CreatGroupScreen (POST /screens/groups)
func (_ BoraRacharServer) CreatGroupScreen(w http.ResponseWriter, r *http.Request, params swagger.CreatGroupScreenParams) {
	flow.CreatGroupScreen(w, r, params, MYSLQ)
}

// GetGroupScreen (GET /screens/groups/{groupId})
func (_ BoraRacharServer) GetGroupScreen(w http.ResponseWriter, r *http.Request, groupId string, params swagger.GetGroupScreenParams) {
	flow.GetGroupScreen(w, r, groupId, params, MYSLQ)
}

// AddParticipant (POST /groups/{groupId}/participants)
func (_ BoraRacharServer) AddParticipant(w http.ResponseWriter, r *http.Request, groupId string) {
	flow.AddParticipant(w, r, groupId, MYSLQ)
}

// RemoveParticipant (DELETE /groups/{groupId}/participants/{userId})
func (_ BoraRacharServer) RemoveParticipant(w http.ResponseWriter, r *http.Request, groupId string, userId string) {
	flow.RemoveParticipant(w, r, groupId, userId, MYSLQ)
}

// UpdateParticipant (PUT /groups/{groupId}/participants/{userId})
func (_ BoraRacharServer) UpdateParticipant(w http.ResponseWriter, r *http.Request, groupId string, userId string) {
	flow.UpdateParticipant(w, r, groupId, userId, MYSLQ)
}

// GetUsers (GET /users)
func (_ BoraRacharServer) GetUsers(w http.ResponseWriter, r *http.Request, params swagger.GetUsersParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// PostUsers (POST /users)
func (_ BoraRacharServer) PostUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// PutUsersUserId (PUT /users/{userId})
func (_ BoraRacharServer) PutUsersUserId(w http.ResponseWriter, r *http.Request, userId string) {
	w.WriteHeader(http.StatusNotImplemented)
}
