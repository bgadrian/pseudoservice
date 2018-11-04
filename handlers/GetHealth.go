package handlers

import (
	"github.com/bgadrian/pseudoservice/restapi/operations"
	"github.com/brianvoe/gofakeit"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
	"time"
)

func (*MyHandlers) GetHealthHandler(params operations.GetHealthParams) middleware.Responder {
	//gofakeit can remain out of entropy and block forever,
	timeout := time.NewTicker(time.Millisecond * 200)
	var done chan bool
	go func() {
		gofakeit.Int64()
		done <- true
	}()
	select {
	case <-timeout.C:
		return operations.NewGetHealthDefault(http.StatusInternalServerError)
	case <-done:
		return operations.NewGetHealthOK()
	}
}
