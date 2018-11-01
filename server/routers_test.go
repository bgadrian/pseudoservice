package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bgadrian/pseudoservice/users"
	"github.com/gorilla/mux"
)

func isEmptyString(strings ...string) bool {
	for _, str := range strings {
		if len(str) == 0 {
			return true
		}
	}
	return false
}

func areValid(users []*users.User) bool {
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

func TestUsersCountGetSeed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://dummy?seed=42", nil)
	request = mux.SetURLVars(request, map[string]string{"count": "100"})
	w := httptest.NewRecorder()
	UsersCountGet(w, request)

	if w.Code != http.StatusOK {
		fmt.Printf("body: '%s'", string(w.Body.Bytes()))
		t.Errorf("status code exp %d got %d", http.StatusOK, w.Code)
	}

	responseStruct := &ResponseModel{}
	json.Unmarshal(w.Body.Bytes(), responseStruct)

	if responseStruct.Seed != 42 {
		fmt.Printf("body: '%s'", string(w.Body.Bytes()))
		t.Errorf("seed exp %d got %d", 42, responseStruct.Seed)
	}

	if responseStruct.Nextseed != 142 {
		t.Errorf("nextseed exp %d got %d", 142, responseStruct.Nextseed)
	}

	if len(responseStruct.Users) != 100 {
		t.Errorf("wrong count of users, exp %d got %d", 100, len(responseStruct.Users))
		return
	}

	if areValid(responseStruct.Users) == false {
		t.Errorf("uesrs are empty, malformed data")
	}
}

func TestUsersCountGetRandom(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://dummy", nil)
	request = mux.SetURLVars(request, map[string]string{"count": "3"})
	w := httptest.NewRecorder()
	UsersCountGet(w, request)

	if w.Code != http.StatusOK {
		fmt.Printf("body: '%s'", string(w.Body.Bytes()))
		t.Errorf("status code exp %d got %d", http.StatusOK, w.Code)
	}

	responseStruct := &ResponseModel{}
	json.Unmarshal(w.Body.Bytes(), responseStruct)

	if len(responseStruct.Users) != 3 {
		t.Errorf("wrong count of users, exp %d got %d", 3, len(responseStruct.Users))
		return
	}

	if areValid(responseStruct.Users) == false {
		t.Errorf("uesrs are empty, malformed data")
	}
}

func TestUsersCountGetWrongInput(t *testing.T) {
	for _, test := range []struct{ url, count string }{
		{"http://dummy", "-1"},
		{"http://dummy?seed=alfa", "3"},
	} {

		request := httptest.NewRequest(http.MethodGet, test.url, nil)
		request = mux.SetURLVars(request, map[string]string{"count": test.count})
		w := httptest.NewRecorder()
		UsersCountGet(w, request)

		if w.Code == http.StatusOK {
			t.Errorf("wrong count, exp error code, got 200 for %s count:%s", test.url, test.count)
		}

	}
}
