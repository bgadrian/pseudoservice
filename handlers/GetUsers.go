package handlers

import (
	"fmt"
	"math"
	"net/http"
	"strings"

	"github.com/bgadrian/pseudoservice/models"
	"github.com/bgadrian/pseudoservice/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"github.com/bgadrian/fastfaker"
)

// GetUsersHandler /users/{count} handler with real data
func (h *MyHandlers) GetUsersHandler(params operations.GetUsersCountParams, principal interface{}) (result middleware.Responder) {
	defer func() {
		if err := recover(); err != nil {
			errorResponse(h, fmt.Errorf("recovered in %v", err))
		}
	}()

	var seed int64
	seedGiven := params.Seed != nil
	if seedGiven {
		seed = *params.Seed
	} else {
		seed = fastfaker.Global.Int64()
	}

	users, nextseed, err := GenerateUsers(seed, int(params.Count), seedGiven)
	if err == nil {
		response := &models.UserResponseModel{}
		response.Seed = seed
		response.Nextseed = nextseed
		response.Users = users
		return operations.NewGetUsersCountOK().WithPayload(response)
	}

	errorResponse := errorResponse(h, err)
	return errorResponse
}

func errorResponse(h *MyHandlers, err error) *operations.GetUsersCountDefault {
	h.Api.Logger("error generateusers '%s'", err)
	code := int32(42)
	message := "internal generator failed"
	response := &models.ErrorModel{&code, &message}
	errorResponse := operations.NewGetUsersCountDefault(http.StatusInternalServerError).WithPayload(response)
	return errorResponse
}

//GenerateUsers deterministic generation of (random) users.
func GenerateUsers(seed int64, count int, deterministic bool) ([]*models.User, int64, error) {
	if math.MaxInt64-int64(count) <= seed {
		//chance to being here is like ... 2^63-count ...is like winning the lottery
		return nil, 0, fmt.Errorf("int overflow, need a smaller seed: %d count: %d", seed, count)
	}
	faker := fastfaker.NewFastFaker()

	if deterministic == false {
		faker.Seed(seed)
	}

	result := make([]*models.User, 0, count)
	friendsIndexs := make([]int, 0, count)

	for i := 0; i < count; i++ {
		//ALERT as long as the following lines remain in order, the calls will be deterministic/backward
		//compatible.

		if seed == 0 {
			seed++ //gofakeit treats 0 as newRandom
		}

		user := &models.User{}
		if deterministic {
			//each seed value must return a specific user, with same data
			//but the performance is 30% worst
			faker.Seed(seed)
		}
		ID := strfmt.UUID(faker.UUID())
		user.ID = &ID
		name := faker.Name()
		user.Name = &name
		//user.Age = faker.Uint8()
		user.Company = faker.BuzzWord() + " " +
			faker.BS() + " " + faker.CompanySuffix()
		user.Position = faker.JobDescriptor() + " " +
			faker.JobLevel() + " " + faker.JobTitle()

		user.Email = strings.Replace(name, " ", "", -1) +
			"@" + faker.DomainName()
		user.Country = faker.Country()

		//FRIENDS from the same batch
		zeroTendency := len(result) / 3                           //at least 33% will have 0 friends
		friendCount := faker.Number(-zeroTendency, len(result)/2) //max of half of users so far

		if friendCount > 0 {
			user.Friends = make([]string, 0, friendCount)
			fcount := 0
			//shuffle them more rarely, for perf reasons
			if i%10 == 0 {
				faker.ShuffleInts(friendsIndexs)
			} else {
				//we do not shuffle it, but we start from a random friend
				fcount = faker.Number(0, len(friendsIndexs)-friendCount-1)
			}

			for ; fcount < friendCount; fcount++ {
				friend := result[friendsIndexs[fcount]]
				user.Friends = append(user.Friends, friend.ID.String())   //him -> me
				friend.Friends = append(friend.Friends, user.ID.String()) //me -> him
			}
		}

		seed++
		result = append(result, user)
		friendsIndexs = append(friendsIndexs, i)
	}

	nextSeed := seed
	if deterministic == false {
		nextSeed = 0
	}

	return result, nextSeed, nil
}
