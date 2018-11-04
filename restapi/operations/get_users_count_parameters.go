// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetUsersCountParams creates a new GetUsersCountParams object
// no default values defined in spec.
func NewGetUsersCountParams() GetUsersCountParams {

	return GetUsersCountParams{}
}

// GetUsersCountParams contains all the bound params for the get users count operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetUsersCount
type GetUsersCountParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*How many users to return
	  Required: true
	  Maximum: 500
	  Minimum: 1
	  In: path
	*/
	Count int32
	/*Seed that will be used to generate the users (deterministic call). The seed will determine the first User data, seed+1 the next user, and so on. If no seed is provided a pseudo-random one will be generated (rand.Int63).
	  In: query
	*/
	Seed *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetUsersCountParams() beforehand.
func (o *GetUsersCountParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rCount, rhkCount, _ := route.Params.GetOK("count")
	if err := o.bindCount(rCount, rhkCount, route.Formats); err != nil {
		res = append(res, err)
	}

	qSeed, qhkSeed, _ := qs.GetOK("seed")
	if err := o.bindSeed(qSeed, qhkSeed, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCount binds and validates parameter Count from path.
func (o *GetUsersCountParams) bindCount(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("count", "path", "int32", raw)
	}
	o.Count = value

	if err := o.validateCount(formats); err != nil {
		return err
	}

	return nil
}

// validateCount carries on validations for parameter Count
func (o *GetUsersCountParams) validateCount(formats strfmt.Registry) error {

	if err := validate.MinimumInt("count", "path", int64(o.Count), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("count", "path", int64(o.Count), 500, false); err != nil {
		return err
	}

	return nil
}

// bindSeed binds and validates parameter Seed from query.
func (o *GetUsersCountParams) bindSeed(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("seed", "query", "int64", raw)
	}
	o.Seed = &value

	return nil
}