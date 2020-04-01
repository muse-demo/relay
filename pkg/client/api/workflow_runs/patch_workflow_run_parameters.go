// Code generated by go-swagger; DO NOT EDIT.

package workflow_runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewPatchWorkflowRunParams creates a new PatchWorkflowRunParams object
// with the default values initialized.
func NewPatchWorkflowRunParams() *PatchWorkflowRunParams {
	var ()
	return &PatchWorkflowRunParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchWorkflowRunParamsWithTimeout creates a new PatchWorkflowRunParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchWorkflowRunParamsWithTimeout(timeout time.Duration) *PatchWorkflowRunParams {
	var ()
	return &PatchWorkflowRunParams{

		timeout: timeout,
	}
}

// NewPatchWorkflowRunParamsWithContext creates a new PatchWorkflowRunParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchWorkflowRunParamsWithContext(ctx context.Context) *PatchWorkflowRunParams {
	var ()
	return &PatchWorkflowRunParams{

		Context: ctx,
	}
}

// NewPatchWorkflowRunParamsWithHTTPClient creates a new PatchWorkflowRunParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchWorkflowRunParamsWithHTTPClient(client *http.Client) *PatchWorkflowRunParams {
	var ()
	return &PatchWorkflowRunParams{
		HTTPClient: client,
	}
}

/*PatchWorkflowRunParams contains all the parameters to send to the API endpoint
for the patch workflow run operation typically these are written to a http.Request
*/
type PatchWorkflowRunParams struct {

	/*Body
	  Update properties of workflow run

	*/
	Body PatchWorkflowRunBody
	/*WorkflowName
	  Workflow name

	*/
	WorkflowName string
	/*WorkflowRunNumber
	  Run number of the associated workflow

	*/
	WorkflowRunNumber int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch workflow run params
func (o *PatchWorkflowRunParams) WithTimeout(timeout time.Duration) *PatchWorkflowRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch workflow run params
func (o *PatchWorkflowRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch workflow run params
func (o *PatchWorkflowRunParams) WithContext(ctx context.Context) *PatchWorkflowRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch workflow run params
func (o *PatchWorkflowRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch workflow run params
func (o *PatchWorkflowRunParams) WithHTTPClient(client *http.Client) *PatchWorkflowRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch workflow run params
func (o *PatchWorkflowRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch workflow run params
func (o *PatchWorkflowRunParams) WithBody(body PatchWorkflowRunBody) *PatchWorkflowRunParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch workflow run params
func (o *PatchWorkflowRunParams) SetBody(body PatchWorkflowRunBody) {
	o.Body = body
}

// WithWorkflowName adds the workflowName to the patch workflow run params
func (o *PatchWorkflowRunParams) WithWorkflowName(workflowName string) *PatchWorkflowRunParams {
	o.SetWorkflowName(workflowName)
	return o
}

// SetWorkflowName adds the workflowName to the patch workflow run params
func (o *PatchWorkflowRunParams) SetWorkflowName(workflowName string) {
	o.WorkflowName = workflowName
}

// WithWorkflowRunNumber adds the workflowRunNumber to the patch workflow run params
func (o *PatchWorkflowRunParams) WithWorkflowRunNumber(workflowRunNumber int64) *PatchWorkflowRunParams {
	o.SetWorkflowRunNumber(workflowRunNumber)
	return o
}

// SetWorkflowRunNumber adds the workflowRunNumber to the patch workflow run params
func (o *PatchWorkflowRunParams) SetWorkflowRunNumber(workflowRunNumber int64) {
	o.WorkflowRunNumber = workflowRunNumber
}

// WriteToRequest writes these params to a swagger request
func (o *PatchWorkflowRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param workflowName
	if err := r.SetPathParam("workflowName", o.WorkflowName); err != nil {
		return err
	}

	// path param workflowRunNumber
	if err := r.SetPathParam("workflowRunNumber", swag.FormatInt64(o.WorkflowRunNumber)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}