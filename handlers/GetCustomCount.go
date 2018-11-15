package handlers

import (
	"regexp"
	"strings"

	"github.com/bgadrian/fastfaker"
	"github.com/bgadrian/pseudoservice/models"
	"github.com/bgadrian/pseudoservice/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

// keyPos the position of a valid key in the template, and the function to call
//to replace it with random data
type keyPos struct {
	start, end int
	f          fakerer
}

// GetCustomCountParams sanitize the HTTP input and add the GenerateCustom result to the HTTP response.
func (h *MyHandlers) GetCustomCountParams(params operations.GetCustomCountParams, principal interface{}) middleware.Responder {

	var seed int64
	seedGiven := params.Seed != nil
	if seedGiven {
		seed = *params.Seed
	} else {
		seed = fastfaker.Global.Int64()
	}
	template := params.Template
	count := params.Count

	result := GenerateCustom(seed, template, count)

	return operations.NewGetCustomCountOK().WithPayload(&models.CustomResponseModel{
		Seed:    seed,
		Results: result,
	})
	return nil
}

var pattern = regexp.MustCompile(`(~[a-zA-Z0-9_-]+~)`)

// GenerateCustom will replace all known ~keys~ from the template with random data, based on the seed. The result will contain (count) such strings.
//For each unique seed/template/count group the result should be the same (deterministic).
func GenerateCustom(seed int64, template string, count int32) []string {
	templateAsByte := []byte(template)
	indexes := pattern.FindAllIndex(templateAsByte, -1)
	result := make([]string, 0, count)

	//filter keys
	var toReplace []keyPos
	for _, match := range indexes {
		start := match[0]
		end := match[1]

		key := string(templateAsByte[start:end])
		if len(key) < 3 {
			//key is empty "~~"
			continue
		}
		//remove the first and last ~
		inner := strings.ToLower(key[1 : len(key)-1])

		f, exists := keys[inner]
		if exists == false {
			//key does not exists
			continue
		}
		toReplace = append(toReplace, keyPos{start, end, f})
	}

	//add random data
	faker := fastfaker.NewFastFaker()
	faker.Seed(seed)
	var i int32
	for ; i < count; i++ {
		if len(indexes) == 0 {
			result = append(result, template)
			continue
		}

		buff := strings.Builder{}
		buff.Grow(len(templateAsByte)) //at least the input

		var lastEnd int
		for _, posToReplace := range toReplace {
			buff.Write(templateAsByte[lastEnd:posToReplace.start])
			buff.WriteString(posToReplace.f(faker))
			lastEnd = posToReplace.end
		}
		buff.Write(templateAsByte[lastEnd:])

		result = append(result, buff.String())
	}
	return result
}
