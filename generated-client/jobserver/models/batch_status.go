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

// BatchStatus BatchStatus holds general information about batch status
//
// swagger:model BatchStatus
type BatchStatus struct {

	// Defines a user defined ID of the batch.
	// Example: 'batch-id-1'
	BatchID string `json:"batchId,omitempty"`

	// BatchName Optional Batch ID of a job
	// Example: 'batch1'
	BatchName string `json:"batchName,omitempty"`

	// BatchType Single job or multiple jobs batch
	// Example: \"job\
	BatchType string `json:"batchType,omitempty"`

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

	// JobStatuses of the jobs in the batch
	JobStatuses []*JobStatus `json:"jobStatuses"`

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

// Validate validates this batch status
func (m *BatchStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobStatuses(formats); err != nil {
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

func (m *BatchStatus) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BatchStatus) validateEnded(formats strfmt.Registry) error {
	if swag.IsZero(m.Ended) { // not required
		return nil
	}

	if err := validate.FormatOf("ended", "body", "date-time", m.Ended.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BatchStatus) validateJobStatuses(formats strfmt.Registry) error {
	if swag.IsZero(m.JobStatuses) { // not required
		return nil
	}

	for i := 0; i < len(m.JobStatuses); i++ {
		if swag.IsZero(m.JobStatuses[i]) { // not required
			continue
		}

		if m.JobStatuses[i] != nil {
			if err := m.JobStatuses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("jobStatuses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("jobStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BatchStatus) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *BatchStatus) validatePodStatuses(formats strfmt.Registry) error {
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

func (m *BatchStatus) validateStarted(formats strfmt.Registry) error {
	if swag.IsZero(m.Started) { // not required
		return nil
	}

	if err := validate.FormatOf("started", "body", "date-time", m.Started.String(), formats); err != nil {
		return err
	}

	return nil
}

var batchStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Running","Succeeded","Failed","Waiting","Stopping","Stopped","Active","Completed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		batchStatusTypeStatusPropEnum = append(batchStatusTypeStatusPropEnum, v)
	}
}

const (

	// BatchStatusStatusRunning captures enum value "Running"
	BatchStatusStatusRunning string = "Running"

	// BatchStatusStatusSucceeded captures enum value "Succeeded"
	BatchStatusStatusSucceeded string = "Succeeded"

	// BatchStatusStatusFailed captures enum value "Failed"
	BatchStatusStatusFailed string = "Failed"

	// BatchStatusStatusWaiting captures enum value "Waiting"
	BatchStatusStatusWaiting string = "Waiting"

	// BatchStatusStatusStopping captures enum value "Stopping"
	BatchStatusStatusStopping string = "Stopping"

	// BatchStatusStatusStopped captures enum value "Stopped"
	BatchStatusStatusStopped string = "Stopped"

	// BatchStatusStatusActive captures enum value "Active"
	BatchStatusStatusActive string = "Active"

	// BatchStatusStatusCompleted captures enum value "Completed"
	BatchStatusStatusCompleted string = "Completed"
)

// prop value enum
func (m *BatchStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, batchStatusTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BatchStatus) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *BatchStatus) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this batch status based on the context it is used
func (m *BatchStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateJobStatuses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePodStatuses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchStatus) contextValidateJobStatuses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.JobStatuses); i++ {

		if m.JobStatuses[i] != nil {

			if swag.IsZero(m.JobStatuses[i]) { // not required
				return nil
			}

			if err := m.JobStatuses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("jobStatuses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("jobStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BatchStatus) contextValidatePodStatuses(ctx context.Context, formats strfmt.Registry) error {

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
func (m *BatchStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BatchStatus) UnmarshalBinary(b []byte) error {
	var res BatchStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
