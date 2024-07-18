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

// ReplicaStatus ReplicaStatus describes the status of a component container inside a pod
//
// swagger:model ReplicaStatus
type ReplicaStatus struct {

	// Status of the container
	// Pending = Container in Waiting state and the reason is ContainerCreating
	// Failed = Container is failed
	// Failing = Container is failed
	// Running = Container in Running state
	// Succeeded = Container in Succeeded state
	// Terminated = Container in Terminated state
	// Stopped = Job has been stopped
	// Example: Running
	// Required: true
	// Enum: [Pending Succeeded Failing Failed Running Terminated Starting Stopped]
	Status *string `json:"status"`
}

// Validate validates this replica status
func (m *ReplicaStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var replicaStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Pending","Succeeded","Failing","Failed","Running","Terminated","Starting","Stopped"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		replicaStatusTypeStatusPropEnum = append(replicaStatusTypeStatusPropEnum, v)
	}
}

const (

	// ReplicaStatusStatusPending captures enum value "Pending"
	ReplicaStatusStatusPending string = "Pending"

	// ReplicaStatusStatusSucceeded captures enum value "Succeeded"
	ReplicaStatusStatusSucceeded string = "Succeeded"

	// ReplicaStatusStatusFailing captures enum value "Failing"
	ReplicaStatusStatusFailing string = "Failing"

	// ReplicaStatusStatusFailed captures enum value "Failed"
	ReplicaStatusStatusFailed string = "Failed"

	// ReplicaStatusStatusRunning captures enum value "Running"
	ReplicaStatusStatusRunning string = "Running"

	// ReplicaStatusStatusTerminated captures enum value "Terminated"
	ReplicaStatusStatusTerminated string = "Terminated"

	// ReplicaStatusStatusStarting captures enum value "Starting"
	ReplicaStatusStatusStarting string = "Starting"

	// ReplicaStatusStatusStopped captures enum value "Stopped"
	ReplicaStatusStatusStopped string = "Stopped"
)

// prop value enum
func (m *ReplicaStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, replicaStatusTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReplicaStatus) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this replica status based on context it is used
func (m *ReplicaStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReplicaStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplicaStatus) UnmarshalBinary(b []byte) error {
	var res ReplicaStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
