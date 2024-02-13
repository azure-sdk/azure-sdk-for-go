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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql/v2"
	"net/http"
	"regexp"
)

// CheckNameAvailabilityServer is a fake server for instances of the armpostgresql.CheckNameAvailabilityClient type.
type CheckNameAvailabilityServer struct {
	// Execute is the fake for method CheckNameAvailabilityClient.Execute
	// HTTP status codes to indicate success: http.StatusOK
	Execute func(ctx context.Context, nameAvailabilityRequest armpostgresql.CheckNameAvailabilityRequest, options *armpostgresql.CheckNameAvailabilityClientExecuteOptions) (resp azfake.Responder[armpostgresql.CheckNameAvailabilityClientExecuteResponse], errResp azfake.ErrorResponder)
}

// NewCheckNameAvailabilityServerTransport creates a new instance of CheckNameAvailabilityServerTransport with the provided implementation.
// The returned CheckNameAvailabilityServerTransport instance is connected to an instance of armpostgresql.CheckNameAvailabilityClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCheckNameAvailabilityServerTransport(srv *CheckNameAvailabilityServer) *CheckNameAvailabilityServerTransport {
	return &CheckNameAvailabilityServerTransport{srv: srv}
}

// CheckNameAvailabilityServerTransport connects instances of armpostgresql.CheckNameAvailabilityClient to instances of CheckNameAvailabilityServer.
// Don't use this type directly, use NewCheckNameAvailabilityServerTransport instead.
type CheckNameAvailabilityServerTransport struct {
	srv *CheckNameAvailabilityServer
}

// Do implements the policy.Transporter interface for CheckNameAvailabilityServerTransport.
func (c *CheckNameAvailabilityServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CheckNameAvailabilityClient.Execute":
		resp, err = c.dispatchExecute(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CheckNameAvailabilityServerTransport) dispatchExecute(req *http.Request) (*http.Response, error) {
	if c.srv.Execute == nil {
		return nil, &nonRetriableError{errors.New("fake for method Execute not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/checkNameAvailability`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armpostgresql.CheckNameAvailabilityRequest](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Execute(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NameAvailability, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
