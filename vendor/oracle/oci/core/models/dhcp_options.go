package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DhcpOptions A set of DHCP options. Used by the VCN to automatically provide configuration
// information to the instances when they boot up. There are two options you can set:
//
// - [DhcpDnsOption](#/en/iaas/20160918/DhcpDnsOption/): Lets you specify how DNS (hostname resolution) is
// handled in the subnets in your VCN.
//
// - [DhcpSearchDomainOption](#/en/iaas/20160918/DhcpSearchDomainOption/): Lets you specify
// a search domain name to use for DNS queries.
//
// For more information, see  [DNS in Your Virtual Cloud Network](/Content/Network/Concepts/dns.htm)
// and [Managing DHCP Options](/Content/Network/Tasks/managingDHCP.htm).
//
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies](/Content/Identity/Concepts/policygetstarted.htm).
//
// swagger:model DhcpOptions
type DhcpOptions struct {

	// The OCID of the compartment containing the set of DHCP options.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	CompartmentID *string `json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	//
	// Max Length: 255
	// Min Length: 1
	DisplayName string `json:"displayName,omitempty"`

	// Oracle ID (OCID) for the set of DHCP options.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	ID *string `json:"id"`

	// The current state of the set of DHCP options.
	// Required: true
	LifecycleState *string `json:"lifecycleState"`

	Options []DhcpOption `json:"options"`

	// Date and time the set of DHCP options was created, in the format defined by RFC3339.
	//
	// Example: `2016-08-25T21:10:29.600Z`
	//
	// Required: true
	TimeCreated *strfmt.DateTime `json:"timeCreated"`

	// The OCID of the VCN the set of DHCP options belongs to.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	VcnID *string `json:"vcnId"`
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *DhcpOptions) UnmarshalJSON(raw []byte) error {
	var data struct {
		CompartmentID *string `json:"compartmentId"`

		DisplayName string `json:"displayName,omitempty"`

		ID *string `json:"id"`

		LifecycleState *string `json:"lifecycleState"`

		Options json.RawMessage `json:"options"`

		TimeCreated *strfmt.DateTime `json:"timeCreated"`

		VcnID *string `json:"vcnId"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var options []DhcpOption

	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	untypedObj := make(map[string]interface{})
	if err := dec.Decode(&untypedObj); err != nil {
		return err
	}
	if untypedOptions, ok := untypedObj["options"]; ok {
		if slcOptions, ok := untypedOptions.([]interface{}); ok {
			for _, slcEl := range slcOptions {
				slcJSON, _ := json.Marshal(slcEl)
				slcObj, err := UnmarshalDhcpOption(bytes.NewBuffer(slcJSON), runtime.JSONConsumer())
				if err != nil && err != io.EOF {
					return err
				}
				options = append(options, slcObj)
			}
		}
	}

	var result DhcpOptions
	result.CompartmentID = data.CompartmentID
	result.DisplayName = data.DisplayName
	result.ID = data.ID
	result.LifecycleState = data.LifecycleState
	result.Options = options
	result.TimeCreated = data.TimeCreated
	result.VcnID = data.VcnID
	*m = result
	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m DhcpOptions) MarshalJSON() ([]byte, error) {
	var b1, b2 []byte
	var err error
	b1, err = json.Marshal(struct {
		CompartmentID *string `json:"compartmentId"`

		DisplayName string `json:"displayName,omitempty"`

		ID *string `json:"id"`

		LifecycleState *string `json:"lifecycleState"`

		TimeCreated *strfmt.DateTime `json:"timeCreated"`

		VcnID *string `json:"vcnId"`
	}{
		CompartmentID:  m.CompartmentID,
		DisplayName:    m.DisplayName,
		ID:             m.ID,
		LifecycleState: m.LifecycleState,
		TimeCreated:    m.TimeCreated,
		VcnID:          m.VcnID,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Options []DhcpOption `json:"options"`
	}{
		Options: m.Options,
	})
	if err != nil {
		return nil, err
	}
	return swag.ConcatJSON(b1, b2), nil
}

// Validate validates this dhcp options
func (m *DhcpOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompartmentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLifecycleState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTimeCreated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVcnID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DhcpOptions) validateCompartmentID(formats strfmt.Registry) error {

	if err := validate.Required("compartmentId", "body", m.CompartmentID); err != nil {
		return err
	}

	if err := validate.MinLength("compartmentId", "body", string(*m.CompartmentID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("compartmentId", "body", string(*m.CompartmentID), 255); err != nil {
		return err
	}

	return nil
}

func (m *DhcpOptions) validateDisplayName(formats strfmt.Registry) error {

	if swag.IsZero(m.DisplayName) { // not required
		return nil
	}

	if err := validate.MinLength("displayName", "body", string(m.DisplayName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("displayName", "body", string(m.DisplayName), 255); err != nil {
		return err
	}

	return nil
}

func (m *DhcpOptions) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", string(*m.ID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("id", "body", string(*m.ID), 255); err != nil {
		return err
	}

	return nil
}

var dhcpOptionsTypeLifecycleStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PROVISIONING","AVAILABLE","TERMINATING","TERMINATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dhcpOptionsTypeLifecycleStatePropEnum = append(dhcpOptionsTypeLifecycleStatePropEnum, v)
	}
}

const (
	// DhcpOptionsLifecycleStatePROVISIONING captures enum value "PROVISIONING"
	DhcpOptionsLifecycleStatePROVISIONING string = "PROVISIONING"
	// DhcpOptionsLifecycleStateAVAILABLE captures enum value "AVAILABLE"
	DhcpOptionsLifecycleStateAVAILABLE string = "AVAILABLE"
	// DhcpOptionsLifecycleStateTERMINATING captures enum value "TERMINATING"
	DhcpOptionsLifecycleStateTERMINATING string = "TERMINATING"
	// DhcpOptionsLifecycleStateTERMINATED captures enum value "TERMINATED"
	DhcpOptionsLifecycleStateTERMINATED string = "TERMINATED"
)

// prop value enum
func (m *DhcpOptions) validateLifecycleStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, dhcpOptionsTypeLifecycleStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DhcpOptions) validateLifecycleState(formats strfmt.Registry) error {

	if err := validate.Required("lifecycleState", "body", m.LifecycleState); err != nil {
		return err
	}

	// value enum
	if err := m.validateLifecycleStateEnum("lifecycleState", "body", *m.LifecycleState); err != nil {
		return err
	}

	return nil
}

func (m *DhcpOptions) validateOptions(formats strfmt.Registry) error {

	if err := validate.Required("options", "body", m.Options); err != nil {
		return err
	}

	for i := 0; i < len(m.Options); i++ {

		if err := m.Options[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("options" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *DhcpOptions) validateTimeCreated(formats strfmt.Registry) error {

	if err := validate.Required("timeCreated", "body", m.TimeCreated); err != nil {
		return err
	}

	return nil
}

func (m *DhcpOptions) validateVcnID(formats strfmt.Registry) error {

	if err := validate.Required("vcnId", "body", m.VcnID); err != nil {
		return err
	}

	if err := validate.MinLength("vcnId", "body", string(*m.VcnID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("vcnId", "body", string(*m.VcnID), 255); err != nil {
		return err
	}

	return nil
}