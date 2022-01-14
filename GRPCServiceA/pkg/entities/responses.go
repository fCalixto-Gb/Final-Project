package entities

type GetUserResp struct {
	Id          string
	Name        string
	Age         int64
	Nationality string
	Job         string
	Created     string
	Email       string
	Error       Status
}

type CreateUserResp struct {
	Id      string
	Created string
	Error   Status
}

type Status struct {
	Message string
	Code    int32
}
