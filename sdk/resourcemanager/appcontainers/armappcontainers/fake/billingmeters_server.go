//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
	"net/http"
	"net/url"
	"regexp"
)

// BillingMetersServer is a fake server for instances of the armappcontainers.BillingMetersClient type.
type BillingMetersServer struct {
	// Get is the fake for method BillingMetersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, options *armappcontainers.BillingMetersClientGetOptions) (resp azfake.Responder[armappcontainers.BillingMetersClientGetResponse], errResp azfake.ErrorResponder)
}

// NewBillingMetersServerTransport creates a new instance of BillingMetersServerTransport with the provided implementation.
// The returned BillingMetersServerTransport instance is connected to an instance of armappcontainers.BillingMetersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBillingMetersServerTransport(srv *BillingMetersServer) *BillingMetersServerTransport {
	return &BillingMetersServerTransport{srv: srv}
}

// BillingMetersServerTransport connects instances of armappcontainers.BillingMetersClient to instances of BillingMetersServer.
// Don't use this type directly, use NewBillingMetersServerTransport instead.
type BillingMetersServerTransport struct {
	srv *BillingMetersServer
}

// Do implements the policy.Transporter interface for BillingMetersServerTransport.
func (b *BillingMetersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "BillingMetersClient.Get":
		resp, err = b.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BillingMetersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/billingMeters`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Get(req.Context(), locationParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BillingMeterCollection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
