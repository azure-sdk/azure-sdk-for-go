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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
	"net/http"
	"regexp"
)

// AvailableSKUsServer is a fake server for instances of the armdataboxedge.AvailableSKUsClient type.
type AvailableSKUsServer struct {
	// NewListPager is the fake for method AvailableSKUsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armdataboxedge.AvailableSKUsClientListOptions) (resp azfake.PagerResponder[armdataboxedge.AvailableSKUsClientListResponse])
}

// NewAvailableSKUsServerTransport creates a new instance of AvailableSKUsServerTransport with the provided implementation.
// The returned AvailableSKUsServerTransport instance is connected to an instance of armdataboxedge.AvailableSKUsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAvailableSKUsServerTransport(srv *AvailableSKUsServer) *AvailableSKUsServerTransport {
	return &AvailableSKUsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armdataboxedge.AvailableSKUsClientListResponse]](),
	}
}

// AvailableSKUsServerTransport connects instances of armdataboxedge.AvailableSKUsClient to instances of AvailableSKUsServer.
// Don't use this type directly, use NewAvailableSKUsServerTransport instead.
type AvailableSKUsServerTransport struct {
	srv          *AvailableSKUsServer
	newListPager *tracker[azfake.PagerResponder[armdataboxedge.AvailableSKUsClientListResponse]]
}

// Do implements the policy.Transporter interface for AvailableSKUsServerTransport.
func (a *AvailableSKUsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *AvailableSKUsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if availableSkUsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = availableSkUsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "AvailableSKUsClient.NewListPager":
				res.resp, res.err = a.dispatchNewListPager(req)
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

func (a *AvailableSKUsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataBoxEdge/availableSkus`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListPager(nil)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armdataboxedge.AvailableSKUsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		a.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to AvailableSKUsServerTransport
var availableSkUsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
