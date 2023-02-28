// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AzureIdentity AzureIdentity describes an identity in Azure
//
// swagger:model AzureIdentity
type AzureIdentity struct {

	// ClientId is the client ID of an Azure User Assigned Managed Identity
	// or the application ID of an Azure AD Application Registration
	// Required: true
	ClientID *string `json:"clientId"`

	// The Service Account name to use when configuring Kubernetes Federation Credentials for the identity
	// Required: true
	ServiceAccountName *string `json:"serviceAccountName"`
}

// Validate validates this azure identity
func (m *AzureIdentity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceAccountName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureIdentity) validateClientID(formats strfmt.Registry) error {

	if err := validate.Required("clientId", "body", m.ClientID); err != nil {
		return err
	}

	return nil
}

func (m *AzureIdentity) validateServiceAccountName(formats strfmt.Registry) error {

	if err := validate.Required("serviceAccountName", "body", m.ServiceAccountName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this azure identity based on context it is used
func (m *AzureIdentity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureIdentity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureIdentity) UnmarshalBinary(b []byte) error {
	var res AzureIdentity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}