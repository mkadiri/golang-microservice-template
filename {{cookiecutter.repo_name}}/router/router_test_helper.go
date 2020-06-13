package router

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ApiTestHelper struct {}

func (ApiTestHelper) ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	Router.ServeHTTP(responseRecorder, req)

	return responseRecorder
}

func (ApiTestHelper) CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected api_response code %d. Got %d\n", expected, actual)
	}
}

// json.NewEncoder(w).Encode() adds an extra line (\n), we should do the same here when making a comparison
func (ApiTestHelper) CheckBodyEqualsJson(t *testing.T, modulesJson []byte, reponseBody *bytes.Buffer) {
	modulesJsonString := string(modulesJson) + "\n"
	body := reponseBody.String()

	if body != modulesJsonString {
		t.Errorf("Expected %s", modulesJsonString)
		t.Errorf("Got %s", body)
	}
}