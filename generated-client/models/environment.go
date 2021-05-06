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

// Environment Environment holds detail information about environment
//
// swagger:model Environment
type Environment struct {

	// BranchMapping The branch mapped to this environment
	// Example: master
	BranchMapping string `json:"branchMapping,omitempty"`

	// Deployments All deployments in environment
	Deployments []*DeploymentSummary `json:"deployments"`

	// Name of the environment
	// Example: prod
	Name string `json:"name,omitempty"`

	// Secrets All secrets in environment
	Secrets []*Secret `json:"secrets"`

	// Status of the environment
	// Pending = Environment exists in Radix config, but not in cluster
	// Consistent = Environment exists in Radix config and in cluster
	// Orphan = Environment does not exist in Radix config, but exists in cluster
	// Example: Consistent
	// Enum: [Pending Consistent Orphan]
	Status string `json:"status,omitempty"`

	// active deployment
	ActiveDeployment *Deployment `json:"activeDeployment,omitempty"`
}

// Validate validates this environment
func (m *Environment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeployments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecrets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActiveDeployment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Environment) validateDeployments(formats strfmt.Registry) error {
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

func (m *Environment) validateSecrets(formats strfmt.Registry) error {
	if swag.IsZero(m.Secrets) { // not required
		return nil
	}

	for i := 0; i < len(m.Secrets); i++ {
		if swag.IsZero(m.Secrets[i]) { // not required
			continue
		}

		if m.Secrets[i] != nil {
			if err := m.Secrets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("secrets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var environmentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Pending","Consistent","Orphan"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		environmentTypeStatusPropEnum = append(environmentTypeStatusPropEnum, v)
	}
}

const (

	// EnvironmentStatusPending captures enum value "Pending"
	EnvironmentStatusPending string = "Pending"

	// EnvironmentStatusConsistent captures enum value "Consistent"
	EnvironmentStatusConsistent string = "Consistent"

	// EnvironmentStatusOrphan captures enum value "Orphan"
	EnvironmentStatusOrphan string = "Orphan"
)

// prop value enum
func (m *Environment) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, environmentTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Environment) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Environment) validateActiveDeployment(formats strfmt.Registry) error {
	if swag.IsZero(m.ActiveDeployment) { // not required
		return nil
	}

	if m.ActiveDeployment != nil {
		if err := m.ActiveDeployment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activeDeployment")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this environment based on the context it is used
func (m *Environment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeployments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecrets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateActiveDeployment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Environment) contextValidateDeployments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Deployments); i++ {

		if m.Deployments[i] != nil {
			if err := m.Deployments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deployments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Environment) contextValidateSecrets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Secrets); i++ {

		if m.Secrets[i] != nil {
			if err := m.Secrets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("secrets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Environment) contextValidateActiveDeployment(ctx context.Context, formats strfmt.Registry) error {

	if m.ActiveDeployment != nil {
		if err := m.ActiveDeployment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activeDeployment")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Environment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Environment) UnmarshalBinary(b []byte) error {
	var res Environment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
