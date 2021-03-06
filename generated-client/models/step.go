// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Step Step holds general information about job step
//
// swagger:model Step
type Step struct {

	// Components associated components
	Components []string `json:"components"`

	// Ended timestamp
	// Example: 2006-01-02T15:04:05Z
	Ended string `json:"ended,omitempty"`

	// Name of the step
	// Example: build
	Name string `json:"name,omitempty"`

	// Started timestamp
	// Example: 2006-01-02T15:04:05Z
	Started string `json:"started,omitempty"`

	// Status of the step
	// Example: Waiting
	// Enum: [Waiting Running Succeeded Failed]
	Status string `json:"status,omitempty"`
}

// Validate validates this step
func (m *Step) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var stepTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Waiting","Running","Succeeded","Failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stepTypeStatusPropEnum = append(stepTypeStatusPropEnum, v)
	}
}

const (

	// StepStatusWaiting captures enum value "Waiting"
	StepStatusWaiting string = "Waiting"

	// StepStatusRunning captures enum value "Running"
	StepStatusRunning string = "Running"

	// StepStatusSucceeded captures enum value "Succeeded"
	StepStatusSucceeded string = "Succeeded"

	// StepStatusFailed captures enum value "Failed"
	StepStatusFailed string = "Failed"
)

// prop value enum
func (m *Step) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, stepTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Step) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this step based on context it is used
func (m *Step) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Step) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Step) UnmarshalBinary(b []byte) error {
	var res Step
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
