package usr

import (
	"encoding/json"
	"net/http"
)

func RequestToUsers(r *http.Request) (users []User, err error) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err = decoder.Decode(&users)

	return users, err
}