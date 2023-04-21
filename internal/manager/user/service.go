package user

import (
	"context"

	userStorage "github.com/Akezhan1/lecvisitor/internal/storage/user"
)

type service struct {
	userRepo userStorage.Storage
}

func NewUserService(userRepo userStorage.Storage) *service {
	return &service{userRepo: userRepo}
}

func (s *service) CreateUser(ctx context.Context, user User) (int, error) {
	return s.userRepo.CreateUser(ctx, user.toStorageModel())
}

func (s *service) GetUser(id int) (*userStorage.User, error) {
	strUser, err := s.userRepo.GetUser(id)
	if err != nil {
		return nil, err
	}

	u := &userStorage.User{
		ID:          strUser.ID,
		Name:        strUser.Name,
		DiscordName: strUser.DiscordName,
		TelegramID:  strUser.TelegramID,
		Role:        strUser.Role,
	}

	return u, nil
}

func (s *service) GetUsers(id int, discord_name, telegram_name string) ([]*userStorage.User, error) {
	strUsers, err := s.userRepo.GetUsers(id, discord_name, telegram_name)
	if err != nil {
		return nil, err
	}
	var u []*userStorage.User
	for _, strUser := range strUsers {
		u = append(u, &userStorage.User{
			ID:          strUser.ID,
			Name:        strUser.Name,
			DiscordName: strUser.DiscordName,
			TelegramID:  strUser.TelegramID,
			Role:        strUser.Role,
		})
	}

	return u, nil
}
