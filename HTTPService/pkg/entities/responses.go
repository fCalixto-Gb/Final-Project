package entities

type GetUserResp struct {
	Id      string
	Name    string
	Age     int64
	Country string
	Job     string
	Created string
	Email   string
}

type CreateUserResp struct {
	Id      string
	Created string
}

type DeleteUserResp struct {
	Deleted string
}

type UpdateUserResp struct {
	Updated string
}

type AuthenticateResp struct {
	Token string
}
