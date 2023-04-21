package user

import (
	userStorage "github.com/Akezhan1/lecvisitor/internal/storage/user"
)

func (u User) toStorageModel() userStorage.User {
	return userStorage.User{
		ID:          u.ID,
		Name:        u.Name,
		DiscordName: u.DiscordName,
		TelegramID:  u.TelegramID,
		Role:        u.Role,
	}
}
