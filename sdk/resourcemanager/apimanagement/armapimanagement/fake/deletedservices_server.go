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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
	"net/http"
	"net/url"
	"regexp"
)

// DeletedServicesServer is a fake server for instances of the armapimanagement.DeletedServicesClient type.
type DeletedServicesServer struct {
	// GetByName is the fake for method DeletedServicesClient.GetByName
	// HTTP status codes to indicate success: http.StatusOK
	GetByName func(ctx context.Context, serviceName string, location string, options *armapimanagement.DeletedServicesClientGetByNameOptions) (resp azfake.Responder[armapimanagement.DeletedServicesClientGetByNameResponse], errResp azfake.ErrorResponder)

	// NewListBySubscriptionPager is the fake for method DeletedServicesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armapimanagement.DeletedServicesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armapimanagement.DeletedServicesClientListBySubscriptionResponse])

	// BeginPurge is the fake for method DeletedServicesClient.BeginPurge
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginPurge func(ctx context.Context, serviceName string, location string, options *armapimanagement.DeletedServicesClientBeginPurgeOptions) (resp azfake.PollerResponder[armapimanagement.DeletedServicesClientPurgeResponse], errResp azfake.ErrorResponder)
}

// NewDeletedServicesServerTransport creates a new instance of DeletedServicesServerTransport with the provided implementation.
// The returned DeletedServicesServerTransport instance is connected to an instance of armapimanagement.DeletedServicesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDeletedServicesServerTransport(srv *DeletedServicesServer) *DeletedServicesServerTransport {
	return &DeletedServicesServerTransport{
		srv:                        srv,
		newListBySubscriptionPager: newTracker[azfake.PagerResponder[armapimanagement.DeletedServicesClientListBySubscriptionResponse]](),
		beginPurge:                 newTracker[azfake.PollerResponder[armapimanagement.DeletedServicesClientPurgeResponse]](),
	}
}

// DeletedServicesServerTransport connects instances of armapimanagement.DeletedServicesClient to instances of DeletedServicesServer.
// Don't use this type directly, use NewDeletedServicesServerTransport instead.
type DeletedServicesServerTransport struct {
	srv                        *DeletedServicesServer
	newListBySubscriptionPager *tracker[azfake.PagerResponder[armapimanagement.DeletedServicesClientListBySubscriptionResponse]]
	beginPurge                 *tracker[azfake.PollerResponder[armapimanagement.DeletedServicesClientPurgeResponse]]
}

// Do implements the policy.Transporter interface for DeletedServicesServerTransport.
func (d *DeletedServicesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DeletedServicesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if deletedServicesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = deletedServicesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "DeletedServicesClient.GetByName":
				res.resp, res.err = d.dispatchGetByName(req)
			case "DeletedServicesClient.NewListBySubscriptionPager":
				res.resp, res.err = d.dispatchNewListBySubscriptionPager(req)
			case "DeletedServicesClient.BeginPurge":
				res.resp, res.err = d.dispatchBeginPurge(req)
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

func (d *DeletedServicesServerTransport) dispatchGetByName(req *http.Request) (*http.Response, error) {
	if d.srv.GetByName == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByName not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deletedservices/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.GetByName(req.Context(), serviceNameParam, locationParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeletedServiceContract, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DeletedServicesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := d.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/deletedservices`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := d.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		d.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armapimanagement.DeletedServicesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		d.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (d *DeletedServicesServerTransport) dispatchBeginPurge(req *http.Request) (*http.Response, error) {
	if d.srv.BeginPurge == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPurge not implemented")}
	}
	beginPurge := d.beginPurge.get(req)
	if beginPurge == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deletedservices/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginPurge(req.Context(), serviceNameParam, locationParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPurge = &respr
		d.beginPurge.add(req, beginPurge)
	}

	resp, err := server.PollerResponderNext(beginPurge, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		d.beginPurge.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPurge) {
		d.beginPurge.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to DeletedServicesServerTransport
var deletedServicesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
