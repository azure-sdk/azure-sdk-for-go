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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
	"net/http"
	"net/url"
	"regexp"
)

// VirtualMachineSizesServer is a fake server for instances of the armmachinelearning.VirtualMachineSizesClient type.
type VirtualMachineSizesServer struct {
	// List is the fake for method VirtualMachineSizesClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, location string, options *armmachinelearning.VirtualMachineSizesClientListOptions) (resp azfake.Responder[armmachinelearning.VirtualMachineSizesClientListResponse], errResp azfake.ErrorResponder)
}

// NewVirtualMachineSizesServerTransport creates a new instance of VirtualMachineSizesServerTransport with the provided implementation.
// The returned VirtualMachineSizesServerTransport instance is connected to an instance of armmachinelearning.VirtualMachineSizesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualMachineSizesServerTransport(srv *VirtualMachineSizesServer) *VirtualMachineSizesServerTransport {
	return &VirtualMachineSizesServerTransport{srv: srv}
}

// VirtualMachineSizesServerTransport connects instances of armmachinelearning.VirtualMachineSizesClient to instances of VirtualMachineSizesServer.
// Don't use this type directly, use NewVirtualMachineSizesServerTransport instead.
type VirtualMachineSizesServerTransport struct {
	srv *VirtualMachineSizesServer
}

// Do implements the policy.Transporter interface for VirtualMachineSizesServerTransport.
func (v *VirtualMachineSizesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualMachineSizesClient.List":
		resp, err = v.dispatchList(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualMachineSizesServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if v.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/vmSizes`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.List(req.Context(), locationParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineSizeListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
