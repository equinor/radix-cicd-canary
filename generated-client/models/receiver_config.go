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

// ReceiverConfig ReceiverConfig receiver configuration
//
// swagger:model ReceiverConfig
type ReceiverConfig struct {

	// slack config
	// Required: true
	SlackConfig *SlackConfig `json:"slackConfig"`
}

// Validate validates this receiver config
func (m *ReceiverConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSlackConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReceiverConfig) validateSlackConfig(formats strfmt.Registry) error {

	if err := validate.Required("slackConfig", "body", m.SlackConfig); err != nil {
		return err
	}

	if m.SlackConfig != nil {
		if err := m.SlackConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slackConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("slackConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this receiver config based on the context it is used
func (m *ReceiverConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSlackConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReceiverConfig) contextValidateSlackConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.SlackConfig != nil {
		if err := m.SlackConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slackConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("slackConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReceiverConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReceiverConfig) UnmarshalBinary(b []byte) error {
	var res ReceiverConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
