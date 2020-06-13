package {{cookiecutter.repo_name}}

import (
	"bytes"
	"encoding/json"
	"github.com/mkadiri/golang-microservice/database"
	"github.com/mkadiri/golang-microservice/router"
	"net/http"
	"testing"
	"time"
)

// these numbers aren't random https://golang.org/src/time/format.go
var dob, _ = time.Parse("2006-01-02", "1989-10-08")

var {{cookiecutter.domain_name_plural}} = []{{cookiecutter.domain_name_title}}{
	{Id: 1, FirstName: "james", LastName: "smith", DateOfBirth: dob},
	{Id: 2, FirstName: "bruce", LastName: "wayne", DateOfBirth: dob},
	{Id: 3, FirstName: "jeremy", LastName: "renner", DateOfBirth: dob},
}

func TestAddUsers(t *testing.T) {
	router.Init()
	InitRouter()

	database.Db.Exec("truncate table {{cookiecutter.domain_name}}")

	modulesJson, err := json.Marshal({{cookiecutter.domain_name_plural}})

	if err != nil {
		t.Errorf("Cannot encode to JSON")
	}

	req, _ := http.NewRequest("PUT", "/{{cookiecutter.domain_name_plural}}", bytes.NewBuffer(modulesJson))
	response := router.ApiTestHelper{}.ExecuteRequest(req)

	router.ApiTestHelper{}.CheckResponseCode(t, http.StatusOK, response.Code)
	router.ApiTestHelper{}.CheckBodyEqualsJson(t, modulesJson, response.Body)
}

func TestGetUsers(t *testing.T) {
	router.Init()
	InitRouter()

	database.Db.Exec("truncate table {{cookiecutter.domain_name}}")
	req, _ := http.NewRequest("GET", "/{{cookiecutter.domain_name_plural}}", nil)
	response := router.ApiTestHelper{}.ExecuteRequest(req)

	router.ApiTestHelper{}.CheckResponseCode(t, http.StatusOK, response.Code)
}