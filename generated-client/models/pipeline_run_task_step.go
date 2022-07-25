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

// PipelineRunTaskStep PipelineRunTaskStep holds general information about pipeline run task steps
//
// swagger:model PipelineRunTaskStep
type PipelineRunTaskStep struct {

	// Ended timestamp
	// Example: 2006-01-02T15:04:05Z
	Ended string `json:"ended,omitempty"`

	// Name of the step
	// Example: build
	// Required: true
	Name *string `json:"name"`

	// Started timestamp
	// Example: 2006-01-02T15:04:05Z
	Started string `json:"started,omitempty"`

	// Status of the task
	// Example: Waiting
	// Enum: [Waiting Running Succeeded Failed]
	Status string `json:"status,omitempty"`

	// StatusMessage of the task
	StatusMessage string `json:"statusMessage,omitempty"`
}

// Validate validates this pipeline run task step
func (m *PipelineRunTaskStep) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PipelineRunTaskStep) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var pipelineRunTaskStepTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Waiting","Running","Succeeded","Failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pipelineRunTaskStepTypeStatusPropEnum = append(pipelineRunTaskStepTypeStatusPropEnum, v)
	}
}

const (

	// PipelineRunTaskStepStatusWaiting captures enum value "Waiting"
	PipelineRunTaskStepStatusWaiting string = "Waiting"

	// PipelineRunTaskStepStatusRunning captures enum value "Running"
	PipelineRunTaskStepStatusRunning string = "Running"

	// PipelineRunTaskStepStatusSucceeded captures enum value "Succeeded"
	PipelineRunTaskStepStatusSucceeded string = "Succeeded"

	// PipelineRunTaskStepStatusFailed captures enum value "Failed"
	PipelineRunTaskStepStatusFailed string = "Failed"
)

// prop value enum
func (m *PipelineRunTaskStep) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pipelineRunTaskStepTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PipelineRunTaskStep) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this pipeline run task step based on context it is used
func (m *PipelineRunTaskStep) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PipelineRunTaskStep) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelineRunTaskStep) UnmarshalBinary(b []byte) error {
	var res PipelineRunTaskStep
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}