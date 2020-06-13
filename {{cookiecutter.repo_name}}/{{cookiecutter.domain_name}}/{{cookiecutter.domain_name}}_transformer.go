package {{cookiecutter.repo_name}}

import (
	"encoding/json"
	"net/http"
)

func RequestToUsers(r *http.Request) ({{cookiecutter.domain_name_plural}} []{{cookiecutter.domain_name_title}}, err error) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err = decoder.Decode(&{{cookiecutter.domain_name_plural}})

	return {{cookiecutter.domain_name_plural}}, err
}