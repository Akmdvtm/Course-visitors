package user

import (
	"bytes"
	"encoding/json"
	"github.com/Akezhan1/lecvisitor/internal/manager/user"
	"github.com/Akezhan1/lecvisitor/internal/manager/user/mock"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHandler_CreateUser(t *testing.T) {
	// defines data for CreateUser mock repo method
	type mockData struct {
		CreatedUser    any
		CreatedUserId  int
		CreatedUserErr error
	}

	var (
		ctrl = gomock.NewController(t)

		testUser = user.User{
			ID:          1,
			Name:        "Oleg Olegovich",
			DiscordName: "oleg_olegovich",
			TelegramID:  "Oleg123",
		}
	)

	testcases := []struct {
		name              string
		mockData          mockData
		httpMethod        string             // http method (GET, POST etc)
		input             user.User          // input for http.Body
		wantedResponse    createUserResponse // response from createUserHandler
		wantedErrResponse errResponse        // err response from creteUserHandler
		wantedStatusCode  int                // status code from createUserHandler
	}{
		{
			name: "ok",
			mockData: mockData{
				CreatedUser:    testUser,
				CreatedUserId:  1,
				CreatedUserErr: nil,
			},
			httpMethod: http.MethodPost,
			input:      testUser,
			wantedResponse: createUserResponse{
				ID: testUser.ID,
			},
			wantedErrResponse: errResponse{},
			wantedStatusCode:  http.StatusCreated,
		},
		{
			name:              "mismatch method",
			httpMethod:        http.MethodGet,
			wantedErrResponse: newErrResponse(errMethodNotAllowed),
			wantedStatusCode:  http.StatusMethodNotAllowed,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {

			// defines user mock service
			userSvc := mock.NewMockService(ctrl)

			userSvc.EXPECT().
				CreateUser(gomock.Any(), testcase.mockData.CreatedUser).
				Return(testcase.mockData.CreatedUserId, testcase.mockData.CreatedUserErr).
				Times(1)

			// create handler
			userHandler := NewHandler(userSvc)

			// create test http server
			server := httptest.NewServer(userHandler.ServeMux())
			defer server.Close()

			// convert user model to bytes for write to http.Body
			inputBytes, err := json.Marshal(testcase.input)
			if err != nil {
				t.Error(err)
				return
			}

			// http body
			body := bytes.NewReader(inputBytes)

			// create http client
			client := http.DefaultClient

			// create http request
			req, err := http.NewRequest(testcase.httpMethod, server.URL+"/create", body)
			if err != nil {
				t.Error(err)
				return
			}

			// do http request
			resp, err := client.Do(req)
			if err != nil {
				t.Error(err)
				return
			}

			// check status codes
			if resp.StatusCode != testcase.wantedStatusCode {
				t.Errorf("mismatch status code. got: %v, want: %v", resp.StatusCode, testcase.wantedStatusCode)
				return
			}

			// status codes from 200 to 299 is successful codes
			// if status code > 299 is error status code it means that body is error response
			if resp.StatusCode > 299 {
				var errResp errResponse

				if err = json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
					t.Error(err)
					return
				}

				// check that returned response is equal to wanted response
				if !reflect.DeepEqual(errResp, testcase.wantedErrResponse) {
					t.Errorf("mismatch err response. got: %v, want: %v", errResp, testcase.wantedErrResponse)
					return
				}
			} else {
				var createResp createUserResponse

				if err = json.NewDecoder(resp.Body).Decode(&createResp); err != nil {
					t.Error(err)
					return
				}

				// check that returned response is equal to wanted response
				if !reflect.DeepEqual(createResp, testcase.wantedResponse) {
					t.Errorf("mismatch err response. got: %v, want: %v", createResp, testcase.wantedResponse)
					return
				}
			}
		})
	}
}
