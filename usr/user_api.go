package usr

import (
	"github.com/mkadiri/golang-microservice/router"
	"net/http"
)

type UserApi struct {}

func InitRouter()  {
	router.Add("/users", UserApi{}.GetUsers, "GET")
	router.Add("/users", UserApi{}.AddUsers, "PUT")
}

func (UserApi) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := UserRepository{}.GetUsers()

	if err != nil {
		response := router.RestResponse{Message: err.Error(), StatusCode: http.StatusOK}
		response.Write(w)
		return
	}

	response := router.RestResponse{Message: users, StatusCode: http.StatusOK}
	response.Write(w)
}

func (UserApi) AddUsers(w http.ResponseWriter, r *http.Request) {
	adaptedUsers, err := RequestToUsers(r)

	if err != nil {
		response := router.RestResponse{Message: err.Error(), StatusCode: http.StatusOK}
		response.Write(w)
		return
	}

	var savedUsers []User

	for _, user := range adaptedUsers  {
		savedUser, err := UserRepository{}.AddUser(user)

		if err != nil {
			response := router.RestResponse{Message: err.Error(), StatusCode: http.StatusInternalServerError}
			response.Write(w)
			return
		}

		savedUsers = append(savedUsers, savedUser)
	}

	response := router.RestResponse{Message: savedUsers, StatusCode: http.StatusOK}
	response.Write(w)
}