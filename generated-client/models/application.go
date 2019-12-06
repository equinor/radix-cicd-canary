// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Application Application details of an application
// swagger:model Application
type Application struct {

	// Creator of the application (user principle name).
	Creator string `json:"creator,omitempty"`

	// Environments List of environments for this application
	Environments []*EnvironmentSummary `json:"environments"`

	// Jobs list of run jobs for the application
	Jobs []*JobSummary `json:"jobs"`

	// Name the name of the application
	Name string `json:"name,omitempty"`

	// Owner of the application (email). Can be a single person or a shared group email
	Owner string `json:"owner,omitempty"`

	// app alias
	AppAlias *ApplicationAlias `json:"appAlias,omitempty"`

	// registration
	Registration *ApplicationRegistration `json:"registration,omitempty"`
}

// Validate validates this application
func (m *Application) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppAlias(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Application) validateEnvironments(formats strfmt.Registry) error {

	if swag.IsZero(m.Environments) { // not required
		return nil
	}

	for i := 0; i < len(m.Environments); i++ {
		if swag.IsZero(m.Environments[i]) { // not required
			continue
		}

		if m.Environments[i] != nil {
			if err := m.Environments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("environments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Application) validateJobs(formats strfmt.Registry) error {

	if swag.IsZero(m.Jobs) { // not required
		return nil
	}

	for i := 0; i < len(m.Jobs); i++ {
		if swag.IsZero(m.Jobs[i]) { // not required
			continue
		}

		if m.Jobs[i] != nil {
			if err := m.Jobs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("jobs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Application) validateAppAlias(formats strfmt.Registry) error {

	if swag.IsZero(m.AppAlias) { // not required
		return nil
	}

	if m.AppAlias != nil {
		if err := m.AppAlias.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appAlias")
			}
			return err
		}
	}

	return nil
}

func (m *Application) validateRegistration(formats strfmt.Registry) error {

	if swag.IsZero(m.Registration) { // not required
		return nil
	}

	if m.Registration != nil {
		if err := m.Registration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Application) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Application) UnmarshalBinary(b []byte) error {
	var res Application
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
