package numkhan

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Candidate struct {
	gorm.Model
	N     int
	Votes uint
}

func SetupDb(name string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(name))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Candidate{})

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
	uuid := uuid.New()

	return uuid, nil
}

func (s *Service) Vote(id uuid.UUID, n int) error {
	// TODO: check if user is voted

	c := Candidate{N: n}
	s.Db.First(&c)
	c.Votes += 1
	s.Db.Save(&c)

	return nil
}
