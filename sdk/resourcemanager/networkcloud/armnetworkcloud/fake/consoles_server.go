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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
	"net/http"
	"net/url"
	"regexp"
)

// ConsolesServer is a fake server for instances of the armnetworkcloud.ConsolesClient type.
type ConsolesServer struct {
	// BeginCreateOrUpdate is the fake for method ConsolesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, virtualMachineName string, consoleName string, consoleParameters armnetworkcloud.Console, options *armnetworkcloud.ConsolesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetworkcloud.ConsolesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ConsolesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, virtualMachineName string, consoleName string, options *armnetworkcloud.ConsolesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetworkcloud.ConsolesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ConsolesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, virtualMachineName string, consoleName string, options *armnetworkcloud.ConsolesClientGetOptions) (resp azfake.Responder[armnetworkcloud.ConsolesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByVirtualMachinePager is the fake for method ConsolesClient.NewListByVirtualMachinePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByVirtualMachinePager func(resourceGroupName string, virtualMachineName string, options *armnetworkcloud.ConsolesClientListByVirtualMachineOptions) (resp azfake.PagerResponder[armnetworkcloud.ConsolesClientListByVirtualMachineResponse])

	// BeginUpdate is the fake for method ConsolesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, virtualMachineName string, consoleName string, consoleUpdateParameters armnetworkcloud.ConsolePatchParameters, options *armnetworkcloud.ConsolesClientBeginUpdateOptions) (resp azfake.PollerResponder[armnetworkcloud.ConsolesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewConsolesServerTransport creates a new instance of ConsolesServerTransport with the provided implementation.
// The returned ConsolesServerTransport instance is connected to an instance of armnetworkcloud.ConsolesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewConsolesServerTransport(srv *ConsolesServer) *ConsolesServerTransport {
	return &ConsolesServerTransport{
		srv:                          srv,
		beginCreateOrUpdate:          newTracker[azfake.PollerResponder[armnetworkcloud.ConsolesClientCreateOrUpdateResponse]](),
		beginDelete:                  newTracker[azfake.PollerResponder[armnetworkcloud.ConsolesClientDeleteResponse]](),
		newListByVirtualMachinePager: newTracker[azfake.PagerResponder[armnetworkcloud.ConsolesClientListByVirtualMachineResponse]](),
		beginUpdate:                  newTracker[azfake.PollerResponder[armnetworkcloud.ConsolesClientUpdateResponse]](),
	}
}

// ConsolesServerTransport connects instances of armnetworkcloud.ConsolesClient to instances of ConsolesServer.
// Don't use this type directly, use NewConsolesServerTransport instead.
type ConsolesServerTransport struct {
	srv                          *ConsolesServer
	beginCreateOrUpdate          *tracker[azfake.PollerResponder[armnetworkcloud.ConsolesClientCreateOrUpdateResponse]]
	beginDelete                  *tracker[azfake.PollerResponder[armnetworkcloud.ConsolesClientDeleteResponse]]
	newListByVirtualMachinePager *tracker[azfake.PagerResponder[armnetworkcloud.ConsolesClientListByVirtualMachineResponse]]
	beginUpdate                  *tracker[azfake.PollerResponder[armnetworkcloud.ConsolesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for ConsolesServerTransport.
func (c *ConsolesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *ConsolesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if consolesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = consolesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ConsolesClient.BeginCreateOrUpdate":
				res.resp, res.err = c.dispatchBeginCreateOrUpdate(req)
			case "ConsolesClient.BeginDelete":
				res.resp, res.err = c.dispatchBeginDelete(req)
			case "ConsolesClient.Get":
				res.resp, res.err = c.dispatchGet(req)
			case "ConsolesClient.NewListByVirtualMachinePager":
				res.resp, res.err = c.dispatchNewListByVirtualMachinePager(req)
			case "ConsolesClient.BeginUpdate":
				res.resp, res.err = c.dispatchBeginUpdate(req)
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

func (c *ConsolesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := c.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/virtualMachines/(?P<virtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/consoles/(?P<consoleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetworkcloud.Console](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualMachineName")])
		if err != nil {
			return nil, err
		}
		consoleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("consoleName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
		var options *armnetworkcloud.ConsolesClientBeginCreateOrUpdateOptions
		if ifMatchParam != nil || ifNoneMatchParam != nil {
			options = &armnetworkcloud.ConsolesClientBeginCreateOrUpdateOptions{
				IfMatch:     ifMatchParam,
				IfNoneMatch: ifNoneMatchParam,
			}
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, virtualMachineNameParam, consoleNameParam, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		c.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		c.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		c.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (c *ConsolesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/virtualMachines/(?P<virtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/consoles/(?P<consoleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualMachineName")])
		if err != nil {
			return nil, err
		}
		consoleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("consoleName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
		var options *armnetworkcloud.ConsolesClientBeginDeleteOptions
		if ifMatchParam != nil || ifNoneMatchParam != nil {
			options = &armnetworkcloud.ConsolesClientBeginDeleteOptions{
				IfMatch:     ifMatchParam,
				IfNoneMatch: ifNoneMatchParam,
			}
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, virtualMachineNameParam, consoleNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		c.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		c.beginDelete.remove(req)
	}

	return resp, nil
}

func (c *ConsolesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/virtualMachines/(?P<virtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/consoles/(?P<consoleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	virtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualMachineName")])
	if err != nil {
		return nil, err
	}
	consoleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("consoleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, virtualMachineNameParam, consoleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Console, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConsolesServerTransport) dispatchNewListByVirtualMachinePager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByVirtualMachinePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByVirtualMachinePager not implemented")}
	}
	newListByVirtualMachinePager := c.newListByVirtualMachinePager.get(req)
	if newListByVirtualMachinePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/virtualMachines/(?P<virtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/consoles`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualMachineName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListByVirtualMachinePager(resourceGroupNameParam, virtualMachineNameParam, nil)
		newListByVirtualMachinePager = &resp
		c.newListByVirtualMachinePager.add(req, newListByVirtualMachinePager)
		server.PagerResponderInjectNextLinks(newListByVirtualMachinePager, req, func(page *armnetworkcloud.ConsolesClientListByVirtualMachineResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByVirtualMachinePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByVirtualMachinePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByVirtualMachinePager) {
		c.newListByVirtualMachinePager.remove(req)
	}
	return resp, nil
}

func (c *ConsolesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := c.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/virtualMachines/(?P<virtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/consoles/(?P<consoleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetworkcloud.ConsolePatchParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualMachineName")])
		if err != nil {
			return nil, err
		}
		consoleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("consoleName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
		var options *armnetworkcloud.ConsolesClientBeginUpdateOptions
		if ifMatchParam != nil || ifNoneMatchParam != nil {
			options = &armnetworkcloud.ConsolesClientBeginUpdateOptions{
				IfMatch:     ifMatchParam,
				IfNoneMatch: ifNoneMatchParam,
			}
		}
		respr, errRespr := c.srv.BeginUpdate(req.Context(), resourceGroupNameParam, virtualMachineNameParam, consoleNameParam, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		c.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		c.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ConsolesServerTransport
var consolesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
