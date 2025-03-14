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

// ComponentSummary ComponentSummary describe a component part of a deployment
//
// swagger:model ComponentSummary
type ComponentSummary struct {

	// CommitID the commit ID of the branch to build
	// REQUIRED for "build" and "build-deploy" pipelines
	// Example: 4faca8595c5283a9d0f17a623b9255a0d9866a2e
	CommitID string `json:"commitID,omitempty"`

	// GitTags the git tags that the git commit hash points to
	// Example: \"v1.22.1 v1.22.3\
	GitTags string `json:"gitTags,omitempty"`

	// Image name
	// Example: radixdev.azurecr.io/app-server:cdgkg
	// Required: true
	Image *string `json:"image"`

	// Name the component
	// Example: server
	// Required: true
	Name *string `json:"name"`

	// SkipDeployment The component should not be deployed, but used existing
	// Example: true
	SkipDeployment bool `json:"skipDeployment,omitempty"`

	// Type of component
	// Example: component
	// Required: true
	// Enum: ["component","job"]
	Type *string `json:"type"`

	// resources
	Resources *ResourceRequirements `json:"resources,omitempty"`

	// runtime
	Runtime *Runtime `json:"runtime,omitempty"`
}

// Validate validates this component summary
func (m *ComponentSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuntime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComponentSummary) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	return nil
}

func (m *ComponentSummary) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var componentSummaryTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["component","job"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		componentSummaryTypeTypePropEnum = append(componentSummaryTypeTypePropEnum, v)
	}
}

const (

	// ComponentSummaryTypeComponent captures enum value "component"
	ComponentSummaryTypeComponent string = "component"

	// ComponentSummaryTypeJob captures enum value "job"
	ComponentSummaryTypeJob string = "job"
)

// prop value enum
func (m *ComponentSummary) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, componentSummaryTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ComponentSummary) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ComponentSummary) validateResources(formats strfmt.Registry) error {
	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	if m.Resources != nil {
		if err := m.Resources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

func (m *ComponentSummary) validateRuntime(formats strfmt.Registry) error {
	if swag.IsZero(m.Runtime) { // not required
		return nil
	}

	if m.Runtime != nil {
		if err := m.Runtime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runtime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runtime")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this component summary based on the context it is used
func (m *ComponentSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRuntime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComponentSummary) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

	if m.Resources != nil {

		if swag.IsZero(m.Resources) { // not required
			return nil
		}

		if err := m.Resources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

func (m *ComponentSummary) contextValidateRuntime(ctx context.Context, formats strfmt.Registry) error {

	if m.Runtime != nil {

		if swag.IsZero(m.Runtime) { // not required
			return nil
		}

		if err := m.Runtime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("runtime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("runtime")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComponentSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComponentSummary) UnmarshalBinary(b []byte) error {
	var res ComponentSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
