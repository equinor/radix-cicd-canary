// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuxiliaryResource AuxiliaryResource describes an auxiliary resources for a component
//
// swagger:model AuxiliaryResource
type AuxiliaryResource struct {

	// oauth2
	Oauth2 *OAuth2AuxiliaryResource `json:"oauth2,omitempty"`
}

// Validate validates this auxiliary resource
func (m *AuxiliaryResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOauth2(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuxiliaryResource) validateOauth2(formats strfmt.Registry) error {

	if swag.IsZero(m.Oauth2) { // not required
		return nil
	}

	if m.Oauth2 != nil {
		if err := m.Oauth2.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oauth2")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this auxiliary resource based on the context it is used
func (m *AuxiliaryResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOauth2(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuxiliaryResource) contextValidateOauth2(ctx context.Context, formats strfmt.Registry) error {

	if m.Oauth2 != nil {
		if err := m.Oauth2.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oauth2")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuxiliaryResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuxiliaryResource) UnmarshalBinary(b []byte) error {
	var res AuxiliaryResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
