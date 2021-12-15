package users

import (
	"database/sql"
)

type User struct {
	id        int
	userName  string
	interests []string
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
		`+InsertInterestsToUsersString(user.interests), user.userName)
	return err
}

func InsertInterestsToUsersString(interests []string) string {
	res := ""
	for _, v := range interests {
		res += "insert into InterestsToUsers (UserId, InterestId) values (@userId, " + v + `)
		`
	}
	return res
}
