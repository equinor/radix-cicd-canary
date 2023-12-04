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

// TLSCertificate TLSCertificate holds information about a TLS certificate
//
// swagger:model TLSCertificate
type TLSCertificate struct {

	// DNSNames defines list of Subject Alternate Names in the certificate
	DNSNames []string `json:"dnsNames"`

	// Issuer contains the distinguished name for the certificate's issuer
	// Example: CN=DigiCert TLS RSA SHA256 2020 CA1,O=DigiCert Inc,C=US
	// Required: true
	Issuer *string `json:"issuer"`

	// NotAfter defines the uppdater date/time validity boundary
	// Required: true
	// Format: date
	NotAfter *strfmt.Date `json:"notAfter"`

	// NotBefore defines the lower date/time validity boundary
	// Required: true
	// Format: date
	NotBefore *strfmt.Date `json:"notBefore"`

	// Subject contains the distinguished name for the certificate
	// Example: CN=mysite.example.com,O=MyOrg,L=MyLocation,C=NO
	// Required: true
	Subject *string `json:"subject"`
}

// Validate validates this TLS certificate
func (m *TLSCertificate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIssuer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotAfter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotBefore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TLSCertificate) validateIssuer(formats strfmt.Registry) error {

	if err := validate.Required("issuer", "body", m.Issuer); err != nil {
		return err
	}

	return nil
}

func (m *TLSCertificate) validateNotAfter(formats strfmt.Registry) error {

	if err := validate.Required("notAfter", "body", m.NotAfter); err != nil {
		return err
	}

	if err := validate.FormatOf("notAfter", "body", "date", m.NotAfter.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TLSCertificate) validateNotBefore(formats strfmt.Registry) error {

	if err := validate.Required("notBefore", "body", m.NotBefore); err != nil {
		return err
	}

	if err := validate.FormatOf("notBefore", "body", "date", m.NotBefore.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TLSCertificate) validateSubject(formats strfmt.Registry) error {

	if err := validate.Required("subject", "body", m.Subject); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this TLS certificate based on context it is used
func (m *TLSCertificate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TLSCertificate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TLSCertificate) UnmarshalBinary(b []byte) error {
	var res TLSCertificate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
