package userService

import (
	"BestPortfolio/internal/entity"
	"BestPortfolio/internal/repo"
	"crypto/sha256"
	"fmt"
	"net/http"
)

type UserUseCase struct {
	userRepo      repo.UserRepo
	activeProfile string
}

func NewUserUseCase(userRepo repo.UserRepo) UserService {
	return &UserUseCase{userRepo: userRepo}
}

func (u *UserUseCase) SetActiveProfile(activeProfile string) {
	u.activeProfile = activeProfile
}

func (u *UserUseCase) GetDummyUser(request *http.Request) *entity.User {
	hashedAddress := u.getUserAddress(request)
	user := &entity.User{Address: hashedAddress}
	return user
}

func (u *UserUseCase) GetUser(request *http.Request) *entity.User {
	hashedAddress := u.getUserAddress(request)
	user, err := u.userRepo.FindUserByAddress(hashedAddress)
	if err != nil {
		panic(err)
	}
	if user != nil {
		return user
	}
	return &entity.User{Address: hashedAddress}
}

func (u *UserUseCase) getUserAddress(request *http.Request) string {
	if u.activeProfile == "" {
		return hashSHA256(request.RemoteAddr + " " + getRequestLocalAddr(request))
	}
	switch u.activeProfile {
	case "production":
		return hashSHA256(request.Header.Get("CF-Connecting-IP"))
	case "dev", "default":
		fallthrough
	default:
		return hashSHA256(request.RemoteAddr + " " + getRequestLocalAddr(request))
	}

}

func getRequestLocalAddr(request *http.Request) string {
	return request.Context().Value(http.LocalAddrContextKey).(string)
}

func hashSHA256(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
