// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PipelineParametersPromote PipelineParametersPromote identify deployment to promote and a target environment
// swagger:model PipelineParametersPromote
type PipelineParametersPromote struct {

	// ID of the deployment to promote
	// REQUIRED for "promote" pipeline
	DeploymentName string `json:"deploymentName,omitempty"`

	// Name of environment where to look for the deployment to be promoted
	// REQUIRED for "promote" pipeline
	FromEnvironment string `json:"fromEnvironment,omitempty"`

	// Name of environment to receive the promoted deployment
	// REQUIRED for "promote" pipeline
	ToEnvironment string `json:"toEnvironment,omitempty"`
}

// Validate validates this pipeline parameters promote
func (m *PipelineParametersPromote) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PipelineParametersPromote) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelineParametersPromote) UnmarshalBinary(b []byte) error {
	var res PipelineParametersPromote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
