package numkhan

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	AlreadyVoted = NumkhanErr("you can't vote multiple times")
)

type NumkhanErr string

func (e NumkhanErr) Error() string {
	return string(e)
}

type Candidate struct {
	gorm.Model
	N     int
	Votes uint
}
type User struct {
	gorm.Model
	Uuid    uuid.UUID
	IsVoted bool
}

func SetupDb(name string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(name))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Candidate{})
	db.AutoMigrate(&User{})

	for n := range 10 {
		row := Candidate{N: n, Votes: 0}
		db.Create(&row)
	}

	return db, nil
}

type Service struct {
	Db *gorm.DB
}

func (s *Service) Votes(n int) (uint, error) {
	c := Candidate{N: n}

	s.Db.First(&c)

	return c.Votes, nil
}

func (s *Service) Subscribe() (uuid.UUID, error) {
	user := User{
		Uuid:    uuid.New(),
		IsVoted: false,
	}

	s.Db.Create(&user)

	return user.Uuid, nil
}

func (s *Service) Vote(id uuid.UUID, n int) error {
	user := User{Uuid: id}
	s.Db.First(&user)
	if user.IsVoted {
		return AlreadyVoted
	}

	c := Candidate{N: n}
	s.Db.First(&c)
	c.Votes += 1
	user.IsVoted = true
	s.Db.Save(&c)
	s.Db.Save(&user)

	return nil
}
