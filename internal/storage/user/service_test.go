package user

import (
	"context"
	"testing"
)

func TestStorage_CreateUser(t *testing.T) {
	var (
		ctx = context.Background()

		newUser = User{
			Name:        "Oleg Olegovich",
			DiscordName: "oleg_olegovich",
			TelegramID:  "Oleg123",
		}

		// user from data_test.sql
		existingUser = User{
			Name:        "User User User",
			DiscordName: "Discord User",
			TelegramID:  "Telegram User",
		}
	)

	testcases := []struct {
		name    string
		input   User
		wantErr bool
	}{
		{
			name:    "new user - ok",
			input:   newUser,
			wantErr: false,
		},
		{
			name:    "existing user - error",
			input:   existingUser,
			wantErr: true,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			// doing test query to test storage
			_, err := testStg.CreateUser(ctx, testcase.input)

			// if err is not nil and testcase don't wanted error - unexpected error
			if err != nil && !testcase.wantErr {
				t.Errorf("testStg.CreateUser() error: %v", err)
			}
		})
	}
}

func TestStorage_GetUser(t *testing.T) {
	var (
		newUser = User{
			ID:          1,
			Name:        "Oleg Olegovich",
			DiscordName: "oleg_olegovich",
			TelegramID:  "Oleg123",
		}

		// user from data_test.sql
		existingUser = User{
			ID:          0,
			Name:        "User User User",
			DiscordName: "Discord User",
			TelegramID:  "Telegram User",
		}
	)

	testcases := []struct {
		name    string
		input   User
		wantErr bool
	}{
		{
			name:    "new user - ok",
			input:   newUser,
			wantErr: false,
		},
		{
			name:    "existing user - error",
			input:   existingUser,
			wantErr: true,
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			_, err := testStg.GetUser(testcase.input.ID)

			if err != nil && !testcase.wantErr {
				t.Errorf("teststg.GetUser() error: %v", err)
			}
		})
	}
}

func TestStorage_GetUsers(t *testing.T) {
	var (
		newUser = User{
			ID:          1,
			Name:        "Oleg Olegovich",
			DiscordName: "oleg_olegovich",
			TelegramID:  "Oleg123",
		}

		// user from data_test.sql
		existingUser = User{
			ID:          0,
			Name:        "User User User",
			DiscordName: "Discord User",
			TelegramID:  "Telegram User",
		}
	)

	testcases := []struct {
		name    string
		input   User
		wantErr bool
	}{
		{
			name:    "new user - ok",
			input:   newUser,
			wantErr: false,
		},
		{
			name:    "existing user - error",
			input:   existingUser,
			wantErr: true,
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			_, err := testStg.GetUsers(testcase.input.ID, testcase.input.DiscordName, testcase.input.TelegramID)

			if err != nil && !testcase.wantErr {
				t.Errorf("teststg.GetUsers() error : %v", err)
			}
		})
	}
}
