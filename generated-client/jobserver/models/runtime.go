// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Runtime Runtime defines the component or job's target runtime requirements
//
// swagger:model Runtime
type Runtime struct {

	// CPU architecture
	// Example: amd64
	Architecture string `json:"architecture,omitempty"`

	// Defines the node type for the component. It is a node-pool label and taint, where the component's or job's pods will be scheduled.
	// More info: https://www.radix.equinor.com/radix-config#nodetype
	// +kubebuilder:validation:MaxLength=120
	// +kubebuilder:validation:Pattern=^(([a-z0-9][-a-z0-9]*)?[a-z0-9])?$
	// +optional
	NodeType string `json:"nodeType,omitempty"`
}

// Validate validates this runtime
func (m *Runtime) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this runtime based on context it is used
func (m *Runtime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Runtime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Runtime) UnmarshalBinary(b []byte) error {
	var res Runtime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
