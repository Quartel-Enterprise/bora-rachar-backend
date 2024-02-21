package repository_model

type JoinUserAndGroupParticipant struct {
	UserId *string `db:"user_id"`
	Avatar *string
	Name   *string
}
