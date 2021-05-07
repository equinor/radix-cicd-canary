// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RegenerateDeployKeyAndSecretData RegenerateDeployKeyAndSecretData Holds regenerated shared secret
//
// swagger:model RegenerateDeployKeyAndSecretData
type RegenerateDeployKeyAndSecretData struct {

	// SharedSecret of the shared secret
	// Required: true
	SharedSecret *string `json:"sharedSecret"`
}

// Validate validates this regenerate deploy key and secret data
func (m *RegenerateDeployKeyAndSecretData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSharedSecret(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegenerateDeployKeyAndSecretData) validateSharedSecret(formats strfmt.Registry) error {

	if err := validate.Required("sharedSecret", "body", m.SharedSecret); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this regenerate deploy key and secret data based on context it is used
func (m *RegenerateDeployKeyAndSecretData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RegenerateDeployKeyAndSecretData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegenerateDeployKeyAndSecretData) UnmarshalBinary(b []byte) error {
	var res RegenerateDeployKeyAndSecretData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
