// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PipelineParametersDeploy PipelineParametersDeploy describes environment to deploy
//
// swagger:model PipelineParametersDeploy
type PipelineParametersDeploy struct {

	// Name of environment to deploy
	// REQUIRED for "deploy" pipeline
	ToEnvironment string `json:"toEnvironment,omitempty"`

	// TriggeredBy of the job - if empty will use user token upn (user principle name)
	TriggeredBy string `json:"triggeredBy,omitempty"`
}

// Validate validates this pipeline parameters deploy
func (m *PipelineParametersDeploy) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PipelineParametersDeploy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelineParametersDeploy) UnmarshalBinary(b []byte) error {
	var res PipelineParametersDeploy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
