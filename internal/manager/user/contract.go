//go:generate mockgen -source=contract.go -package=mock -destination=mock/user_mock.go
package user

import (
	"context"

	userStorage "github.com/Akezhan1/lecvisitor/internal/storage/user"
)

type (
	User struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		DiscordName string `json:"discord_name"`
		TelegramID  string `json:"telegram_id"`
		Role        string `json:"role"`
	}

	Service interface {
		CreateUser(ctx context.Context, user User) (int, error)
		GetUser(id int) (*userStorage.User, error)
		GetUsers(id int, discord_name, telegram_name string) ([]*userStorage.User, error)
	}
)
