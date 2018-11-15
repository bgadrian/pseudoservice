package handlers

import (
	"net/http"
	"time"

	"github.com/bgadrian/fastfaker"
	"github.com/bgadrian/pseudoservice/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

// GetHealthHandler basic /health checks for math.rand entropy
func (*MyHandlers) GetHealthHandler(params operations.GetHealthParams, principal interface{}) middleware.Responder {
	//fastfaker can remain out of entropy and block forever,
	timeout := time.NewTicker(time.Millisecond * 200)
	var done chan bool
	go func() {
		fastfaker.Global.Int64()
		done <- true
	}()
	select {
	case <-timeout.C:
		return operations.NewGetHealthDefault(http.StatusInternalServerError)
	case <-done:
		return operations.NewGetHealthOK()
	}
}
