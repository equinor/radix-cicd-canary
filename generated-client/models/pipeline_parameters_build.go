// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PipelineParametersBuild PipelineParametersBuild describe branch to build and its commit ID
// swagger:model PipelineParametersBuild
type PipelineParametersBuild struct {

	// Branch the branch to build
	// REQUIRED for "build" and "build-deploy" pipelines
	Branch string `json:"branch,omitempty"`

	// CommitID the commit ID of the branch to build
	// REQUIRED for "build" and "build-deploy" pipelines
	CommitID string `json:"commitID,omitempty"`

	// PushImage should image be pushed to container registry. Defaults pushing
	PushImage string `json:"pushImage,omitempty"`

	// TriggeredBy of the job - if empty will use user token upn (user principle name)
	TriggeredBy string `json:"triggeredBy,omitempty"`
}

// Validate validates this pipeline parameters build
func (m *PipelineParametersBuild) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PipelineParametersBuild) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PipelineParametersBuild) UnmarshalBinary(b []byte) error {
	var res PipelineParametersBuild
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
