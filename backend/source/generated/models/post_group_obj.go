// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostGroupObj post group obj
//
// swagger:model postGroupObj
type PostGroupObj struct {

	// group's desciption
	Description string `json:"description,omitempty"`

	// group name
	GroupName string `json:"group_name,omitempty"`

	// group owner
	GroupOwner string `json:"group_owner,omitempty"`
}

// Validate validates this post group obj
func (m *PostGroupObj) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post group obj based on context it is used
func (m *PostGroupObj) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostGroupObj) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostGroupObj) UnmarshalBinary(b []byte) error {
	var res PostGroupObj
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
