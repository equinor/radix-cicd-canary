// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// JobSummary JobSummary holds general information about job
// swagger:model JobSummary
type JobSummary struct {

	// AppName of the application
	AppName string `json:"appName,omitempty"`

	// Branch branch to build from
	Branch string `json:"branch,omitempty"`

	// CommitID the commit ID of the branch to build
	CommitID string `json:"commitID,omitempty"`

	// Created timestamp
	Created string `json:"created,omitempty"`

	// Ended timestamp
	Ended string `json:"ended,omitempty"`

	// Environments the job deployed to
	Environments []string `json:"environments"`

	// Name of the job
	Name string `json:"name,omitempty"`

	// Name of the pipeline
	// Enum: [build-deploy  build]
	Pipeline string `json:"pipeline,omitempty"`

	// Started timestamp
	Started string `json:"started,omitempty"`

	// Status of the job
	// Enum: [Waiting Running Succeeded Stopping Stopped Failed]
	Status string `json:"status,omitempty"`

	// TriggeredBy user that triggered the job. If through webhook = sender.login. If through api - usertoken.upn
	TriggeredBy string `json:"triggeredBy,omitempty"`
}

// Validate validates this job summary
func (m *JobSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePipeline(formats); err != nil {
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

var jobSummaryTypePipelinePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["build-deploy"," build"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jobSummaryTypePipelinePropEnum = append(jobSummaryTypePipelinePropEnum, v)
	}
}

const (

	// JobSummaryPipelineBuildDeploy captures enum value "build-deploy"
	JobSummaryPipelineBuildDeploy string = "build-deploy"

	// JobSummaryPipelineBuild captures enum value " build"
	JobSummaryPipelineBuild string = " build"
)

// prop value enum
func (m *JobSummary) validatePipelineEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, jobSummaryTypePipelinePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *JobSummary) validatePipeline(formats strfmt.Registry) error {

	if swag.IsZero(m.Pipeline) { // not required
		return nil
	}

	// value enum
	if err := m.validatePipelineEnum("pipeline", "body", m.Pipeline); err != nil {
		return err
	}

	return nil
}

var jobSummaryTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Waiting","Running","Succeeded","Stopping","Stopped","Failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jobSummaryTypeStatusPropEnum = append(jobSummaryTypeStatusPropEnum, v)
	}
}

const (

	// JobSummaryStatusWaiting captures enum value "Waiting"
	JobSummaryStatusWaiting string = "Waiting"

	// JobSummaryStatusRunning captures enum value "Running"
	JobSummaryStatusRunning string = "Running"

	// JobSummaryStatusSucceeded captures enum value "Succeeded"
	JobSummaryStatusSucceeded string = "Succeeded"

	// JobSummaryStatusStopping captures enum value "Stopping"
	JobSummaryStatusStopping string = "Stopping"

	// JobSummaryStatusStopped captures enum value "Stopped"
	JobSummaryStatusStopped string = "Stopped"

	// JobSummaryStatusFailed captures enum value "Failed"
	JobSummaryStatusFailed string = "Failed"
)

// prop value enum
func (m *JobSummary) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, jobSummaryTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *JobSummary) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobSummary) UnmarshalBinary(b []byte) error {
	var res JobSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
