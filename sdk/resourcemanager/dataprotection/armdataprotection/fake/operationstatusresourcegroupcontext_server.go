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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v2"
	"net/http"
	"net/url"
	"regexp"
)

// OperationStatusResourceGroupContextServer is a fake server for instances of the armdataprotection.OperationStatusResourceGroupContextClient type.
type OperationStatusResourceGroupContextServer struct {
	// Get is the fake for method OperationStatusResourceGroupContextClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, operationID string, options *armdataprotection.OperationStatusResourceGroupContextClientGetOptions) (resp azfake.Responder[armdataprotection.OperationStatusResourceGroupContextClientGetResponse], errResp azfake.ErrorResponder)
}

// NewOperationStatusResourceGroupContextServerTransport creates a new instance of OperationStatusResourceGroupContextServerTransport with the provided implementation.
// The returned OperationStatusResourceGroupContextServerTransport instance is connected to an instance of armdataprotection.OperationStatusResourceGroupContextClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOperationStatusResourceGroupContextServerTransport(srv *OperationStatusResourceGroupContextServer) *OperationStatusResourceGroupContextServerTransport {
	return &OperationStatusResourceGroupContextServerTransport{srv: srv}
}

// OperationStatusResourceGroupContextServerTransport connects instances of armdataprotection.OperationStatusResourceGroupContextClient to instances of OperationStatusResourceGroupContextServer.
// Don't use this type directly, use NewOperationStatusResourceGroupContextServerTransport instead.
type OperationStatusResourceGroupContextServerTransport struct {
	srv *OperationStatusResourceGroupContextServer
}

// Do implements the policy.Transporter interface for OperationStatusResourceGroupContextServerTransport.
func (o *OperationStatusResourceGroupContextServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "OperationStatusResourceGroupContextClient.Get":
		resp, err = o.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (o *OperationStatusResourceGroupContextServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if o.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataProtection/operationStatus/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.Get(req.Context(), resourceGroupNameParam, operationIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OperationResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
