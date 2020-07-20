// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha1Cluster Cluster is the definition of a cluster resource
//
// swagger:model v1alpha1Cluster
type V1alpha1Cluster struct {

	// config
	Config *V1alpha1ClusterConfig `json:"config,omitempty"`

	// connection state
	ConnectionState *V1alpha1ConnectionState `json:"connectionState,omitempty"`

	// Name of the cluster. If omitted, will use the server address
	Name string `json:"name,omitempty"`

	// Holds list of namespaces which are accessible in that cluster. Cluster level resources would be ignored if namespace list if not empty.
	Namespaces []string `json:"namespaces"`

	// Server is the API server URL of the Kubernetes cluster
	Server string `json:"server,omitempty"`

	// The server version
	ServerVersion string `json:"serverVersion,omitempty"`
}

// Validate validates this v1alpha1 cluster
func (m *V1alpha1Cluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1Cluster) validateConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1Cluster) validateConnectionState(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionState) { // not required
		return nil
	}

	if m.ConnectionState != nil {
		if err := m.ConnectionState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionState")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1Cluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1Cluster) UnmarshalBinary(b []byte) error {
	var res V1alpha1Cluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}