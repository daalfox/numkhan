package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/daalfox/numkhan/internal/numkhan"
)

func TestRestAPI(t *testing.T) {
	dbName := "numkhan.db"
	db, _ := numkhan.SetupDb(dbName)
	defer os.Remove(dbName)

	s := numkhan.NewHttpService(db)

	t.Run("Get Candidates endpoint", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		s.GetCandidates(res, req)

		assertStatus(t, res.Code, http.StatusOK)

		candidates := []numkhan.Candidate{}
		for n := range 10 {
			candidates = append(candidates, numkhan.Candidate{ID: n + 1, N: n, Votes: 0})
		}

		var got []numkhan.Candidate
		json.NewDecoder(res.Body).Decode(&got)
		want, _ := json.Marshal(candidates)

		if reflect.DeepEqual(want, got) {
			t.Errorf("\ngot: %v\nwnt: %v", got, want)
		}
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected %d status code, but got %d", want, got)
	}
}
