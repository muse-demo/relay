// Code generated by go-swagger; DO NOT EDIT.

package workflow_secrets_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/puppetlabs/nebula/pkg/client/api/models"
)

// NewUpdateSecretByKeyAndWorkflowIDParams creates a new UpdateSecretByKeyAndWorkflowIDParams object
// with the default values initialized.
func NewUpdateSecretByKeyAndWorkflowIDParams() *UpdateSecretByKeyAndWorkflowIDParams {
	var ()
	return &UpdateSecretByKeyAndWorkflowIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSecretByKeyAndWorkflowIDParamsWithTimeout creates a new UpdateSecretByKeyAndWorkflowIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSecretByKeyAndWorkflowIDParamsWithTimeout(timeout time.Duration) *UpdateSecretByKeyAndWorkflowIDParams {
	var ()
	return &UpdateSecretByKeyAndWorkflowIDParams{

		timeout: timeout,
	}
}

// NewUpdateSecretByKeyAndWorkflowIDParamsWithContext creates a new UpdateSecretByKeyAndWorkflowIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSecretByKeyAndWorkflowIDParamsWithContext(ctx context.Context) *UpdateSecretByKeyAndWorkflowIDParams {
	var ()
	return &UpdateSecretByKeyAndWorkflowIDParams{

		Context: ctx,
	}
}

// NewUpdateSecretByKeyAndWorkflowIDParamsWithHTTPClient creates a new UpdateSecretByKeyAndWorkflowIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSecretByKeyAndWorkflowIDParamsWithHTTPClient(client *http.Client) *UpdateSecretByKeyAndWorkflowIDParams {
	var ()
	return &UpdateSecretByKeyAndWorkflowIDParams{
		HTTPClient: client,
	}
}

/*UpdateSecretByKeyAndWorkflowIDParams contains all the parameters to send to the API endpoint
for the update secret by key and workflow ID operation typically these are written to a http.Request
*/
type UpdateSecretByKeyAndWorkflowIDParams struct {

	/*Accept
	  The version of the API, in this case should be "application/nebula-api.v1+json"

	*/
	Accept string
	/*Body
	  Secret to key value pair to update

	*/
	Body *models.UpdateSecret
	/*SecretKey
	  Key for a workflow secret

	*/
	SecretKey string
	/*WorkflowName
	  Workflow name. Must be unique within a user account

	*/
	WorkflowName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update secret by key and workflow ID params
func (o *UpdateSecretByKeyAndWorkflowIDParams) WithTimeout(timeout time.Duration) *UpdateSecretByKeyAndWorkflowIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update secret by key and workflow ID params
func (o *UpdateSecretByKeyAndWorkflowIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update secret by key and workflow ID params
func (o *UpdateSecretByKeyAndWorkflowIDParams) WithContext(ctx context.Context) *UpdateSecretByKeyAndWorkflowIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update secret by key and workflow ID params
func (o *UpdateSecretByKeyAndWorkflowIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update secret by key and workflow ID params
func (o *UpdateSecretByKeyAndWorkflowIDParams) WithHTTPClient(client *http.Client) *UpdateSecretByKeyAndWorkflowIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update secret by key and workflow ID params
func (o *UpdateSecretByKeyAndWorkflowIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccept adds the accept to the update secret by key and workflow ID params
func (o *UpdateSecretByKeyAndWorkflowIDParams) WithAccept(accept string) *UpdateSecretByKeyAndWorkflowIDParams {
	o.SetAccept(accept)
	return o
}

// SetAccept adds the accept to the update secret by key and workflow ID params
func (o *UpdateSecretByKeyAndWorkflowIDParams) SetAccept(accept string) {
	o.Accept = accept
}

// WithBody adds the body to the update secret by key and workflow ID params
func (o *UpdateSecretByKeyAndWorkflowIDParams) WithBody(body *models.UpdateSecret) *UpdateSecretByKeyAndWorkflowIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update secret by key and workflow ID params
func (o *UpdateSecretByKeyAndWorkflowIDParams) SetBody(body *models.UpdateSecret) {
	o.Body = body
}

// WithSecretKey adds the secretKey to the update secret by key and workflow ID params
func (o *UpdateSecretByKeyAndWorkflowIDParams) WithSecretKey(secretKey string) *UpdateSecretByKeyAndWorkflowIDParams {
	o.SetSecretKey(secretKey)
	return o
}

// SetSecretKey adds the secretKey to the update secret by key and workflow ID params
func (o *UpdateSecretByKeyAndWorkflowIDParams) SetSecretKey(secretKey string) {
	o.SecretKey = secretKey
}

// WithWorkflowName adds the workflowName to the update secret by key and workflow ID params
func (o *UpdateSecretByKeyAndWorkflowIDParams) WithWorkflowName(workflowName string) *UpdateSecretByKeyAndWorkflowIDParams {
	o.SetWorkflowName(workflowName)
	return o
}

// SetWorkflowName adds the workflowName to the update secret by key and workflow ID params
func (o *UpdateSecretByKeyAndWorkflowIDParams) SetWorkflowName(workflowName string) {
	o.WorkflowName = workflowName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSecretByKeyAndWorkflowIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Accept
	if err := r.SetHeaderParam("Accept", o.Accept); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param secret_key
	if err := r.SetPathParam("secret_key", o.SecretKey); err != nil {
		return err
	}

	// path param workflow_name
	if err := r.SetPathParam("workflow_name", o.WorkflowName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}