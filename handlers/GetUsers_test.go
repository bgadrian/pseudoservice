package handlers

import (
	"fmt"
	"github.com/bgadrian/pseudoservice/models"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func isEmptyString(strings ...string) bool {
	for _, str := range strings {
		if len(str) == 0 {
			return true
		}
	}
	return false
}

func areValid(users []*models.User) bool {
	for _, user := range users {
		if user == nil {
			return false
		}

		//TODO replace with user.Validate()
		if isEmptyString(*user.Name, user.Email, user.Position,
			user.Company, user.Country) {
			return false
		}
	}
	return true
}

func TestGenerateUsersDeterministic(t *testing.T) {
	seeds := []int64{-422, 0, 1, 100, 123456123456}

	randUsers, _, err := GenerateUsers(-999999, 30, false)
	if err != nil {
		t.Error(err)
	}

	if len(randUsers) != 30 {
		t.Error("generateUsers failed to return {count} users")
	}

	for _, seed := range seeds {
		usersA, nextA, err := GenerateUsers(seed, 30, true)
		if err != nil {
			t.Error(err)
		}
		usersB, nextB, err := GenerateUsers(seed, 30, true)
		if err != nil {
			t.Error(err)
		}

		if len(usersA) != 30 {
			t.Error("generateUsers failed to return {count} users")
		}

		if nextA != nextB {
			t.Errorf("different nextSeed for seed:%d count %d", seed, 3)
		}

		if cmp.Equal(usersA, usersB) == false {
			t.Errorf("different users for seed:%d count %d", seed, 3)
			fmt.Printf("A: '%v' \nB: '%v'\n", usersA[0], usersB[0])
		}

		if cmp.Equal(usersA, randUsers) {
			t.Error("is generating the same users for any seed")
		}
	}

}
func TestGenerateUsersHasData(t *testing.T) {
	sizes := []int{0, 1, 100, 1000}
	for _, size := range sizes {
		users, _, err := GenerateUsers(42, size, false)

		if err != nil {
			t.Error(err)
			return
		}

		if len(users) != size {
			t.Errorf("exp %d users, got %d", size, len(users))
		}

		if areValid(users) == false {
			t.Errorf("empty users generated for seed: %d count %d sample: %v",
				42, size, users[0])
		}
	}
}

func TestUsersCountGetSeed(t *testing.T) {
	_, nextA, err := GenerateUsers(42, 30, true)
	if err != nil {
		t.Error(err)
	}
	if nextA != 42+30 {
		t.Errorf("next seed is wrong, exp %d got %d", 42+30, nextA)
	}

	_, nextA, err = GenerateUsers(42, 30, false)
	if err != nil {
		t.Error(err)
	}
	if nextA != 0 {
		t.Errorf("next seed is wrong, exp %d got %d", 0, nextA)
	}
}

func BenchmarkGenerateUsers100Deterministic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateUsers(42, 100, true)
	}
}

func BenchmarkGenerateUsers100NonDeterministic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateUsers(42, 100, false)
	}
}

func BenchmarkGenerateUsers300Deterministic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateUsers(42, 300, true)
	}
}

func BenchmarkGenerateUsers300NonDeterministic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GenerateUsers(42, 300, false)
	}
}
