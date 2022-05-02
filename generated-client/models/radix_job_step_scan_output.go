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

// RadixJobStepScanOutput RadixJobStepScanOutput holds information about output from a single scan step
//
// swagger:model RadixJobStepScanOutput
type RadixJobStepScanOutput struct {

	// Reason for the status
	Reason string `json:"reason,omitempty"`

	// VulnerabilityListConfigMap defines the name of the ConfigMap with list of information about vulnerabilities found during scan
	// the ConfigMap must be in the same namespace as the RadixJob
	VulnerabilityListConfigMap string `json:"vulnerabilityListConfigMap,omitempty"`

	// VulnerabilityListKey defines the key in VulnerabilityListConfigMap where vulnerability details are stored
	VulnerabilityListKey string `json:"vulnerabilityListKey,omitempty"`

	// status
	Status ScanStatus `json:"status,omitempty"`

	// vulnerabilities
	Vulnerabilities VulnerabilityMap `json:"vulnerabilities,omitempty"`
}

// Validate validates this radix job step scan output
func (m *RadixJobStepScanOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVulnerabilities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RadixJobStepScanOutput) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *RadixJobStepScanOutput) validateVulnerabilities(formats strfmt.Registry) error {
	if swag.IsZero(m.Vulnerabilities) { // not required
		return nil
	}

	if m.Vulnerabilities != nil {
		if err := m.Vulnerabilities.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vulnerabilities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vulnerabilities")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this radix job step scan output based on the context it is used
func (m *RadixJobStepScanOutput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVulnerabilities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RadixJobStepScanOutput) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *RadixJobStepScanOutput) contextValidateVulnerabilities(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Vulnerabilities.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("vulnerabilities")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("vulnerabilities")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RadixJobStepScanOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RadixJobStepScanOutput) UnmarshalBinary(b []byte) error {
	var res RadixJobStepScanOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}