package users

import (
	"database/sql"
)

type User struct {
	Id        int      `json:"id"`
	UserName  string   `json:"UserName"`
	Interests []string `json:"Interests"`
}

type Store struct {
	Db *sql.DB
}

func GetStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) CreateUser(user *User) error {
	_, err := s.Db.Query(
		`declare @userName varchar(50) = ?
		insert into Users (UserName) values (@userName)
		declare @userId int = (select id from Users where UserName = @userName)
		`+InsertInterestsToUsersString(user.Interests), GetQueryArgs(user)...)
	return err
}

func GetQueryArgs(user *User) []interface{} {
	res := make([]interface{}, len(user.Interests)+1)
	res[0] = user.UserName

	for i, v := range user.Interests {
		res[i+1] = v
	}

	return res
}

func InsertInterestsToUsersString(interests []string) string {
	res := ""
	for range interests {
		res += `insert into InterestsToUsers (UserId, InterestId) values (@userId, ?)
		`
	}
	return res
}
