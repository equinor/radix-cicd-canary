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

// PipelineRun PipelineRun holds general information about pipeline run
//
// swagger:model PipelineRun
type PipelineRun struct {

	// Ended timestamp
	// Example: 2006-01-02T15:04:05Z
	Ended string `json:"ended,omitempty"`

	// Env Environment of the pipeline run
	// Example: prod
	// Required: true
	Env *string `json:"env"`

	// Name Original name of the pipeline run
	// Example: build-pipeline
	// Required: true
	Name *string `json:"name"`

	// RealName Name of the pipeline run in the namespace
	// Example: radix-tekton-pipelinerun-dev-2022-05-09-abcde
	// Required: true
	RealName *string `json:"realName"`

	// Started timestamp
	// Example: 2006-01-02T15:04:05Z
	Started string `json:"started,omitempty"`

	// Status of the step
	// Example: Waiting
	// Enum: [Waiting Running Succeeded Failed]
	Status string `json:"status,omitempty"`

	// StatusMessage of the task
	StatusMessage string `json:"statusMessage,omitempty"`
}

// Validate validates this pipeline run
func (m *PipelineRun) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRealName(formats); err != nil {
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

func (m *PipelineRun) validateEnv(formats strfmt.Registry) error {

	if err := validate.Required("env", "body", m.Env); err != nil {
		return err
	}

	return nil
}

func (m *PipelineRun) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PipelineRun) validateRealName(formats strfmt.Registry) error {

	if err := validate.Required("realName", "body", m.RealName); err != nil {
		return err
	}

	return nil
}

var pipelineRunTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Waiting","Running","Succeeded","Failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pipelineRunTypeStatusPropEnum = append(pipelineRunTypeStatusPropEnum, v)
	}
}

const (

	// PipelineRunStatusWaiting captures enum value "Waiting"
	PipelineRunStatusWaiting string = "Waiting"

	// PipelineRunStatusRunning captures enum value "Running"
	PipelineRunStatusRunning string = "Running"

	// PipelineRunStatusSucceeded captures enum value "Succeeded"
	PipelineRunStatusSucceeded string = "Succeeded"

	// PipelineRunStatusFailed captures enum value "Failed"
	PipelineRunStatusFailed string = "Failed"
)

// prop value enum
func (m *PipelineRun) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pipelineRunTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PipelineRun) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this pipeline run based on context it is used
func (m *PipelineRun) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PipelineRun) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelineRun) UnmarshalBinary(b []byte) error {
	var res PipelineRun
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
