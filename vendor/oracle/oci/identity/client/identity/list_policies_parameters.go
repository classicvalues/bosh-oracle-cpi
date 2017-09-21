package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListPoliciesParams creates a new ListPoliciesParams object
// with the default values initialized.
func NewListPoliciesParams() *ListPoliciesParams {
	var ()
	return &ListPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListPoliciesParamsWithTimeout creates a new ListPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListPoliciesParamsWithTimeout(timeout time.Duration) *ListPoliciesParams {
	var ()
	return &ListPoliciesParams{

		timeout: timeout,
	}
}

// NewListPoliciesParamsWithContext creates a new ListPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListPoliciesParamsWithContext(ctx context.Context) *ListPoliciesParams {
	var ()
	return &ListPoliciesParams{

		Context: ctx,
	}
}

// NewListPoliciesParamsWithHTTPClient creates a new ListPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListPoliciesParamsWithHTTPClient(client *http.Client) *ListPoliciesParams {
	var ()
	return &ListPoliciesParams{
		HTTPClient: client,
	}
}

/*ListPoliciesParams contains all the parameters to send to the API endpoint
for the list policies operation typically these are written to a http.Request
*/
type ListPoliciesParams struct {

	/*CompartmentID
	  The OCID of the compartment (remember that the tenancy is simply the root compartment).


	*/
	CompartmentID string
	/*Limit
	  The maximum number of items to return in a paginated "List" call.


	*/
	Limit *int64
	/*Page
	  The value of the `opc-next-page` response header from the previous "List" call.


	*/
	Page *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list policies params
func (o *ListPoliciesParams) WithTimeout(timeout time.Duration) *ListPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list policies params
func (o *ListPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list policies params
func (o *ListPoliciesParams) WithContext(ctx context.Context) *ListPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list policies params
func (o *ListPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list policies params
func (o *ListPoliciesParams) WithHTTPClient(client *http.Client) *ListPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list policies params
func (o *ListPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompartmentID adds the compartmentID to the list policies params
func (o *ListPoliciesParams) WithCompartmentID(compartmentID string) *ListPoliciesParams {
	o.SetCompartmentID(compartmentID)
	return o
}

// SetCompartmentID adds the compartmentId to the list policies params
func (o *ListPoliciesParams) SetCompartmentID(compartmentID string) {
	o.CompartmentID = compartmentID
}

// WithLimit adds the limit to the list policies params
func (o *ListPoliciesParams) WithLimit(limit *int64) *ListPoliciesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list policies params
func (o *ListPoliciesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the list policies params
func (o *ListPoliciesParams) WithPage(page *string) *ListPoliciesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list policies params
func (o *ListPoliciesParams) SetPage(page *string) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *ListPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param compartmentId
	qrCompartmentID := o.CompartmentID
	qCompartmentID := qrCompartmentID
	if qCompartmentID != "" {
		if err := r.SetQueryParam("compartmentId", qCompartmentID); err != nil {
			return err
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage string
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := qrPage
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}