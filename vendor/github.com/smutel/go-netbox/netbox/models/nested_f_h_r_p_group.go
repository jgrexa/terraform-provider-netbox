// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2020 Samuel Mutel <12967891+smutel@users.noreply.github.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NestedFHRPGroup nested f h r p group
//
// swagger:model NestedFHRPGroup
type NestedFHRPGroup struct {

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// Group ID
	// Required: true
	// Maximum: 32767
	// Minimum: 0
	GroupID *int64 `json:"group_id"`

	// Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Protocol
	// Required: true
	// Enum: [vrrp2 vrrp3 carp clusterxl hsrp glbp other]
	Protocol *string `json:"protocol"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this nested f h r p group
func (m *NestedFHRPGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedFHRPGroup) validateGroupID(formats strfmt.Registry) error {

	if err := validate.Required("group_id", "body", m.GroupID); err != nil {
		return err
	}

	if err := validate.MinimumInt("group_id", "body", *m.GroupID, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("group_id", "body", *m.GroupID, 32767, false); err != nil {
		return err
	}

	return nil
}

var nestedFHRPGroupTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["vrrp2","vrrp3","carp","clusterxl","hsrp","glbp","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nestedFHRPGroupTypeProtocolPropEnum = append(nestedFHRPGroupTypeProtocolPropEnum, v)
	}
}

const (

	// NestedFHRPGroupProtocolVrrp2 captures enum value "vrrp2"
	NestedFHRPGroupProtocolVrrp2 string = "vrrp2"

	// NestedFHRPGroupProtocolVrrp3 captures enum value "vrrp3"
	NestedFHRPGroupProtocolVrrp3 string = "vrrp3"

	// NestedFHRPGroupProtocolCarp captures enum value "carp"
	NestedFHRPGroupProtocolCarp string = "carp"

	// NestedFHRPGroupProtocolClusterxl captures enum value "clusterxl"
	NestedFHRPGroupProtocolClusterxl string = "clusterxl"

	// NestedFHRPGroupProtocolHsrp captures enum value "hsrp"
	NestedFHRPGroupProtocolHsrp string = "hsrp"

	// NestedFHRPGroupProtocolGlbp captures enum value "glbp"
	NestedFHRPGroupProtocolGlbp string = "glbp"

	// NestedFHRPGroupProtocolOther captures enum value "other"
	NestedFHRPGroupProtocolOther string = "other"
)

// prop value enum
func (m *NestedFHRPGroup) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nestedFHRPGroupTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NestedFHRPGroup) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", m.Protocol); err != nil {
		return err
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", *m.Protocol); err != nil {
		return err
	}

	return nil
}

func (m *NestedFHRPGroup) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nested f h r p group based on the context it is used
func (m *NestedFHRPGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedFHRPGroup) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *NestedFHRPGroup) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *NestedFHRPGroup) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedFHRPGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedFHRPGroup) UnmarshalBinary(b []byte) error {
	var res NestedFHRPGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
