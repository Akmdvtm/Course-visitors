//go:generate mockgen -source=contract.go -package=mock -destination=mock/user_mock.go
package user

import (
	"context"
	"time"
)

type (
	// User defines user database model
	User struct {
		ID          int
		Name        string
		DiscordName string
		TelegramID  string
		Role        string
		CreatedAt   time.Time
	}

	// Storage defines user storage methods
	Storage interface {
		// CreateUser creates user in database and returns inserted id and error
		CreateUser(ctx context.Context, user User) (int, error)
		// GetUser returns user information by its id
		GetUser(id int) (*User, error)
		// getUsers return users infomation by its id, telegram and discord name
		GetUsers(id int, discord_name, telegram_name string) ([]*User, error)
	}
)
