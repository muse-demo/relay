// Code generated by go-swagger; DO NOT EDIT.

package workflows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/puppetlabs/nebula-cli/pkg/client/api/models"
)

// GetWorkflowReader is a Reader for the GetWorkflow structure.
type GetWorkflowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWorkflowDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWorkflowOK creates a GetWorkflowOK with default headers values
func NewGetWorkflowOK() *GetWorkflowOK {
	return &GetWorkflowOK{}
}

/*GetWorkflowOK handles this case with default header values.

The retrieved workflow information
*/
type GetWorkflowOK struct {
	Payload *GetWorkflowOKBody
}

func (o *GetWorkflowOK) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflowName}][%d] getWorkflowOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowOK) GetPayload() *GetWorkflowOKBody {
	return o.Payload
}

func (o *GetWorkflowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowDefault creates a GetWorkflowDefault with default headers values
func NewGetWorkflowDefault(code int) *GetWorkflowDefault {
	return &GetWorkflowDefault{
		_statusCode: code,
	}
}

/*GetWorkflowDefault handles this case with default header values.

An error occurred
*/
type GetWorkflowDefault struct {
	_statusCode int

	Payload *GetWorkflowDefaultBody
}

// Code gets the status code for the get workflow default response
func (o *GetWorkflowDefault) Code() int {
	return o._statusCode
}

func (o *GetWorkflowDefault) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflowName}][%d] getWorkflow default  %+v", o._statusCode, o.Payload)
}

func (o *GetWorkflowDefault) GetPayload() *GetWorkflowDefaultBody {
	return o.Payload
}

func (o *GetWorkflowDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetWorkflowDefaultBody Error response
swagger:model GetWorkflowDefaultBody
*/
type GetWorkflowDefaultBody struct {

	// error
	Error *models.Error `json:"error,omitempty"`
}

// Validate validates this get workflow default body
func (o *GetWorkflowDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowDefaultBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWorkflow default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWorkflowOKBody get workflow o k body
swagger:model GetWorkflowOKBody
*/
type GetWorkflowOKBody struct {
	models.WorkflowEntity

	// workflow
	Workflow *models.Workflow `json:"workflow,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetWorkflowOKBody) UnmarshalJSON(raw []byte) error {
	// GetWorkflowOKBodyAO0
	var getWorkflowOKBodyAO0 models.WorkflowEntity
	if err := swag.ReadJSON(raw, &getWorkflowOKBodyAO0); err != nil {
		return err
	}
	o.WorkflowEntity = getWorkflowOKBodyAO0

	// GetWorkflowOKBodyAO1
	var dataGetWorkflowOKBodyAO1 struct {
		Workflow *models.Workflow `json:"workflow,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetWorkflowOKBodyAO1); err != nil {
		return err
	}

	o.Workflow = dataGetWorkflowOKBodyAO1.Workflow

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetWorkflowOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getWorkflowOKBodyAO0, err := swag.WriteJSON(o.WorkflowEntity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getWorkflowOKBodyAO0)

	var dataGetWorkflowOKBodyAO1 struct {
		Workflow *models.Workflow `json:"workflow,omitempty"`
	}

	dataGetWorkflowOKBodyAO1.Workflow = o.Workflow

	jsonDataGetWorkflowOKBodyAO1, errGetWorkflowOKBodyAO1 := swag.WriteJSON(dataGetWorkflowOKBodyAO1)
	if errGetWorkflowOKBodyAO1 != nil {
		return nil, errGetWorkflowOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetWorkflowOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get workflow o k body
func (o *GetWorkflowOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.WorkflowEntity
	if err := o.WorkflowEntity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWorkflow(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowOKBody) validateWorkflow(formats strfmt.Registry) error {

	if swag.IsZero(o.Workflow) { // not required
		return nil
	}

	if o.Workflow != nil {
		if err := o.Workflow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWorkflowOK" + "." + "workflow")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowOKBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}