// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeviceType device type
// swagger:model DeviceType
type DeviceType struct {

	// Comments
	Comments string `json:"comments,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Instance count
	// Read Only: true
	InstanceCount int64 `json:"instance_count,omitempty"`

	// interface ordering
	// Required: true
	InterfaceOrdering *DeviceTypeInterfaceOrdering `json:"interface_ordering"`

	// Is a console server
	//
	// This type of device has console server ports
	IsConsoleServer bool `json:"is_console_server,omitempty"`

	// Is full depth
	//
	// Device consumes both front and rear rack faces
	IsFullDepth bool `json:"is_full_depth,omitempty"`

	// Is a network device
	//
	// This type of device has network interfaces
	IsNetworkDevice bool `json:"is_network_device,omitempty"`

	// Is a PDU
	//
	// This type of device has power outlets
	IsPdu bool `json:"is_pdu,omitempty"`

	// manufacturer
	// Required: true
	Manufacturer *NestedManufacturer `json:"manufacturer"`

	// Model
	// Required: true
	// Max Length: 50
	Model *string `json:"model"`

	// Part number
	//
	// Discrete part number (optional)
	// Max Length: 50
	PartNumber string `json:"part_number,omitempty"`

	// Slug
	// Required: true
	// Max Length: 50
	// Pattern: ^[-a-zA-Z0-9_]+$
	Slug *string `json:"slug"`

	// subdevice role
	// Required: true
	SubdeviceRole *DeviceTypeSubdeviceRole `json:"subdevice_role"`

	// Height (U)
	// Maximum: 32767
	// Minimum: 0
	UHeight *int64 `json:"u_height,omitempty"`
}

// Validate validates this device type
func (m *DeviceType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInterfaceOrdering(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateManufacturer(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePartNumber(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSlug(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubdeviceRole(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUHeight(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceType) validateInterfaceOrdering(formats strfmt.Registry) error {

	if err := validate.Required("interface_ordering", "body", m.InterfaceOrdering); err != nil {
		return err
	}

	if m.InterfaceOrdering != nil {

		if err := m.InterfaceOrdering.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface_ordering")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceType) validateManufacturer(formats strfmt.Registry) error {

	if err := validate.Required("manufacturer", "body", m.Manufacturer); err != nil {
		return err
	}

	if m.Manufacturer != nil {

		if err := m.Manufacturer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("manufacturer")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceType) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	if err := validate.MaxLength("model", "body", string(*m.Model), 50); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) validatePartNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.PartNumber) { // not required
		return nil
	}

	if err := validate.MaxLength("part_number", "body", string(m.PartNumber), 50); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) validateSlug(formats strfmt.Registry) error {

	if err := validate.Required("slug", "body", m.Slug); err != nil {
		return err
	}

	if err := validate.MaxLength("slug", "body", string(*m.Slug), 50); err != nil {
		return err
	}

	if err := validate.Pattern("slug", "body", string(*m.Slug), `^[-a-zA-Z0-9_]+$`); err != nil {
		return err
	}

	return nil
}

func (m *DeviceType) validateSubdeviceRole(formats strfmt.Registry) error {

	if err := validate.Required("subdevice_role", "body", m.SubdeviceRole); err != nil {
		return err
	}

	if m.SubdeviceRole != nil {

		if err := m.SubdeviceRole.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subdevice_role")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceType) validateUHeight(formats strfmt.Registry) error {

	if swag.IsZero(m.UHeight) { // not required
		return nil
	}

	if err := validate.MinimumInt("u_height", "body", int64(*m.UHeight), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("u_height", "body", int64(*m.UHeight), 32767, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceType) UnmarshalBinary(b []byte) error {
	var res DeviceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}