package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/takahashis-shun/go-swagger-tutorial/models"
	"github.com/takahashis-shun/go-swagger-tutorial/restapi/operations"
)

var (
	Users map[uint64]models.User
)

func init() {
	Users = map[uint64]models.User{}
}

type PostUserHandler struct{}

func (h *PostUserHandler) Handle(params operations.PostUserParams) middleware.Responder {
	u := models.User{
		ID:   params.Body.ID,
		Name: params.Body.Name,
	}

	Users[params.Body.ID] = u

	return operations.NewPostUserCreated()
}
