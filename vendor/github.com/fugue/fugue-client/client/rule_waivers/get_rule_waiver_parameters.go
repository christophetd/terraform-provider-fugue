// Code generated by go-swagger; DO NOT EDIT.

package rule_waivers

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
)

// NewGetRuleWaiverParams creates a new GetRuleWaiverParams object
// with the default values initialized.
func NewGetRuleWaiverParams() *GetRuleWaiverParams {
	var ()
	return &GetRuleWaiverParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRuleWaiverParamsWithTimeout creates a new GetRuleWaiverParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRuleWaiverParamsWithTimeout(timeout time.Duration) *GetRuleWaiverParams {
	var ()
	return &GetRuleWaiverParams{

		timeout: timeout,
	}
}

// NewGetRuleWaiverParamsWithContext creates a new GetRuleWaiverParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRuleWaiverParamsWithContext(ctx context.Context) *GetRuleWaiverParams {
	var ()
	return &GetRuleWaiverParams{

		Context: ctx,
	}
}

// NewGetRuleWaiverParamsWithHTTPClient creates a new GetRuleWaiverParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRuleWaiverParamsWithHTTPClient(client *http.Client) *GetRuleWaiverParams {
	var ()
	return &GetRuleWaiverParams{
		HTTPClient: client,
	}
}

/*GetRuleWaiverParams contains all the parameters to send to the API endpoint
for the get rule waiver operation typically these are written to a http.Request
*/
type GetRuleWaiverParams struct {

	/*RuleWaiverID
	  The ID of rule waiver to update

	*/
	RuleWaiverID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get rule waiver params
func (o *GetRuleWaiverParams) WithTimeout(timeout time.Duration) *GetRuleWaiverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get rule waiver params
func (o *GetRuleWaiverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get rule waiver params
func (o *GetRuleWaiverParams) WithContext(ctx context.Context) *GetRuleWaiverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get rule waiver params
func (o *GetRuleWaiverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get rule waiver params
func (o *GetRuleWaiverParams) WithHTTPClient(client *http.Client) *GetRuleWaiverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get rule waiver params
func (o *GetRuleWaiverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRuleWaiverID adds the ruleWaiverID to the get rule waiver params
func (o *GetRuleWaiverParams) WithRuleWaiverID(ruleWaiverID string) *GetRuleWaiverParams {
	o.SetRuleWaiverID(ruleWaiverID)
	return o
}

// SetRuleWaiverID adds the ruleWaiverId to the get rule waiver params
func (o *GetRuleWaiverParams) SetRuleWaiverID(ruleWaiverID string) {
	o.RuleWaiverID = ruleWaiverID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRuleWaiverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param rule_waiver_id
	if err := r.SetPathParam("rule_waiver_id", o.RuleWaiverID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}