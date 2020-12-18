// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Scan A scan belonging to an environment.
//
// swagger:model Scan
type Scan struct {

	// Time the scan was created.
	CreatedAt int64 `json:"created_at,omitempty"`

	// ID of the environment the scan belongs to.
	EnvironmentID string `json:"environment_id,omitempty"`

	// Time the scan was finished.
	FinishedAt int64 `json:"finished_at,omitempty"`

	// ID of the scan.
	ID string `json:"id,omitempty"`

	// Message related to the scan.
	Message string `json:"message,omitempty"`

	// Indicates whether there were any remediation errors on the scan.
	RemediationError bool `json:"remediation_error,omitempty"`

	// Status of the scan.
	// Enum: [CREATED QUEUED IN_PROGRESS ERROR SUCCESS CANCELED]
	Status string `json:"status,omitempty"`

	// Time the scan was last updated.
	UpdatedAt int64 `json:"updated_at,omitempty"`
}

// Validate validates this scan
func (m *Scan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var scanTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATED","QUEUED","IN_PROGRESS","ERROR","SUCCESS","CANCELED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scanTypeStatusPropEnum = append(scanTypeStatusPropEnum, v)
	}
}

const (

	// ScanStatusCREATED captures enum value "CREATED"
	ScanStatusCREATED string = "CREATED"

	// ScanStatusQUEUED captures enum value "QUEUED"
	ScanStatusQUEUED string = "QUEUED"

	// ScanStatusINPROGRESS captures enum value "IN_PROGRESS"
	ScanStatusINPROGRESS string = "IN_PROGRESS"

	// ScanStatusERROR captures enum value "ERROR"
	ScanStatusERROR string = "ERROR"

	// ScanStatusSUCCESS captures enum value "SUCCESS"
	ScanStatusSUCCESS string = "SUCCESS"

	// ScanStatusCANCELED captures enum value "CANCELED"
	ScanStatusCANCELED string = "CANCELED"
)

// prop value enum
func (m *Scan) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, scanTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Scan) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Scan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Scan) UnmarshalBinary(b []byte) error {
	var res Scan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
