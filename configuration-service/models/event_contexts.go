// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// EventContexts event contexts
// swagger:model EventContexts
type EventContexts struct {

	// event contexts
	EventContexts []*EventContext `json:"eventContexts"`

	// Pointer to next page, base64 encoded
	NextPageKey string `json:"nextPageKey,omitempty"`

	// Size of returned page
	PageSize float64 `json:"pageSize,omitempty"`

	// Total number of stages
	TotalCount float64 `json:"totalCount,omitempty"`
}

// Validate validates this event contexts
func (m *EventContexts) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventContexts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventContexts) validateEventContexts(formats strfmt.Registry) error {

	if swag.IsZero(m.EventContexts) { // not required
		return nil
	}

	for i := 0; i < len(m.EventContexts); i++ {
		if swag.IsZero(m.EventContexts[i]) { // not required
			continue
		}

		if m.EventContexts[i] != nil {
			if err := m.EventContexts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("eventContexts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EventContexts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventContexts) UnmarshalBinary(b []byte) error {
	var res EventContexts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}