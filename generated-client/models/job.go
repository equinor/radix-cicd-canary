// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Job Job holds general information about job
// swagger:model Job
type Job struct {

	// Branch branch to build from
	Branch string `json:"branch,omitempty"`

	// CommitID the commit ID of the branch to build
	CommitID string `json:"commitID,omitempty"`

	// Components (array of ComponentSummary) created by the job
	Components []*ComponentSummary `json:"components"`

	// Created timestamp
	Created string `json:"created,omitempty"`

	// Array of deployments
	Deployments []*DeploymentSummary `json:"deployments"`

	// Ended timestamp
	Ended string `json:"ended,omitempty"`

	// Name of the job
	Name string `json:"name,omitempty"`

	// Name of the pipeline
	// Enum: [build-deploy]
	Pipeline string `json:"pipeline,omitempty"`

	// Started timestamp
	Started string `json:"started,omitempty"`

	// Status of the job
	// Enum: [Waiting Running Succeeded Failed]
	Status string `json:"status,omitempty"`

	// Array of steps
	Steps []*Step `json:"steps"`
}

// Validate validates this job
func (m *Job) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeployments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePipeline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSteps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Job) validateComponents(formats strfmt.Registry) error {

	if swag.IsZero(m.Components) { // not required
		return nil
	}

	for i := 0; i < len(m.Components); i++ {
		if swag.IsZero(m.Components[i]) { // not required
			continue
		}

		if m.Components[i] != nil {
			if err := m.Components[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Job) validateDeployments(formats strfmt.Registry) error {

	if swag.IsZero(m.Deployments) { // not required
		return nil
	}

	for i := 0; i < len(m.Deployments); i++ {
		if swag.IsZero(m.Deployments[i]) { // not required
			continue
		}

		if m.Deployments[i] != nil {
			if err := m.Deployments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deployments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var jobTypePipelinePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["build-deploy"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jobTypePipelinePropEnum = append(jobTypePipelinePropEnum, v)
	}
}

const (

	// JobPipelineBuildDeploy captures enum value "build-deploy"
	JobPipelineBuildDeploy string = "build-deploy"
)

// prop value enum
func (m *Job) validatePipelineEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, jobTypePipelinePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Job) validatePipeline(formats strfmt.Registry) error {

	if swag.IsZero(m.Pipeline) { // not required
		return nil
	}

	// value enum
	if err := m.validatePipelineEnum("pipeline", "body", m.Pipeline); err != nil {
		return err
	}

	return nil
}

var jobTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Waiting","Running","Succeeded","Failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jobTypeStatusPropEnum = append(jobTypeStatusPropEnum, v)
	}
}

const (

	// JobStatusWaiting captures enum value "Waiting"
	JobStatusWaiting string = "Waiting"

	// JobStatusRunning captures enum value "Running"
	JobStatusRunning string = "Running"

	// JobStatusSucceeded captures enum value "Succeeded"
	JobStatusSucceeded string = "Succeeded"

	// JobStatusFailed captures enum value "Failed"
	JobStatusFailed string = "Failed"
)

// prop value enum
func (m *Job) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, jobTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Job) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Job) validateSteps(formats strfmt.Registry) error {

	if swag.IsZero(m.Steps) { // not required
		return nil
	}

	for i := 0; i < len(m.Steps); i++ {
		if swag.IsZero(m.Steps[i]) { // not required
			continue
		}

		if m.Steps[i] != nil {
			if err := m.Steps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Job) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Job) UnmarshalBinary(b []byte) error {
	var res Job
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
