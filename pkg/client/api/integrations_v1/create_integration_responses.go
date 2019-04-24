// Code generated by go-swagger; DO NOT EDIT.

package integrations_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/puppetlabs/nebula/pkg/client/api/models"
)

// CreateIntegrationReader is a Reader for the CreateIntegration structure.
type CreateIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateIntegrationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateIntegrationCreated creates a CreateIntegrationCreated with default headers values
func NewCreateIntegrationCreated() *CreateIntegrationCreated {
	return &CreateIntegrationCreated{}
}

/*CreateIntegrationCreated handles this case with default header values.

The newly created integration
*/
type CreateIntegrationCreated struct {
	Payload *models.ShowIntegration
}

func (o *CreateIntegrationCreated) Error() string {
	return fmt.Sprintf("[POST /api/integrations][%d] createIntegrationCreated  %+v", 201, o.Payload)
}

func (o *CreateIntegrationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ShowIntegration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}