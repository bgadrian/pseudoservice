package users

import (
	"fmt"
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

func areValid(users []*User) bool {
	for _, user := range users {
		if user == nil {
			return false
		}

		if isEmptyString(user.Name, user.Email, user.Position,
			user.Company, user.Country) {
			return false
		}
	}
	return true
}

func TestGenerateUsersDeterministic(t *testing.T) {
	seeds := []int64{-422, 0, 1, 100, 123456123456}
	for _, seed := range seeds {
		usersA, nextA, err := GenerateUsers(seed, 3)
		if err != nil {
			t.Error(err)
		}
		usersB, nextB, err := GenerateUsers(seed, 3)
		if err != nil {
			t.Error(err)
		}

		if nextA != nextB {
			t.Errorf("different nextSeed for seed:%d count %d", seed, 3)
		}

		if cmp.Equal(usersA, usersB) == false {
			t.Errorf("different users for seed:%d count %d", seed, 3)
			fmt.Printf("A: '%v' \nB: '%v'\n", usersA[0], usersB[0])
		}
	}

}
func TestGenerateUsersHasData(t *testing.T) {
	sizes := []int{0, 1, 100, 1000}
	for _, size := range sizes {
		users, _, err := GenerateUsers(42, size)

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
