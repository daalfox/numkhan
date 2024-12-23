package numkhan

import (
	"os"
	"reflect"
	"testing"
)

func TestSetupDb(t *testing.T) {
	db, _ := SetupDb("test_db.db")
	defer os.Remove("test_db.db")

	var candidates []Candidate
	db.Find(&candidates)

	var numbers []int
	for _, c := range candidates {
		numbers = append(numbers, c.N)
	}

	got := numbers
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v\ngot %v", want, got)
	}
}

func TestGetCandidates(t *testing.T) {
	db, _ := SetupDb("test_db.db")
	defer os.Remove("test_db.db")

	service := Service{db}

	candidates := service.Candidates()

	var numbers []int
	for _, c := range candidates {
		numbers = append(numbers, c.N)
	}

	got := numbers
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v\ngot %v", want, got)
	}
}
