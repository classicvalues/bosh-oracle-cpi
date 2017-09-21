package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CrossConnectLocation An individual FastConnect location.
// swagger:model CrossConnectLocation
type CrossConnectLocation struct {

	// A description of the location.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	Description *string `json:"description"`

	// The name of the location.
	//
	// Example: `CyrusOne, Chandler, AZ`
	//
	// Required: true
	// Max Length: 255
	// Min Length: 1
	Name *string `json:"name"`
}

// Validate validates this cross connect location
func (m *CrossConnectLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CrossConnectLocation) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 255); err != nil {
		return err
	}

	return nil
}

func (m *CrossConnectLocation) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 255); err != nil {
		return err
	}

	return nil
}