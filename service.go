package main

import "database/sql"

type Service struct {
	S Storage
}

func NewService(s Storage) Service {
	return Service{
		S: s,
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

func (s *Service) Test() {
	if err := s.S.Psql.Ping(); err != nil {
		log("service ping failed")
		panic(err)
	}
	log("service ping psql success")
}
