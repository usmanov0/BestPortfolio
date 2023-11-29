package repo

import "BestPortfolio/internal/entity"

type UserRepo interface {
	Create(user *entity.User) error
	FindUserByAddress(address string) (*entity.User, error)
}
