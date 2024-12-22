package service

import (
	"os"
	"testing"

	"github.com/daalfox/numkhan/internal/numkhan"
)

func TestNewVoter(t *testing.T) {
	dbName := "test_db.db"
	db, _ := numkhan.SetupDb(dbName)
	defer os.Remove(dbName)

	service := numkhan.Service{Db: db}

	votesBefore, _ := service.Votes(0)

	id, _ := service.Subscribe()
	service.Vote(id, 0)

	votesAfter, _ := service.Votes(0)

	if votesAfter != votesBefore+1 {
		t.Errorf("expected votes to go from %d to %d. before: %d, after: %d", votesBefore, votesBefore+1, votesBefore, votesAfter)
	}
}
