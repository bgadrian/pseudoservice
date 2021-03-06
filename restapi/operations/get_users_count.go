// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetUsersCountHandlerFunc turns a function with the right signature into a get users count handler
type GetUsersCountHandlerFunc func(GetUsersCountParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUsersCountHandlerFunc) Handle(params GetUsersCountParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetUsersCountHandler interface for that can handle valid get users count params
type GetUsersCountHandler interface {
	Handle(GetUsersCountParams, interface{}) middleware.Responder
}

// NewGetUsersCount creates a new http.Handler for the get users count operation
func NewGetUsersCount(ctx *middleware.Context, handler GetUsersCountHandler) *GetUsersCount {
	return &GetUsersCount{Context: ctx, Handler: handler}
}

/*GetUsersCount swagger:route GET /users/{count} getUsersCount

Get a random user

*/
type GetUsersCount struct {
	Context *middleware.Context
	Handler GetUsersCountHandler
}

func (o *GetUsersCount) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUsersCountParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
