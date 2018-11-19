package handlers

import (
	"log"

	"github.com/bgadrian/fastfaker/faker"
	"github.com/bgadrian/pseudoservice/models"
	"github.com/bgadrian/pseudoservice/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

// GetCustomCountParams sanitize the HTTP input and add the GenerateCustom result to the HTTP response.
func (h *MyHandlers) GetCustomCountParams(params operations.GetCustomCountParams, principal interface{}) middleware.Responder {

	var seed int64
	seedGiven := params.Seed != nil
	if seedGiven {
		seed = *params.Seed
	} else {
		seed = faker.Global.Int64()
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

// GenerateCustom will replace all known ~keys~ from the template with random data, based on the seed. The result will contain (count) such strings.
//For each unique seed/template/count group the result should be the same (deterministic).
func GenerateCustom(seed int64, template string, count int32) []string {
	//add random data
	fastFaker := faker.NewFastFaker()
	fastFaker.Seed(seed)
	result := make([]string, 0, count)
	for i := 0; i < int(count); i++ {
		res, err := fastFaker.TemplateCustom(template, "~", "~")
		if err != nil {
			log.Printf("faker failed: %s", err)
			continue
		}
		result = append(result, res)
	}
	return result
}
