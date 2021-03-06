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

// StepLog StepLog holds logs for a given step
//
// swagger:model StepLog
type StepLog struct {

	// Log of step
	Log string `json:"log,omitempty"`

	// Name of the step
	// Required: true
	Name *string `json:"name"`

	// Name of the logged pod
	PodName string `json:"podname,omitempty"`

	// sort
	Sort int32 `json:"sort,omitempty"`
}

// Validate validates this step log
func (m *StepLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StepLog) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this step log based on context it is used
func (m *StepLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StepLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StepLog) UnmarshalBinary(b []byte) error {
	var res StepLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
