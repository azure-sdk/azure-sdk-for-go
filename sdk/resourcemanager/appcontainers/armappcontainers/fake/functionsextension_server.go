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
	"net/http"
	"net/url"
	"regexp"

	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// FunctionsExtensionServer is a fake server for instances of the armappcontainers.FunctionsExtensionClient type.
type FunctionsExtensionServer struct {
	// InvokeFunctionsHost is the fake for method FunctionsExtensionClient.InvokeFunctionsHost
	// HTTP status codes to indicate success: http.StatusOK
	InvokeFunctionsHost func(ctx context.Context, resourceGroupName string, containerAppName string, revisionName string, functionAppName string, options *armappcontainers.FunctionsExtensionClientInvokeFunctionsHostOptions) (resp azfake.Responder[armappcontainers.FunctionsExtensionClientInvokeFunctionsHostResponse], errResp azfake.ErrorResponder)
}

// NewFunctionsExtensionServerTransport creates a new instance of FunctionsExtensionServerTransport with the provided implementation.
// The returned FunctionsExtensionServerTransport instance is connected to an instance of armappcontainers.FunctionsExtensionClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFunctionsExtensionServerTransport(srv *FunctionsExtensionServer) *FunctionsExtensionServerTransport {
	return &FunctionsExtensionServerTransport{srv: srv}
}

// FunctionsExtensionServerTransport connects instances of armappcontainers.FunctionsExtensionClient to instances of FunctionsExtensionServer.
// Don't use this type directly, use NewFunctionsExtensionServerTransport instead.
type FunctionsExtensionServerTransport struct {
	srv *FunctionsExtensionServer
}

// Do implements the policy.Transporter interface for FunctionsExtensionServerTransport.
func (f *FunctionsExtensionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FunctionsExtensionClient.InvokeFunctionsHost":
		resp, err = f.dispatchInvokeFunctionsHost(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FunctionsExtensionServerTransport) dispatchInvokeFunctionsHost(req *http.Request) (*http.Response, error) {
	if f.srv.InvokeFunctionsHost == nil {
		return nil, &nonRetriableError{errors.New("fake for method InvokeFunctionsHost not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/revisions/(?P<revisionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/functions/(?P<functionAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/invoke`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
	if err != nil {
		return nil, err
	}
	revisionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("revisionName")])
	if err != nil {
		return nil, err
	}
	functionAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("functionAppName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.InvokeFunctionsHost(req.Context(), resourceGroupNameParam, containerAppNameParam, revisionNameParam, functionAppNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
