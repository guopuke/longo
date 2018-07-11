package user

type CreateResponse struct {
	Username string `json:"username"`
}

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
