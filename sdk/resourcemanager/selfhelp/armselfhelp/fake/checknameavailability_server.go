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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/selfhelp/armselfhelp/v2"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// CheckNameAvailabilityServer is a fake server for instances of the armselfhelp.CheckNameAvailabilityClient type.
type CheckNameAvailabilityServer struct {
	// Post is the fake for method CheckNameAvailabilityClient.Post
	// HTTP status codes to indicate success: http.StatusOK
	Post func(ctx context.Context, scope string, options *armselfhelp.CheckNameAvailabilityClientPostOptions) (resp azfake.Responder[armselfhelp.CheckNameAvailabilityClientPostResponse], errResp azfake.ErrorResponder)
}

// NewCheckNameAvailabilityServerTransport creates a new instance of CheckNameAvailabilityServerTransport with the provided implementation.
// The returned CheckNameAvailabilityServerTransport instance is connected to an instance of armselfhelp.CheckNameAvailabilityClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCheckNameAvailabilityServerTransport(srv *CheckNameAvailabilityServer) *CheckNameAvailabilityServerTransport {
	return &CheckNameAvailabilityServerTransport{srv: srv}
}

// CheckNameAvailabilityServerTransport connects instances of armselfhelp.CheckNameAvailabilityClient to instances of CheckNameAvailabilityServer.
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
	case "CheckNameAvailabilityClient.Post":
		resp, err = c.dispatchPost(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CheckNameAvailabilityServerTransport) dispatchPost(req *http.Request) (*http.Response, error) {
	if c.srv.Post == nil {
		return nil, &nonRetriableError{errors.New("fake for method Post not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Help/checkNameAvailability`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armselfhelp.CheckNameAvailabilityRequest](req)
	if err != nil {
		return nil, err
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	var options *armselfhelp.CheckNameAvailabilityClientPostOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armselfhelp.CheckNameAvailabilityClientPostOptions{
			CheckNameAvailabilityRequest: &body,
		}
	}
	respr, errRespr := c.srv.Post(req.Context(), scopeParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CheckNameAvailabilityResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
