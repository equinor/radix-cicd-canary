// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RadixNode RadixNode defines node attributes, where container should be scheduled
//
// swagger:model RadixNode
type RadixNode struct {

	// Defines rules for allowed GPU types.
	// More info: https://www.radix.equinor.com/references/reference-radix-config/#gpu
	// +optional
	Gpu string `json:"gpu,omitempty"`

	// Defines minimum number of required GPUs.
	// +optional
	GpuCount string `json:"gpuCount,omitempty"`
}

// Validate validates this radix node
func (m *RadixNode) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this radix node based on context it is used
func (m *RadixNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RadixNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RadixNode) UnmarshalBinary(b []byte) error {
	var res RadixNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
