// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// User user
// swagger:model User
type User struct {

	// company
	Company string `json:"Company,omitempty"`

	// age
	// Maximum: 255
	// Minimum: 0
	Age *int32 `json:"age,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// Random list of UUIDs from the same response. To increase the chance of having friends request for bigger batches.
	Friends []string `json:"friends"`

	// UUID for this user
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// First and Last name
	// Required: true
	Name *string `json:"name"`

	// position
	Position string `json:"position,omitempty"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateAge(formats strfmt.Registry) error {

	if swag.IsZero(m.Age) { // not required
		return nil
	}

	if err := validate.Minimum("age", "body", float64(*m.Age), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("age", "body", float64(*m.Age), 255, false); err != nil {
		return err
	}

	return nil
}

func (m *User) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *User) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
