package flow

import (
	"context"
	swagger "github.com/Quartel-Enterprise/bora-rachar-backend/src/cmd/generated-code"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/logger"
	repository_model "github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/repository/model"
	repository_query "github.com/Quartel-Enterprise/bora-rachar-backend/src/internal/infra/repository/query"
	"github.com/Quartel-Enterprise/bora-rachar-backend/src/util"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func GetGroupListScreen(w http.ResponseWriter, r *http.Request, params swagger.GetGroupListScreenParams, db *sqlx.DB) {
	var response swagger.GroupsScreenResponse
	var groupsResponse = make([]swagger.GroupsScreenGroup, 0)
	var userId = params.UserId

	groups, err := repository_query.GetGroups(userId, db)

	if err != nil {
		logger.Error.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, g := range groups {
		participants, err := repository_query.GetParticipants(context.Background(), g.Id, db)
		if err != nil {
			util.HttpResponse(w, http.StatusInternalServerError, err)
			return
		}
		groupsResponse = append(groupsResponse, mergeGroupWithParticipants(g, participants))
	}

	response.Balance, err = getBalance(db, userId)
	if err != nil {
		util.HttpResponse(w, http.StatusInternalServerError, err)
		return
	}

	response.Groups = &groupsResponse

	util.HttpResponseWihBody(w, http.StatusOK, response, err)
}

func getBalance(db *sqlx.DB, userId string) (*swagger.SummaryBalance, error) {
	ammountToPay, err := repository_query.GetTotalAmountToPay(context.Background(), db, userId)

	if err != nil {
		return nil, err
	}

	ammountToReceive, err := repository_query.GetTotalAmountToReceive(context.Background(), db, userId)

	if err != nil {
		return nil, err
	}

	var balance = &swagger.SummaryBalance{
		AmountToPay:     &ammountToPay,
		AmountToReceive: &ammountToReceive,
	}

	return balance, nil
}

func mergeGroupWithParticipants(group repository_model.Group, participants []repository_model.JoinUserAndGroupParticipant) swagger.GroupsScreenGroup {
	var participantsDTO = make([]swagger.GroupsScreenParticipant, 0)

	for _, p := range participants {
		participantsDTO = append(participantsDTO, swagger.GroupsScreenParticipant{
			UserId:    p.UserId,
			UserName:  p.Name,
			UserPhoto: p.Avatar,
		})
	}

	var groupId = group.Id

	return swagger.GroupsScreenGroup{
		GroupId:      &groupId,
		GroupName:    &group.Name,
		GroupPhoto:   &group.Avatar,
		Participants: &participantsDTO,
	}
}
