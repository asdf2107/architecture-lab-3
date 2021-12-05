package forums

import (
	"database/sql"
	"strings"
)

type Store struct {
	Db *sql.DB
}

type QueryForum struct {
	name         string
	topicKeyword string
	usersString  string
}

type Forum struct {
	name         string
	topicKeyword string
	users        []string
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
		if err := rows.Scan(&forum.name, &forum.topicKeyword, &forum.usersString); err != nil {
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
		name:         qforum.name,
		topicKeyword: qforum.topicKeyword,
		users:        strings.Split(qforum.usersString, ";"),
	}
}
