// Code generated by go-swagger; DO NOT EDIT.

package workflow_triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/puppetlabs/nebula-cli/pkg/client/api/models"
)

// UpdateWorkflowTriggerReader is a Reader for the UpdateWorkflowTrigger structure.
type UpdateWorkflowTriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateWorkflowTriggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateWorkflowTriggerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateWorkflowTriggerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateWorkflowTriggerOK creates a UpdateWorkflowTriggerOK with default headers values
func NewUpdateWorkflowTriggerOK() *UpdateWorkflowTriggerOK {
	return &UpdateWorkflowTriggerOK{}
}

/*UpdateWorkflowTriggerOK handles this case with default header values.

An updated workflow trigger
*/
type UpdateWorkflowTriggerOK struct {
	Payload *UpdateWorkflowTriggerOKBody
}

func (o *UpdateWorkflowTriggerOK) Error() string {
	return fmt.Sprintf("[PUT /api/workflows/{workflowName}/triggers/{workflowTriggerId}][%d] updateWorkflowTriggerOK  %+v", 200, o.Payload)
}

func (o *UpdateWorkflowTriggerOK) GetPayload() *UpdateWorkflowTriggerOKBody {
	return o.Payload
}

func (o *UpdateWorkflowTriggerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateWorkflowTriggerOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateWorkflowTriggerDefault creates a UpdateWorkflowTriggerDefault with default headers values
func NewUpdateWorkflowTriggerDefault(code int) *UpdateWorkflowTriggerDefault {
	return &UpdateWorkflowTriggerDefault{
		_statusCode: code,
	}
}

/*UpdateWorkflowTriggerDefault handles this case with default header values.

An error occurred
*/
type UpdateWorkflowTriggerDefault struct {
	_statusCode int

	Payload *UpdateWorkflowTriggerDefaultBody
}

// Code gets the status code for the update workflow trigger default response
func (o *UpdateWorkflowTriggerDefault) Code() int {
	return o._statusCode
}

func (o *UpdateWorkflowTriggerDefault) Error() string {
	return fmt.Sprintf("[PUT /api/workflows/{workflowName}/triggers/{workflowTriggerId}][%d] updateWorkflowTrigger default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateWorkflowTriggerDefault) GetPayload() *UpdateWorkflowTriggerDefaultBody {
	return o.Payload
}

func (o *UpdateWorkflowTriggerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateWorkflowTriggerDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateWorkflowTriggerBody The request type for updatig a workflow trigger
swagger:model UpdateWorkflowTriggerBody
*/
type UpdateWorkflowTriggerBody struct {

	// binding
	Binding *models.WorkflowTriggerBinding `json:"binding,omitempty"`

	// settings
	Settings models.WorkflowTriggerSettings `json:"settings,omitempty"`
}

// Validate validates this update workflow trigger body
func (o *UpdateWorkflowTriggerBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBinding(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateWorkflowTriggerBody) validateBinding(formats strfmt.Registry) error {

	if swag.IsZero(o.Binding) { // not required
		return nil
	}

	if o.Binding != nil {
		if err := o.Binding.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "binding")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateWorkflowTriggerBody) validateSettings(formats strfmt.Registry) error {

	if swag.IsZero(o.Settings) { // not required
		return nil
	}

	if err := o.Settings.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "settings")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateWorkflowTriggerBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateWorkflowTriggerBody) UnmarshalBinary(b []byte) error {
	var res UpdateWorkflowTriggerBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateWorkflowTriggerDefaultBody Error response
swagger:model UpdateWorkflowTriggerDefaultBody
*/
type UpdateWorkflowTriggerDefaultBody struct {

	// error
	Error *models.Error `json:"error,omitempty"`
}

// Validate validates this update workflow trigger default body
func (o *UpdateWorkflowTriggerDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateWorkflowTriggerDefaultBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateWorkflowTrigger default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateWorkflowTriggerDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateWorkflowTriggerDefaultBody) UnmarshalBinary(b []byte) error {
	var res UpdateWorkflowTriggerDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateWorkflowTriggerOKBody The response type for the updated trigger
swagger:model UpdateWorkflowTriggerOKBody
*/
type UpdateWorkflowTriggerOKBody struct {

	// trigger
	Trigger *models.WorkflowTrigger `json:"trigger,omitempty"`
}

// Validate validates this update workflow trigger o k body
func (o *UpdateWorkflowTriggerOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTrigger(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateWorkflowTriggerOKBody) validateTrigger(formats strfmt.Registry) error {

	if swag.IsZero(o.Trigger) { // not required
		return nil
	}

	if o.Trigger != nil {
		if err := o.Trigger.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateWorkflowTriggerOK" + "." + "trigger")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateWorkflowTriggerOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateWorkflowTriggerOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateWorkflowTriggerOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}