package healthcheck

import (
	"github.com/mkadiri/golang-microservice/router"
	"net/http"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	router.Init()
	InitRouter()
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	response := router.ApiTestHelper{}.ExecuteRequest(req)
	router.ApiTestHelper{}.CheckResponseCode(t, http.StatusOK, response.Code)
}