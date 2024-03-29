// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Status Status is a return value for calls that don't return other objects or when a request returns an error
//
// swagger:model Status
type Status struct {

	// Suggested HTTP return code for this status, 0 if not set.
	// Example: 404
	Code int64 `json:"code,omitempty"`

	// A human-readable description of the status of this operation.
	// Example: job job123 is not found
	Message string `json:"message,omitempty"`

	// Status of the operation.
	// One of: "Success" or "Failure".
	// Example: Failure
	Status string `json:"status,omitempty"`

	// reason
	Reason StatusReason `json:"reason,omitempty"`
}

// Validate validates this status
func (m *Status) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Status) validateReason(formats strfmt.Registry) error {
	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	if err := m.Reason.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reason")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("reason")
		}
		return err
	}

	return nil
}

// ContextValidate validate this status based on the context it is used
func (m *Status) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Status) contextValidateReason(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	if err := m.Reason.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reason")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("reason")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Status) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Status) UnmarshalBinary(b []byte) error {
	var res Status
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
