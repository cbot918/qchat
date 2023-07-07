package main

import (
	"database/sql"

	"github.com/rs/zerolog"
)

type Service struct {
	S Storage
	L zerolog.Logger
}

func NewService(s Storage, l zerolog.Logger) Service {
	return Service{
		S: s,
		L: l,
	}
}

type LoginDbParam struct {
	Id        string `db:"id"`
	Name      string `db:"name"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	CreatedAt string `db:"created_at"`
}

func (s *Service) LoginService(l *LoginParam) (string, error) {
	v := LoginDbParam{}
	// str := fmt.Sprintf("select * from users where email='yale918@gmail.com';", l.Email)
	str := "select * from users where email=$1"
	err := s.S.Psql.Get(&v, str, l.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return "no this user", err
		}
		return "error executing db Get", err
	}
	log(v)

	// authorization to be done

	mockToken := "12345"
	return mockToken, nil
}

type ListFriendDbParam struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

func (s *Service) ListFriendService() ([]ListFriendDbParam, error) {

	q := "SELECT users.id,name FROM users JOIN friend ON users.id = to_user WHERE from_user = (select id from users where email=$1)"

	friends := []ListFriendDbParam{}

	err := s.S.Psql.Select(&friends, q, "yale918@gmail.com")
	if err != nil {
		log("friend list query db failed")
		log(err)
		return nil, err
	}
	return friends, nil
}

func (s *Service) Test() {
	if err := s.S.Psql.Ping(); err != nil {
		log("service ping failed")
		panic(err)
	}
	log("service ping psql success")
}
