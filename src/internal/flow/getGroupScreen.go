package flow

import (
	"context"
	"fmt"
	swagger "github.com/Quartel-Enterprise/bora-rachar-backend/src/cmd/generated-code"
	repository_model "github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/repository/model"
	repository_query "github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/repository/query"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/util"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func GetGroupScreen(w http.ResponseWriter, r *http.Request, groupId string, params swagger.GetGroupScreenParams, db *sqlx.DB) {
	groupInfo, err := repository_query.GetGroup(context.Background(), groupId, db)
	if err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err)
		return
	}

	participants, err := repository_query.GetParticipants(context.Background(), groupId, db)
	if err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err)
		return
	}

	expensesSplitToPay, err := repository_query.GetExpensePaymentSplitNoSettle(context.Background(), db, params.UserId)
	if err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Printf("%v, %v, %v", groupInfo, participants, expensesSplitToPay)

	util.HttpResponseWihBody(w, http.StatusOK, nil, nil)
}

func mapResponse(group repository_model.Group, groupParticipants []repository_model.JoinUserAndGroupParticipant) swagger.GroupDetailsScreenResponse {

	var groupId = group.Id
	var groupName = group.Name
	var groupPhoto = group.Avatar
	var participants = mapToParticipantsResponse(groupParticipants)

	var groupInfo = swagger.GroupDetailsScreenGroupInfo{
		Expenses:     nil,
		Id:           &groupId,
		Name:         &groupName,
		Participants: &participants,
		Photo:        &groupPhoto,
	}

	var meta = swagger.Pagination{
		CurrentPage:  nil,
		TotalPages:   nil,
		TotalRecords: nil,
	}

	return swagger.GroupDetailsScreenResponse{
		GroupInfo: &groupInfo,
		Meta:      &meta,
	}
}

func mapToParticipantsResponse(participant []repository_model.JoinUserAndGroupParticipant) []swagger.GroupDetailsScreenParticipant {
	var response = make([]swagger.GroupDetailsScreenParticipant, 0)

	for _, p := range participant {
		var item = swagger.GroupDetailsScreenParticipant{
			UserId:    p.UserId,
			UserName:  p.Name,
			UserPhoto: p.Avatar,
		}

		response = append(response, item)
	}

	return response
}
