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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scvmm/armscvmm/v2"
	"net/http"
	"net/url"
	"regexp"
)

// GuestAgentsServer is a fake server for instances of the armscvmm.GuestAgentsClient type.
type GuestAgentsServer struct {
	// BeginCreate is the fake for method GuestAgentsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceURI string, resource armscvmm.GuestAgent, options *armscvmm.GuestAgentsClientBeginCreateOptions) (resp azfake.PollerResponder[armscvmm.GuestAgentsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method GuestAgentsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceURI string, options *armscvmm.GuestAgentsClientDeleteOptions) (resp azfake.Responder[armscvmm.GuestAgentsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method GuestAgentsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceURI string, options *armscvmm.GuestAgentsClientGetOptions) (resp azfake.Responder[armscvmm.GuestAgentsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByVirtualMachineInstancePager is the fake for method GuestAgentsClient.NewListByVirtualMachineInstancePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByVirtualMachineInstancePager func(resourceURI string, options *armscvmm.GuestAgentsClientListByVirtualMachineInstanceOptions) (resp azfake.PagerResponder[armscvmm.GuestAgentsClientListByVirtualMachineInstanceResponse])
}

// NewGuestAgentsServerTransport creates a new instance of GuestAgentsServerTransport with the provided implementation.
// The returned GuestAgentsServerTransport instance is connected to an instance of armscvmm.GuestAgentsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGuestAgentsServerTransport(srv *GuestAgentsServer) *GuestAgentsServerTransport {
	return &GuestAgentsServerTransport{
		srv:                                  srv,
		beginCreate:                          newTracker[azfake.PollerResponder[armscvmm.GuestAgentsClientCreateResponse]](),
		newListByVirtualMachineInstancePager: newTracker[azfake.PagerResponder[armscvmm.GuestAgentsClientListByVirtualMachineInstanceResponse]](),
	}
}

// GuestAgentsServerTransport connects instances of armscvmm.GuestAgentsClient to instances of GuestAgentsServer.
// Don't use this type directly, use NewGuestAgentsServerTransport instead.
type GuestAgentsServerTransport struct {
	srv                                  *GuestAgentsServer
	beginCreate                          *tracker[azfake.PollerResponder[armscvmm.GuestAgentsClientCreateResponse]]
	newListByVirtualMachineInstancePager *tracker[azfake.PagerResponder[armscvmm.GuestAgentsClientListByVirtualMachineInstanceResponse]]
}

// Do implements the policy.Transporter interface for GuestAgentsServerTransport.
func (g *GuestAgentsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return g.dispatchToMethodFake(req, method)
}

func (g *GuestAgentsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if guestAgentsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = guestAgentsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "GuestAgentsClient.BeginCreate":
				res.resp, res.err = g.dispatchBeginCreate(req)
			case "GuestAgentsClient.Delete":
				res.resp, res.err = g.dispatchDelete(req)
			case "GuestAgentsClient.Get":
				res.resp, res.err = g.dispatchGet(req)
			case "GuestAgentsClient.NewListByVirtualMachineInstancePager":
				res.resp, res.err = g.dispatchNewListByVirtualMachineInstancePager(req)
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

func (g *GuestAgentsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if g.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := g.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ScVmm/virtualMachineInstances/default/guestAgents/default`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armscvmm.GuestAgent](req)
		if err != nil {
			return nil, err
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginCreate(req.Context(), resourceURIParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		g.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		g.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		g.beginCreate.remove(req)
	}

	return resp, nil
}

func (g *GuestAgentsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if g.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ScVmm/virtualMachineInstances/default/guestAgents/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Delete(req.Context(), resourceURIParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GuestAgentsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if g.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ScVmm/virtualMachineInstances/default/guestAgents/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Get(req.Context(), resourceURIParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GuestAgent, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GuestAgentsServerTransport) dispatchNewListByVirtualMachineInstancePager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListByVirtualMachineInstancePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByVirtualMachineInstancePager not implemented")}
	}
	newListByVirtualMachineInstancePager := g.newListByVirtualMachineInstancePager.get(req)
	if newListByVirtualMachineInstancePager == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ScVmm/virtualMachineInstances/default/guestAgents`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		resp := g.srv.NewListByVirtualMachineInstancePager(resourceURIParam, nil)
		newListByVirtualMachineInstancePager = &resp
		g.newListByVirtualMachineInstancePager.add(req, newListByVirtualMachineInstancePager)
		server.PagerResponderInjectNextLinks(newListByVirtualMachineInstancePager, req, func(page *armscvmm.GuestAgentsClientListByVirtualMachineInstanceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByVirtualMachineInstancePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.newListByVirtualMachineInstancePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByVirtualMachineInstancePager) {
		g.newListByVirtualMachineInstancePager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to GuestAgentsServerTransport
var guestAgentsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
