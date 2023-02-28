// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EnvVarMetadata EnvVarMetadata Environment variable metadata, holding state of creating or changing of value in Radix console
//
// swagger:model EnvVarMetadata
type EnvVarMetadata struct {

	// Value of the environment variable in radixconfig.yaml
	// Example: value1
	RadixConfigValue string `json:"radixConfigValue,omitempty"`
}

// Validate validates this env var metadata
func (m *EnvVarMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this env var metadata based on context it is used
func (m *EnvVarMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EnvVarMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvVarMetadata) UnmarshalBinary(b []byte) error {
	var res EnvVarMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}