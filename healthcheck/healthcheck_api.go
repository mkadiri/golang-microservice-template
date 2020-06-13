package healthcheck

import (
	"github.com/mkadiri/golang-microservice/router"
	"net/http"
)

func InitRouter() {
	router.Add("/healthcheck", HealthCheck,"GET")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := router.RestResponse{Message: "OK", StatusCode: http.StatusOK}
	response.Write(w)
}