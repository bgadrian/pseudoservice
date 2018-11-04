package handlers

import "github.com/bgadrian/pseudoservice/restapi/operations"

// MyHandlers we need this struct to inject dependencies and access them from handlers
//like logger.
type MyHandlers struct {
	Api *operations.PseudoServiceAPI
}
