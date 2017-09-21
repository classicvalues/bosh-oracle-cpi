package virtual_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPrivateIPParams creates a new GetPrivateIPParams object
// with the default values initialized.
func NewGetPrivateIPParams() *GetPrivateIPParams {
	var ()
	return &GetPrivateIPParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateIPParamsWithTimeout creates a new GetPrivateIPParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateIPParamsWithTimeout(timeout time.Duration) *GetPrivateIPParams {
	var ()
	return &GetPrivateIPParams{

		timeout: timeout,
	}
}

// NewGetPrivateIPParamsWithContext creates a new GetPrivateIPParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateIPParamsWithContext(ctx context.Context) *GetPrivateIPParams {
	var ()
	return &GetPrivateIPParams{

		Context: ctx,
	}
}

// NewGetPrivateIPParamsWithHTTPClient creates a new GetPrivateIPParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateIPParamsWithHTTPClient(client *http.Client) *GetPrivateIPParams {
	var ()
	return &GetPrivateIPParams{
		HTTPClient: client,
	}
}

/*GetPrivateIPParams contains all the parameters to send to the API endpoint
for the get private Ip operation typically these are written to a http.Request
*/
type GetPrivateIPParams struct {

	/*PrivateIPID
	  The private IP's OCID.

	*/
	PrivateIPID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private Ip params
func (o *GetPrivateIPParams) WithTimeout(timeout time.Duration) *GetPrivateIPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private Ip params
func (o *GetPrivateIPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private Ip params
func (o *GetPrivateIPParams) WithContext(ctx context.Context) *GetPrivateIPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private Ip params
func (o *GetPrivateIPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private Ip params
func (o *GetPrivateIPParams) WithHTTPClient(client *http.Client) *GetPrivateIPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private Ip params
func (o *GetPrivateIPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPrivateIPID adds the privateIPID to the get private Ip params
func (o *GetPrivateIPParams) WithPrivateIPID(privateIPID string) *GetPrivateIPParams {
	o.SetPrivateIPID(privateIPID)
	return o
}

// SetPrivateIPID adds the privateIpId to the get private Ip params
func (o *GetPrivateIPParams) SetPrivateIPID(privateIPID string) {
	o.PrivateIPID = privateIPID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateIPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param privateIpId
	if err := r.SetPathParam("privateIpId", o.PrivateIPID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}