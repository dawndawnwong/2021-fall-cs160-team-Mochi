// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetFileResponse get file response
//
// swagger:model getFileResponse
type GetFileResponse struct {

	// note id for pdf file
	NoteID string `json:"note_id,omitempty"`

	// note reference for pdf file
	NoteReference string `json:"note_reference,omitempty"`

	// pdf content
	PdfData interface{} `json:"pdf_data,omitempty"`
}

// Validate validates this get file response
func (m *GetFileResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get file response based on context it is used
func (m *GetFileResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetFileResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetFileResponse) UnmarshalBinary(b []byte) error {
	var res GetFileResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
