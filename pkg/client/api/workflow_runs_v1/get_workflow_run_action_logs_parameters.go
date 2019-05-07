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
)

// NewGetWorkflowRunActionLogsParams creates a new GetWorkflowRunActionLogsParams object
// with the default values initialized.
func NewGetWorkflowRunActionLogsParams() *GetWorkflowRunActionLogsParams {
	var ()
	return &GetWorkflowRunActionLogsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkflowRunActionLogsParamsWithTimeout creates a new GetWorkflowRunActionLogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkflowRunActionLogsParamsWithTimeout(timeout time.Duration) *GetWorkflowRunActionLogsParams {
	var ()
	return &GetWorkflowRunActionLogsParams{

		timeout: timeout,
	}
}

// NewGetWorkflowRunActionLogsParamsWithContext creates a new GetWorkflowRunActionLogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkflowRunActionLogsParamsWithContext(ctx context.Context) *GetWorkflowRunActionLogsParams {
	var ()
	return &GetWorkflowRunActionLogsParams{

		Context: ctx,
	}
}

// NewGetWorkflowRunActionLogsParamsWithHTTPClient creates a new GetWorkflowRunActionLogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkflowRunActionLogsParamsWithHTTPClient(client *http.Client) *GetWorkflowRunActionLogsParams {
	var ()
	return &GetWorkflowRunActionLogsParams{
		HTTPClient: client,
	}
}

/*GetWorkflowRunActionLogsParams contains all the parameters to send to the API endpoint
for the get workflow run action logs operation typically these are written to a http.Request
*/
type GetWorkflowRunActionLogsParams struct {

	/*Accept
	  The version of the API, in this case should be "application/nebula-api.v1+json"

	*/
	Accept string
	/*Authorization
	  The JWT bearer token

	*/
	Authorization string
	/*ActionName
	  Unique workflow action name

	*/
	ActionName string
	/*Rid
	  ID of the workflow run we want to know about

	*/
	Rid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workflow run action logs params
func (o *GetWorkflowRunActionLogsParams) WithTimeout(timeout time.Duration) *GetWorkflowRunActionLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workflow run action logs params
func (o *GetWorkflowRunActionLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workflow run action logs params
func (o *GetWorkflowRunActionLogsParams) WithContext(ctx context.Context) *GetWorkflowRunActionLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workflow run action logs params
func (o *GetWorkflowRunActionLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workflow run action logs params
func (o *GetWorkflowRunActionLogsParams) WithHTTPClient(client *http.Client) *GetWorkflowRunActionLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workflow run action logs params
func (o *GetWorkflowRunActionLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccept adds the accept to the get workflow run action logs params
func (o *GetWorkflowRunActionLogsParams) WithAccept(accept string) *GetWorkflowRunActionLogsParams {
	o.SetAccept(accept)
	return o
}

// SetAccept adds the accept to the get workflow run action logs params
func (o *GetWorkflowRunActionLogsParams) SetAccept(accept string) {
	o.Accept = accept
}

// WithAuthorization adds the authorization to the get workflow run action logs params
func (o *GetWorkflowRunActionLogsParams) WithAuthorization(authorization string) *GetWorkflowRunActionLogsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get workflow run action logs params
func (o *GetWorkflowRunActionLogsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithActionName adds the actionName to the get workflow run action logs params
func (o *GetWorkflowRunActionLogsParams) WithActionName(actionName string) *GetWorkflowRunActionLogsParams {
	o.SetActionName(actionName)
	return o
}

// SetActionName adds the actionName to the get workflow run action logs params
func (o *GetWorkflowRunActionLogsParams) SetActionName(actionName string) {
	o.ActionName = actionName
}

// WithRid adds the rid to the get workflow run action logs params
func (o *GetWorkflowRunActionLogsParams) WithRid(rid string) *GetWorkflowRunActionLogsParams {
	o.SetRid(rid)
	return o
}

// SetRid adds the rid to the get workflow run action logs params
func (o *GetWorkflowRunActionLogsParams) SetRid(rid string) {
	o.Rid = rid
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkflowRunActionLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param action_name
	if err := r.SetPathParam("action_name", o.ActionName); err != nil {
		return err
	}

	// path param rid
	if err := r.SetPathParam("rid", o.Rid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
