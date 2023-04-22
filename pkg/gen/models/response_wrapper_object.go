// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResponseWrapperObject ResponseWrapper«object»
//
// swagger:model ResponseWrapper«object»
type ResponseWrapperObject struct {

	// Code 1：Success
	Code int32 `json:"code,omitempty"`

	// Response message
	Message string `json:"message,omitempty"`

	// Response result
	Result interface{} `json:"result,omitempty"`
}

// Validate validates this response wrapper object
func (m *ResponseWrapperObject) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponseWrapperObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseWrapperObject) UnmarshalBinary(b []byte) error {
	var res ResponseWrapperObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
