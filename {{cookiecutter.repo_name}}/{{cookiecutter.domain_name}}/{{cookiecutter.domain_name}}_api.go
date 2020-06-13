package {{cookiecutter.domain_name}}

import (
	"github.com/mkadiri/golang-microservice/router"
	"net/http"
)

type UserApi struct {}

func InitRouter()  {
	router.Add("/{{cookiecutter.domain_name_plural}}", UserApi{}.GetUsers, "GET")
	router.Add("/{{cookiecutter.domain_name_plural}}", UserApi{}.AddUsers, "PUT")
}

func (UserApi) GetUsers(w http.ResponseWriter, r *http.Request) {
	{{cookiecutter.domain_name_plural}}, err := UserRepository{}.GetUsers()

	if err != nil {
		response := router.RestResponse{Message: err.Error(), StatusCode: http.StatusOK}
		response.Write(w)
		return
	}

	response := router.RestResponse{Message: {{cookiecutter.domain_name_plural}}, StatusCode: http.StatusOK}
	response.Write(w)
}

func (UserApi) AddUsers(w http.ResponseWriter, r *http.Request) {
	adaptedUsers, err := RequestToUsers(r)

	if err != nil {
		response := router.RestResponse{Message: err.Error(), StatusCode: http.StatusOK}
		response.Write(w)
		return
	}

	var savedUsers []{{cookiecutter.domain_name_title}}

	for _, {{cookiecutter.domain_name}} := range adaptedUsers  {
		savedUser, err := UserRepository{}.AddUser({{cookiecutter.domain_name}})

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