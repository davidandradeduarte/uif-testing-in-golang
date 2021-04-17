package main

import (
	"errors"
	"fmt"

	"github.com/davidandradeduarte/uif-testing-in-golang-udemy/src/sql-client/sqlclient"
)

const (
	queryGetUser = "SELECT id, email FROM users WHERE id=%d"
)

var (
	dbClient sqlclient.SqlClient
)

type User struct {
	Id    int64
	Email string
}

func init() {
	sqlclient.StartMockServer()
	var err error
	dbClient, err = sqlclient.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "root", "127.0.0.1:3306", "users_db"))
	if err != nil {
		panic(err)
	}
}

func main() {
	user, err := GetUser(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(user.Email)
}

func GetUser(id int64) (*User, error) {
	sqlclient.AddMock(sqlclient.Mock{
		Query:   "SELECT id, email FROM users WHERE id=?;",
		Args:    []interface{}{1},
		Error:   errors.New("error creating query"),
		Columns: []string{"id", "email"},
		Rows: [][]interface{}{
			{1, "email"},
			{2, "email"},
		},
	})
	rows, err := dbClient.Query(fmt.Sprintf(queryGetUser, id))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var user User
	for rows.HasNext() {
		if err := rows.Scan(&user.Id, &user.Email); err != nil {
			return nil, err
		}
		return &user, nil
	}

	return nil, errors.New("User not found")
}
