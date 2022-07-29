package rest

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	fmt.Println("about start test cases....")
	rest.StartMockupServer()
	flag.Parse()
	os.Exit(m.Run())
}

func TestLoginUserTimeOutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8080/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email:"test@test.com", "password":"123456"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("test@test.com", "123456")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid rest client response when trying to login user", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8080/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email:"test@test.com", "password":"123456"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"invalid Login credentials", "status":"404", "error":"not_found"}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("test@test.com", "123456")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid error interface when trying to login user", err.Message)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8080/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email:"test@test.com", "password":"123456"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"invalid Login credentials", "status":404, "error":"not_found"}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("test@test.com", "123456")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "invalid Login credentials", err.Message)
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8080/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email:"test@test.com", "password":"123456"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":"1", "first_name":"test", "last_name":"user", "email":"test@test.com"}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("test@test.com", "123456")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "error when trying to unmarshall users from response", err.Message)
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8080/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email:"test@test.com", "password":"123456"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":"1", "first_name":"test", "last_name":"user", "email":"test@test.com"}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("test@test.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "test", user.FirstName)
	assert.EqualValues(t, "user", user.LastName)
	assert.EqualValues(t, "test@test.com", user.Email)
}
