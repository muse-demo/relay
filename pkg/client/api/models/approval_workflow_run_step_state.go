// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ApprovalWorkflowRunStepState approval workflow run step state
// swagger:model ApprovalWorkflowRunStepState
type ApprovalWorkflowRunStepState struct {
	AnyWorkflowRunStepState

	// Workflow run step approval for manual gates
	// Enum: [waiting rejected approved]
	Approval *string `json:"approval,omitempty"`

	// Time at which the step was approved or rejected
	// Format: date-time
	ApprovalSubmittedAt *strfmt.DateTime `json:"approval_submitted_at,omitempty"`

	// approver
	Approver *UserSummary `json:"approver,omitempty"`

	// Type of step
	// Required: true
	// Enum: [approval]
	Type *string `json:"type"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ApprovalWorkflowRunStepState) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 AnyWorkflowRunStepState
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.AnyWorkflowRunStepState = aO0

	// AO1
	var dataAO1 struct {
		Approval *string `json:"approval,omitempty"`

		ApprovalSubmittedAt *strfmt.DateTime `json:"approval_submitted_at,omitempty"`

		Approver *UserSummary `json:"approver,omitempty"`

		Type *string `json:"type"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Approval = dataAO1.Approval

	m.ApprovalSubmittedAt = dataAO1.ApprovalSubmittedAt

	m.Approver = dataAO1.Approver

	m.Type = dataAO1.Type

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ApprovalWorkflowRunStepState) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.AnyWorkflowRunStepState)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Approval *string `json:"approval,omitempty"`

		ApprovalSubmittedAt *strfmt.DateTime `json:"approval_submitted_at,omitempty"`

		Approver *UserSummary `json:"approver,omitempty"`

		Type *string `json:"type"`
	}

	dataAO1.Approval = m.Approval

	dataAO1.ApprovalSubmittedAt = m.ApprovalSubmittedAt

	dataAO1.Approver = m.Approver

	dataAO1.Type = m.Type

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this approval workflow run step state
func (m *ApprovalWorkflowRunStepState) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with AnyWorkflowRunStepState
	if err := m.AnyWorkflowRunStepState.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApproval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApprovalSubmittedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApprover(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var approvalWorkflowRunStepStateTypeApprovalPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["waiting","rejected","approved"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		approvalWorkflowRunStepStateTypeApprovalPropEnum = append(approvalWorkflowRunStepStateTypeApprovalPropEnum, v)
	}
}

// property enum
func (m *ApprovalWorkflowRunStepState) validateApprovalEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, approvalWorkflowRunStepStateTypeApprovalPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ApprovalWorkflowRunStepState) validateApproval(formats strfmt.Registry) error {

	if swag.IsZero(m.Approval) { // not required
		return nil
	}

	// value enum
	if err := m.validateApprovalEnum("approval", "body", *m.Approval); err != nil {
		return err
	}

	return nil
}

func (m *ApprovalWorkflowRunStepState) validateApprovalSubmittedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ApprovalSubmittedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("approval_submitted_at", "body", "date-time", m.ApprovalSubmittedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ApprovalWorkflowRunStepState) validateApprover(formats strfmt.Registry) error {

	if swag.IsZero(m.Approver) { // not required
		return nil
	}

	if m.Approver != nil {
		if err := m.Approver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("approver")
			}
			return err
		}
	}

	return nil
}

var approvalWorkflowRunStepStateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["approval"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		approvalWorkflowRunStepStateTypeTypePropEnum = append(approvalWorkflowRunStepStateTypeTypePropEnum, v)
	}
}

// property enum
func (m *ApprovalWorkflowRunStepState) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, approvalWorkflowRunStepStateTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ApprovalWorkflowRunStepState) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApprovalWorkflowRunStepState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApprovalWorkflowRunStepState) UnmarshalBinary(b []byte) error {
	var res ApprovalWorkflowRunStepState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
