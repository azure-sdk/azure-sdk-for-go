// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/loadtesting/armloadtesting/v2"
	"net/http"
)

// OperationsServer is a fake server for instances of the armloadtesting.OperationsClient type.
type OperationsServer struct {
	// NewListPager is the fake for method OperationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armloadtesting.OperationsClientListOptions) (resp azfake.PagerResponder[armloadtesting.OperationsClientListResponse])
}

// NewOperationsServerTransport creates a new instance of OperationsServerTransport with the provided implementation.
// The returned OperationsServerTransport instance is connected to an instance of armloadtesting.OperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOperationsServerTransport(srv *OperationsServer) *OperationsServerTransport {
	return &OperationsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armloadtesting.OperationsClientListResponse]](),
	}
}

// OperationsServerTransport connects instances of armloadtesting.OperationsClient to instances of OperationsServer.
// Don't use this type directly, use NewOperationsServerTransport instead.
type OperationsServerTransport struct {
	srv          *OperationsServer
	newListPager *tracker[azfake.PagerResponder[armloadtesting.OperationsClientListResponse]]
}

// Do implements the policy.Transporter interface for OperationsServerTransport.
func (o *OperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return o.dispatchToMethodFake(req, method)
}

func (o *OperationsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if operationsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = operationsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "OperationsClient.NewListPager":
				res.resp, res.err = o.dispatchNewListPager(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (o *OperationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if o.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := o.newListPager.get(req)
	if newListPager == nil {
		resp := o.srv.NewListPager(nil)
		newListPager = &resp
		o.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armloadtesting.OperationsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		o.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		o.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to OperationsServerTransport
var operationsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
