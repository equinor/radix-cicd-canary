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

// ApplicationSummary ApplicationSummary describe an application
//
// swagger:model ApplicationSummary
type ApplicationSummary struct {

	// Name the name of the application
	// Example: radix-canary-golang
	Name string `json:"name,omitempty"`

	// latest job
	LatestJob *JobSummary `json:"latestJob,omitempty"`
}

// Validate validates this application summary
func (m *ApplicationSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLatestJob(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSummary) validateLatestJob(formats strfmt.Registry) error {
	if swag.IsZero(m.LatestJob) { // not required
		return nil
	}

	if m.LatestJob != nil {
		if err := m.LatestJob.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latestJob")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latestJob")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application summary based on the context it is used
func (m *ApplicationSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLatestJob(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSummary) contextValidateLatestJob(ctx context.Context, formats strfmt.Registry) error {

	if m.LatestJob != nil {
		if err := m.LatestJob.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latestJob")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latestJob")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationSummary) UnmarshalBinary(b []byte) error {
	var res ApplicationSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
