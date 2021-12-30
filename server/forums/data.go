package forums

import (
	"database/sql"
	"strings"
)

type Store struct {
	Db *sql.DB
}

type QueryForum struct {
	Name         string `json:"Name"`
	TopicKeyword string `json:"TopicKeyword"`
	UsersString  string `json:"UsersString"`
}

type Forum struct {
	Name         string   `json:"Name"`
	TopicKeyword string   `json:"TopicKeyword"`
	Users        []string `json:"Users"`
}

func GetStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) ListForums() ([]*Forum, error) {
	rows, err := s.Db.Query(
		`select f.Name, i.Name, string_agg(u.UserName, ';') from Forums f
		inner join Interests i on i.id = f.InterestId
		inner join UsersToForums utf on utf.ForumId = f.id
		inner join Users u on u.id = utf.UserId
		group by f.Name, i.Name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*Forum
	for rows.Next() {
		var forum QueryForum
		if err := rows.Scan(&forum.Name, &forum.TopicKeyword, &forum.UsersString); err != nil {
			return nil, err
		}
		res = append(res, forum.ToForum())
	}
	if res == nil {
		res = make([]*Forum, 0)
	}
	return res, nil
}

func (qforum *QueryForum) ToForum() *Forum {
	return &Forum{
		Name:         qforum.Name,
		TopicKeyword: qforum.TopicKeyword,
		Users:        strings.Split(qforum.UsersString, ";"),
	}
}
