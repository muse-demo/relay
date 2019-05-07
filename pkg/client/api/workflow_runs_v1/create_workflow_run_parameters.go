// Code generated by go-swagger; DO NOT EDIT.

package workflow_runs_v1

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

// NewCreateWorkflowRunParams creates a new CreateWorkflowRunParams object
// with the default values initialized.
func NewCreateWorkflowRunParams() *CreateWorkflowRunParams {
	var ()
	return &CreateWorkflowRunParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateWorkflowRunParamsWithTimeout creates a new CreateWorkflowRunParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateWorkflowRunParamsWithTimeout(timeout time.Duration) *CreateWorkflowRunParams {
	var ()
	return &CreateWorkflowRunParams{

		timeout: timeout,
	}
}

// NewCreateWorkflowRunParamsWithContext creates a new CreateWorkflowRunParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateWorkflowRunParamsWithContext(ctx context.Context) *CreateWorkflowRunParams {
	var ()
	return &CreateWorkflowRunParams{

		Context: ctx,
	}
}

// NewCreateWorkflowRunParamsWithHTTPClient creates a new CreateWorkflowRunParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateWorkflowRunParamsWithHTTPClient(client *http.Client) *CreateWorkflowRunParams {
	var ()
	return &CreateWorkflowRunParams{
		HTTPClient: client,
	}
}

/*CreateWorkflowRunParams contains all the parameters to send to the API endpoint
for the create workflow run operation typically these are written to a http.Request
*/
type CreateWorkflowRunParams struct {

	/*Accept
	  The version of the API, in this case should be "application/nebula-api.v1+json"

	*/
	Accept string
	/*Authorization
	  The JWT bearer token

	*/
	Authorization string
	/*Body
	  Workflow to create

	*/
	Body *models.CreateWorkflowRunSubmission
	/*ID
	  ID of the workflow whose runs we want to view

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create workflow run params
func (o *CreateWorkflowRunParams) WithTimeout(timeout time.Duration) *CreateWorkflowRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create workflow run params
func (o *CreateWorkflowRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create workflow run params
func (o *CreateWorkflowRunParams) WithContext(ctx context.Context) *CreateWorkflowRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create workflow run params
func (o *CreateWorkflowRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create workflow run params
func (o *CreateWorkflowRunParams) WithHTTPClient(client *http.Client) *CreateWorkflowRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create workflow run params
func (o *CreateWorkflowRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccept adds the accept to the create workflow run params
func (o *CreateWorkflowRunParams) WithAccept(accept string) *CreateWorkflowRunParams {
	o.SetAccept(accept)
	return o
}

// SetAccept adds the accept to the create workflow run params
func (o *CreateWorkflowRunParams) SetAccept(accept string) {
	o.Accept = accept
}

// WithAuthorization adds the authorization to the create workflow run params
func (o *CreateWorkflowRunParams) WithAuthorization(authorization string) *CreateWorkflowRunParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the create workflow run params
func (o *CreateWorkflowRunParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the create workflow run params
func (o *CreateWorkflowRunParams) WithBody(body *models.CreateWorkflowRunSubmission) *CreateWorkflowRunParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create workflow run params
func (o *CreateWorkflowRunParams) SetBody(body *models.CreateWorkflowRunSubmission) {
	o.Body = body
}

// WithID adds the id to the create workflow run params
func (o *CreateWorkflowRunParams) WithID(id string) *CreateWorkflowRunParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create workflow run params
func (o *CreateWorkflowRunParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateWorkflowRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Accept
	if err := r.SetHeaderParam("Accept", o.Accept); err != nil {
		return err
	}

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
