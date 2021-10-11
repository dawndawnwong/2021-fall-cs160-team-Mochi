// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GroupObj group obj
//
// swagger:model groupObj
type GroupObj struct {

	// group's desciption
	Description string `json:"description,omitempty"`

	// group id
	GroupID string `json:"group_id,omitempty"`

	// group name
	GroupName string `json:"group_name,omitempty"`

	// group owner
	GroupOwner string `json:"group_owner,omitempty"`
}

// Validate validates this group obj
func (m *GroupObj) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this group obj based on context it is used
func (m *GroupObj) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GroupObj) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupObj) UnmarshalBinary(b []byte) error {
	var res GroupObj
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
