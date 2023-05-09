// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Component Component describe an component part of an deployment
//
// swagger:model Component
type Component struct {

	// Image name
	// Example: radixdev.azurecr.io/app-server:cdgkg
	// Required: true
	Image *string `json:"image"`

	// Name the component
	// Example: server
	// Required: true
	Name *string `json:"name"`

	// Ports defines the port number and protocol that a component is exposed for internally in environment
	Ports []*Port `json:"ports"`

	// Array of ReplicaSummary
	ReplicaList []*ReplicaSummary `json:"replicaList"`

	// Array of pod names
	// Example: ["server-78fc8857c4-hm76l","server-78fc8857c4-asfa2"]
	Replicas []string `json:"replicas"`

	// ScheduledJobPayloadPath defines the payload path, where payload for Job Scheduler will be mapped as a file. From radixconfig.yaml
	// Example: \"/tmp/payload\
	ScheduledJobPayloadPath string `json:"scheduledJobPayloadPath,omitempty"`

	// SchedulerPort defines the port number that a Job Scheduler is exposed internally in environment
	// Example: 8080
	SchedulerPort int32 `json:"schedulerPort,omitempty"`

	// Component secret names. From radixconfig.yaml
	// Example: ["DB_CON","A_SECRET"]
	Secrets []string `json:"secrets"`

	// Status of the component
	// Example: Consistent
	Status string `json:"status,omitempty"`

	// Type of component
	// Example: component
	// Required: true
	Type *string `json:"type"`

	// Variable names map to values. From radixconfig.yaml
	Variables map[string]string `json:"variables,omitempty"`

	// horizontal scaling summary
	HorizontalScalingSummary *HorizontalScalingSummary `json:"horizontalScalingSummary,omitempty"`

	// identity
	Identity *Identity `json:"identity,omitempty"`

	// notifications
	Notifications *Notifications `json:"notifications,omitempty"`

	// oauth2
	Oauth2 *OAuth2AuxiliaryResource `json:"oauth2,omitempty"`
}

// Validate validates this component
func (m *Component) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicaList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHorizontalScalingSummary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOauth2(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Component) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Component) validatePorts(formats strfmt.Registry) error {
	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	for i := 0; i < len(m.Ports); i++ {
		if swag.IsZero(m.Ports[i]) { // not required
			continue
		}

		if m.Ports[i] != nil {
			if err := m.Ports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Component) validateReplicaList(formats strfmt.Registry) error {
	if swag.IsZero(m.ReplicaList) { // not required
		return nil
	}

	for i := 0; i < len(m.ReplicaList); i++ {
		if swag.IsZero(m.ReplicaList[i]) { // not required
			continue
		}

		if m.ReplicaList[i] != nil {
			if err := m.ReplicaList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("replicaList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("replicaList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Component) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Component) validateHorizontalScalingSummary(formats strfmt.Registry) error {
	if swag.IsZero(m.HorizontalScalingSummary) { // not required
		return nil
	}

	if m.HorizontalScalingSummary != nil {
		if err := m.HorizontalScalingSummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("horizontalScalingSummary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("horizontalScalingSummary")
			}
			return err
		}
	}

	return nil
}

func (m *Component) validateIdentity(formats strfmt.Registry) error {
	if swag.IsZero(m.Identity) { // not required
		return nil
	}

	if m.Identity != nil {
		if err := m.Identity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identity")
			}
			return err
		}
	}

	return nil
}

func (m *Component) validateNotifications(formats strfmt.Registry) error {
	if swag.IsZero(m.Notifications) { // not required
		return nil
	}

	if m.Notifications != nil {
		if err := m.Notifications.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifications")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notifications")
			}
			return err
		}
	}

	return nil
}

func (m *Component) validateOauth2(formats strfmt.Registry) error {
	if swag.IsZero(m.Oauth2) { // not required
		return nil
	}

	if m.Oauth2 != nil {
		if err := m.Oauth2.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oauth2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oauth2")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this component based on the context it is used
func (m *Component) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReplicaList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHorizontalScalingSummary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotifications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOauth2(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Component) contextValidatePorts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ports); i++ {

		if m.Ports[i] != nil {
			if err := m.Ports[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Component) contextValidateReplicaList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ReplicaList); i++ {

		if m.ReplicaList[i] != nil {
			if err := m.ReplicaList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("replicaList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("replicaList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Component) contextValidateHorizontalScalingSummary(ctx context.Context, formats strfmt.Registry) error {

	if m.HorizontalScalingSummary != nil {
		if err := m.HorizontalScalingSummary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("horizontalScalingSummary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("horizontalScalingSummary")
			}
			return err
		}
	}

	return nil
}

func (m *Component) contextValidateIdentity(ctx context.Context, formats strfmt.Registry) error {

	if m.Identity != nil {
		if err := m.Identity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identity")
			}
			return err
		}
	}

	return nil
}

func (m *Component) contextValidateNotifications(ctx context.Context, formats strfmt.Registry) error {

	if m.Notifications != nil {
		if err := m.Notifications.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifications")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notifications")
			}
			return err
		}
	}

	return nil
}

func (m *Component) contextValidateOauth2(ctx context.Context, formats strfmt.Registry) error {

	if m.Oauth2 != nil {
		if err := m.Oauth2.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oauth2")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oauth2")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Component) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Component) UnmarshalBinary(b []byte) error {
	var res Component
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
