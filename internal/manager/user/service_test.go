package user

import (
	"context"
	"errors"
	"testing"

	userStorage "github.com/Akezhan1/lecvisitor/internal/storage/user"
	"github.com/Akezhan1/lecvisitor/internal/storage/user/mock"
	"github.com/golang/mock/gomock"
)

func TestStorage_CreateUser(t *testing.T) {
	// defines data for CreateUser mock repo method
	type mockData struct {
		CreatedUser    any
		CreatedUserId  int
		CreatedUserErr error
	}

	var (
		ctx = context.Background()

		// create mock controller
		ctrl = gomock.NewController(t)

		testUser = User{
			ID:          1,
			Name:        "Oleg Olegovich",
			DiscordName: "oleg_olegovich",
			TelegramID:  "Oleg123",
		}
	)

	testcases := []struct {
		name     string
		mockData mockData
		input    User
		wantErr  bool
	}{
		{
			name: "new user - ok",
			mockData: mockData{
				CreatedUser:    gomock.Any(),
				CreatedUserErr: nil,
				CreatedUserId:  1,
			},
			input:   testUser,
			wantErr: false,
		},
		{
			name: "error - error from repo",
			mockData: mockData{
				CreatedUser:    gomock.Any(),
				CreatedUserId:  0,
				CreatedUserErr: errors.New("error"),
			},
			input:   testUser,
			wantErr: true,
		},
	}

	for _, testcase := range testcases {
		// create mock storage
		userStrg := mock.NewMockStorage(ctrl)

		// define behavior for mock storage
		userStrg.EXPECT().
			CreateUser(gomock.Any(), testcase.mockData.CreatedUser).
			Return(testcase.mockData.CreatedUserId, testcase.mockData.CreatedUserErr).
			Times(1)

		userSvc := NewUserService(userStrg)

		t.Run(testcase.name, func(t *testing.T) {
			_, err := userSvc.CreateUser(ctx, testcase.input)

			// if err is not nil and testcase don't wanted error - unexpected error
			if err != nil && !testcase.wantErr {
				t.Errorf("userSvc.CreateUser() error: %v", err)
			}
		})
	}
}

func TestStorage_GetUser(t *testing.T) {
	type mockData struct {
		GetUserId  int
		GetUser    *userStorage.User
		GetUserErr error
	}
	var (
		ctrl     = gomock.NewController(t)
		testUser = userStorage.User{
			ID:          1,
			Name:        "Oleg Olegovich",
			DiscordName: "oleg_olegovich",
			TelegramID:  "Oleg123",
		}
	)
	testcases := []struct {
		name           string
		mockData       mockData
		ID             int
		expectedResult *userStorage.User
		wantErr        bool
	}{
		{
			name: "name user - ok",
			mockData: mockData{
				GetUserId:  1,
				GetUser:    &testUser,
				GetUserErr: nil,
			},
			ID:             1,
			expectedResult: &testUser,
			wantErr:        false,
		},
		{
			name: "name user - error",
			mockData: mockData{
				GetUser:    &testUser,
				GetUserId:  0,
				GetUserErr: errors.New("error"),
			},
			ID:             0,
			expectedResult: &testUser,
			wantErr:        true,
		},
	}
	for _, testcase := range testcases {
		userStrg := mock.NewMockStorage(ctrl)

		userStrg.EXPECT().
			GetUser(testcase.mockData.GetUserId).
			Return(testcase.mockData.GetUser, testcase.mockData.GetUserErr).
			Times(1)
		userSvc := NewUserService(userStrg)
		t.Run(testcase.name, func(t *testing.T) {
			result, err := userSvc.GetUser(testcase.ID)
			if err != nil && err.Error() != "error" {
				t.Errorf("userSvc.GetUser() error: %v", err)
			}
			if !testcase.wantErr && result.ID != testcase.expectedResult.ID {
				t.Error("wrong data whith return GetUser")
			}
		})
	}
}

func TestStorage_GetUsers(t *testing.T) {
	type mockData struct {
		GetUsersId           int
		GetUsersDiscordName  string
		GetUsersTelegramName string
		GetUsers             []*userStorage.User
		GetUsersErr          error
	}
	var (
		ctrl     = gomock.NewController(t)
		testUser = []*userStorage.User{
			{
				ID:          1,
				Name:        "Oleg Olegovich",
				DiscordName: "oleg_olegovich",
				TelegramID:  "Oleg123",
			},
		}
	)
	testcases := []struct {
		name           string
		mockData       mockData
		ID             int
		expectedResult []*userStorage.User
		wantErr        bool
	}{
		{
			name: "name user - ok",
			mockData: mockData{
				GetUsersId:           1,
				GetUsersDiscordName:  "discord_name",
				GetUsersTelegramName: "telegram_name",
				GetUsers:             testUser,
				GetUsersErr:          nil,
			},
			ID:             1,
			expectedResult: testUser,
			wantErr:        false,
		},
		{
			name: "name user - error",
			mockData: mockData{
				GetUsersId:           0,
				GetUsersDiscordName:  "discord_name",
				GetUsersTelegramName: "telegram_name",
				GetUsers:             testUser,
				GetUsersErr:          errors.New("error"),
			},
			ID:             0,
			expectedResult: testUser,
			wantErr:        true,
		},
	}
	for index, testcase := range testcases {
		useStrg := mock.NewMockStorage(ctrl)

		useStrg.EXPECT().
			GetUsers(testcase.mockData.GetUsersId, testcase.mockData.GetUsersDiscordName, testcase.mockData.GetUsersTelegramName).
			Return(testcase.mockData.GetUsers, testcase.mockData.GetUsersErr).
			Times(1)
		userSvc := NewUserService(useStrg)
		t.Run(testcase.name, func(t *testing.T) {
			result, err := userSvc.GetUsers(testcase.mockData.GetUsersId, testcase.mockData.GetUsersDiscordName, testcase.mockData.GetUsersTelegramName)
			if err != nil && !testcase.wantErr {
				t.Errorf("userSvc.GetUsers() error: %v", err)
			}
			if !testcase.wantErr && result[index].ID != testcase.ID {
				t.Error("wrong data whith result GetUsers")
			}
		})

	}
}
