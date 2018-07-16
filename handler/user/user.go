package user

import "github.com/guopuke/longo/model"

type CreateResponse struct {
	Username string `json:"username"`
}

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}
