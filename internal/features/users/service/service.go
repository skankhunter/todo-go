package users_service

import (
	"context"

	"github.com/skankhunter/todo-go/internal/core/domain"
)

type UsersService struct {
	usersRepository UsersRepository
}

type UsersRepository interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
}

func NewUsersRepository(usersRepository UsersRepository) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
	}
}
