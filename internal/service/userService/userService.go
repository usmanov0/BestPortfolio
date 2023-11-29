package userService

import (
	"BestPortfolio/internal/entity"
	"net/http"
)

type UserService interface {
	SetActiveProfile(activeProfile string)
	GetDummyUser(request *http.Request) *entity.User
	GetUser(request *http.Request) *entity.User
	getUserAddress(request *http.Request) string
}
