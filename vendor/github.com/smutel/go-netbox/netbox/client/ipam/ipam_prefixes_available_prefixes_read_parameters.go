// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2020 Samuel Mutel <12967891+smutel@users.noreply.github.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
//

package ipam

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

// NewIpamPrefixesAvailablePrefixesReadParams creates a new IpamPrefixesAvailablePrefixesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamPrefixesAvailablePrefixesReadParams() *IpamPrefixesAvailablePrefixesReadParams {
	return &IpamPrefixesAvailablePrefixesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamPrefixesAvailablePrefixesReadParamsWithTimeout creates a new IpamPrefixesAvailablePrefixesReadParams object
// with the ability to set a timeout on a request.
func NewIpamPrefixesAvailablePrefixesReadParamsWithTimeout(timeout time.Duration) *IpamPrefixesAvailablePrefixesReadParams {
	return &IpamPrefixesAvailablePrefixesReadParams{
		timeout: timeout,
	}
}

// NewIpamPrefixesAvailablePrefixesReadParamsWithContext creates a new IpamPrefixesAvailablePrefixesReadParams object
// with the ability to set a context for a request.
func NewIpamPrefixesAvailablePrefixesReadParamsWithContext(ctx context.Context) *IpamPrefixesAvailablePrefixesReadParams {
	return &IpamPrefixesAvailablePrefixesReadParams{
		Context: ctx,
	}
}

// NewIpamPrefixesAvailablePrefixesReadParamsWithHTTPClient creates a new IpamPrefixesAvailablePrefixesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamPrefixesAvailablePrefixesReadParamsWithHTTPClient(client *http.Client) *IpamPrefixesAvailablePrefixesReadParams {
	return &IpamPrefixesAvailablePrefixesReadParams{
		HTTPClient: client,
	}
}

/* IpamPrefixesAvailablePrefixesReadParams contains all the parameters to send to the API endpoint
   for the ipam prefixes available prefixes read operation.

   Typically these are written to a http.Request.
*/
type IpamPrefixesAvailablePrefixesReadParams struct {

	/* ID.

	   A unique integer value identifying this prefix.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam prefixes available prefixes read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesAvailablePrefixesReadParams) WithDefaults() *IpamPrefixesAvailablePrefixesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam prefixes available prefixes read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesAvailablePrefixesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam prefixes available prefixes read params
func (o *IpamPrefixesAvailablePrefixesReadParams) WithTimeout(timeout time.Duration) *IpamPrefixesAvailablePrefixesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes available prefixes read params
func (o *IpamPrefixesAvailablePrefixesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes available prefixes read params
func (o *IpamPrefixesAvailablePrefixesReadParams) WithContext(ctx context.Context) *IpamPrefixesAvailablePrefixesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes available prefixes read params
func (o *IpamPrefixesAvailablePrefixesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes available prefixes read params
func (o *IpamPrefixesAvailablePrefixesReadParams) WithHTTPClient(client *http.Client) *IpamPrefixesAvailablePrefixesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes available prefixes read params
func (o *IpamPrefixesAvailablePrefixesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam prefixes available prefixes read params
func (o *IpamPrefixesAvailablePrefixesReadParams) WithID(id int64) *IpamPrefixesAvailablePrefixesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam prefixes available prefixes read params
func (o *IpamPrefixesAvailablePrefixesReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamPrefixesAvailablePrefixesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}