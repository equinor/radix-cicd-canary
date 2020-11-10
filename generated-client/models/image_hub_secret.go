// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ImageHubSecret ImageHubSecret holds general information about image hubs
//
// swagger:model ImageHubSecret
type ImageHubSecret struct {

	// Email provided in radixconfig.yaml
	Email string `json:"email,omitempty"`

	// Server name of the image hub
	// Required: true
	Server *string `json:"server"`

	// Status of the secret
	// Pending = Secret value is not set
	// Consistent = Secret value is set
	Status string `json:"status,omitempty"`

	// Username for connecting to private image hub
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this image hub secret
func (m *ImageHubSecret) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageHubSecret) validateServer(formats strfmt.Registry) error {

	if err := validate.Required("server", "body", m.Server); err != nil {
		return err
	}

	return nil
}

func (m *ImageHubSecret) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageHubSecret) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageHubSecret) UnmarshalBinary(b []byte) error {
	var res ImageHubSecret
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
