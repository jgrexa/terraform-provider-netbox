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

package wireless

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

	"github.com/smutel/go-netbox/netbox/models"
)

// NewWirelessWirelessLanGroupsCreateParams creates a new WirelessWirelessLanGroupsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWirelessWirelessLanGroupsCreateParams() *WirelessWirelessLanGroupsCreateParams {
	return &WirelessWirelessLanGroupsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWirelessWirelessLanGroupsCreateParamsWithTimeout creates a new WirelessWirelessLanGroupsCreateParams object
// with the ability to set a timeout on a request.
func NewWirelessWirelessLanGroupsCreateParamsWithTimeout(timeout time.Duration) *WirelessWirelessLanGroupsCreateParams {
	return &WirelessWirelessLanGroupsCreateParams{
		timeout: timeout,
	}
}

// NewWirelessWirelessLanGroupsCreateParamsWithContext creates a new WirelessWirelessLanGroupsCreateParams object
// with the ability to set a context for a request.
func NewWirelessWirelessLanGroupsCreateParamsWithContext(ctx context.Context) *WirelessWirelessLanGroupsCreateParams {
	return &WirelessWirelessLanGroupsCreateParams{
		Context: ctx,
	}
}

// NewWirelessWirelessLanGroupsCreateParamsWithHTTPClient creates a new WirelessWirelessLanGroupsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewWirelessWirelessLanGroupsCreateParamsWithHTTPClient(client *http.Client) *WirelessWirelessLanGroupsCreateParams {
	return &WirelessWirelessLanGroupsCreateParams{
		HTTPClient: client,
	}
}

/* WirelessWirelessLanGroupsCreateParams contains all the parameters to send to the API endpoint
   for the wireless wireless lan groups create operation.

   Typically these are written to a http.Request.
*/
type WirelessWirelessLanGroupsCreateParams struct {

	// Data.
	Data *models.WritableWirelessLANGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the wireless wireless lan groups create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WirelessWirelessLanGroupsCreateParams) WithDefaults() *WirelessWirelessLanGroupsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the wireless wireless lan groups create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WirelessWirelessLanGroupsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the wireless wireless lan groups create params
func (o *WirelessWirelessLanGroupsCreateParams) WithTimeout(timeout time.Duration) *WirelessWirelessLanGroupsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the wireless wireless lan groups create params
func (o *WirelessWirelessLanGroupsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the wireless wireless lan groups create params
func (o *WirelessWirelessLanGroupsCreateParams) WithContext(ctx context.Context) *WirelessWirelessLanGroupsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the wireless wireless lan groups create params
func (o *WirelessWirelessLanGroupsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the wireless wireless lan groups create params
func (o *WirelessWirelessLanGroupsCreateParams) WithHTTPClient(client *http.Client) *WirelessWirelessLanGroupsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the wireless wireless lan groups create params
func (o *WirelessWirelessLanGroupsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the wireless wireless lan groups create params
func (o *WirelessWirelessLanGroupsCreateParams) WithData(data *models.WritableWirelessLANGroup) *WirelessWirelessLanGroupsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the wireless wireless lan groups create params
func (o *WirelessWirelessLanGroupsCreateParams) SetData(data *models.WritableWirelessLANGroup) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *WirelessWirelessLanGroupsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
