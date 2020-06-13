package {{cookiecutter.repo_name}}

import (
	"github.com/mkadiri/golang-microservice/database"
)

type UserRepository struct {}

func (UserRepository) GetUsers() ({{cookiecutter.domain_name_plural}} []{{cookiecutter.domain_name_title}}, err error) {
	rows, err := database.Db.Query("select * from `{{cookiecutter.domain_name}}`")

	if err != nil {
		return {{cookiecutter.domain_name_plural}}, err
	}

	for rows.Next() {
		var {{cookiecutter.domain_name}} {{cookiecutter.domain_name_title}}
		err = rows.Scan(&{{cookiecutter.domain_name}}.Id, &{{cookiecutter.domain_name}}.FirstName, &{{cookiecutter.domain_name}}.LastName, &{{cookiecutter.domain_name}}.DateOfBirth)

		if err != nil {
			return {{cookiecutter.domain_name_plural}}, err
		}

		{{cookiecutter.domain_name_plural}} = append({{cookiecutter.domain_name_plural}}, {{cookiecutter.domain_name}})
	}

	return {{cookiecutter.domain_name_plural}}, err
}

func (userRepository UserRepository) AddUser({{cookiecutter.domain_name}} {{cookiecutter.domain_name_title}}) (savedUser {{cookiecutter.domain_name_title}}, err error) {
	stmt, err := database.Db.Prepare("insert into `{{cookiecutter.domain_name}}` (first_name, last_name, date_of_birth) values(?, ?, ?)")
	defer stmt.Close()
	
	if err != nil {
		return savedUser, err
	}
	
	res, err := stmt.Exec({{cookiecutter.domain_name}}.FirstName, {{cookiecutter.domain_name}}.LastName, {{cookiecutter.domain_name}}.DateOfBirth)

	if err != nil {
		return savedUser, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return savedUser, err
	}

	savedUser, err = userRepository.GetUserById(int(id))
	
	return savedUser, err
}

func (UserRepository) GetUserById(id int) ({{cookiecutter.domain_name}} {{cookiecutter.domain_name_title}}, err error) {
	err = database.Db.QueryRow("select * from `{{cookiecutter.domain_name}}` where `id` = ?", id).
		Scan(&{{cookiecutter.domain_name}}.Id, &{{cookiecutter.domain_name}}.FirstName, &{{cookiecutter.domain_name}}.LastName, &{{cookiecutter.domain_name}}.DateOfBirth)

	return {{cookiecutter.domain_name}}, err
}