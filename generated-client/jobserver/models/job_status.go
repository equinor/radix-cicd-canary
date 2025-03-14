// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// JobStatus JobStatus holds general information about job status
//
// swagger:model JobStatus
type JobStatus struct {

	// Defines a user defined ID of the batch.
	// Example: 'batch-id-1'
	BatchID string `json:"batchId,omitempty"`

	// BatchName Optional Batch ID of a job
	// Example: 'batch1'
	BatchName string `json:"batchName,omitempty"`

	// Created timestamp
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// DeploymentName for this batch
	DeploymentName string `json:"DeploymentName,omitempty"`

	// Ended timestamp
	// Format: date-time
	Ended strfmt.DateTime `json:"ended,omitempty"`

	// The number of times the container for the job has failed.
	// +optional
	Failed int32 `json:"failed,omitempty"`

	// JobId Optional ID of a job
	// Example: 'job1'
	JobID string `json:"jobId,omitempty"`

	// Message, if any, of the job
	// Example: \"Error occurred\
	Message string `json:"message,omitempty"`

	// Name of the job
	// Example: calculator
	// Required: true
	Name *string `json:"name"`

	// PodStatuses for each pod of the job
	PodStatuses []*PodStatus `json:"podStatuses"`

	// Timestamp of the job restart, if applied.
	// +optional
	Restart string `json:"restart,omitempty"`

	// Started timestamp
	// Format: date-time
	Started strfmt.DateTime `json:"started,omitempty"`

	// Status of the job
	// Running = Job is running
	// Succeeded = Job has succeeded
	// Failed = Job has failed
	// Waiting = Job is waiting
	// Stopping = Job is stopping
	// Stopped = Job has been stopped
	// Active = Job is active
	// Completed = Job is completed
	// Example: Waiting
	// Enum: ["Running","Succeeded","Failed","Waiting","Stopping","Stopped","Active","Completed"]
	Status string `json:"status,omitempty"`

	// Updated timestamp when the status was updated
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this job status
func (m *JobStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodStatuses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStarted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobStatus) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JobStatus) validateEnded(formats strfmt.Registry) error {
	if swag.IsZero(m.Ended) { // not required
		return nil
	}

	if err := validate.FormatOf("ended", "body", "date-time", m.Ended.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JobStatus) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *JobStatus) validatePodStatuses(formats strfmt.Registry) error {
	if swag.IsZero(m.PodStatuses) { // not required
		return nil
	}

	for i := 0; i < len(m.PodStatuses); i++ {
		if swag.IsZero(m.PodStatuses[i]) { // not required
			continue
		}

		if m.PodStatuses[i] != nil {
			if err := m.PodStatuses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("podStatuses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("podStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *JobStatus) validateStarted(formats strfmt.Registry) error {
	if swag.IsZero(m.Started) { // not required
		return nil
	}

	if err := validate.FormatOf("started", "body", "date-time", m.Started.String(), formats); err != nil {
		return err
	}

	return nil
}

var jobStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Running","Succeeded","Failed","Waiting","Stopping","Stopped","Active","Completed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jobStatusTypeStatusPropEnum = append(jobStatusTypeStatusPropEnum, v)
	}
}

const (

	// JobStatusStatusRunning captures enum value "Running"
	JobStatusStatusRunning string = "Running"

	// JobStatusStatusSucceeded captures enum value "Succeeded"
	JobStatusStatusSucceeded string = "Succeeded"

	// JobStatusStatusFailed captures enum value "Failed"
	JobStatusStatusFailed string = "Failed"

	// JobStatusStatusWaiting captures enum value "Waiting"
	JobStatusStatusWaiting string = "Waiting"

	// JobStatusStatusStopping captures enum value "Stopping"
	JobStatusStatusStopping string = "Stopping"

	// JobStatusStatusStopped captures enum value "Stopped"
	JobStatusStatusStopped string = "Stopped"

	// JobStatusStatusActive captures enum value "Active"
	JobStatusStatusActive string = "Active"

	// JobStatusStatusCompleted captures enum value "Completed"
	JobStatusStatusCompleted string = "Completed"
)

// prop value enum
func (m *JobStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, jobStatusTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *JobStatus) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *JobStatus) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this job status based on the context it is used
func (m *JobStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePodStatuses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobStatus) contextValidatePodStatuses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PodStatuses); i++ {

		if m.PodStatuses[i] != nil {

			if swag.IsZero(m.PodStatuses[i]) { // not required
				return nil
			}

			if err := m.PodStatuses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("podStatuses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("podStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobStatus) UnmarshalBinary(b []byte) error {
	var res JobStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
