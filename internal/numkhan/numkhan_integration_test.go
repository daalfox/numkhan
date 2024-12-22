package numkhan

import (
	"os"
	"testing"
)

func TestNewVoter(t *testing.T) {
	dbName := "test_db.db"
	db, _ := SetupDb(dbName)
	defer os.Remove(dbName)

	service := Service{Db: db}

	votesBefore, _ := service.Votes(0)

	id, _ := service.Subscribe()
	service.Vote(id, 0)

	votesAfter, _ := service.Votes(0)

	if votesAfter != votesBefore+1 {
		t.Errorf("expected votes to go from %d to %d. before: %d, after: %d", votesBefore, votesBefore+1, votesBefore, votesAfter)
	}
}

func TestPreventVotingMultipleTimes(t *testing.T) {
	dbName := "test_db.db"
	db, _ := SetupDb(dbName)
	defer os.Remove(dbName)

	service := Service{Db: db}

	id, _ := service.Subscribe()

	// first vote
	service.Vote(id, 0)
	// second vote
	err := service.Vote(id, 0)
	if err.Error() != AlreadyVoted.Error() {
		t.Error("expected to fail on second vote, but didn't")
	}
}
