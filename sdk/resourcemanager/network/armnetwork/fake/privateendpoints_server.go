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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateEndpointsServer is a fake server for instances of the armnetwork.PrivateEndpointsClient type.
type PrivateEndpointsServer struct {
	// BeginCreateOrUpdate is the fake for method PrivateEndpointsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, privateEndpointName string, parameters armnetwork.PrivateEndpoint, options *armnetwork.PrivateEndpointsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.PrivateEndpointsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PrivateEndpointsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, privateEndpointName string, options *armnetwork.PrivateEndpointsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.PrivateEndpointsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PrivateEndpointsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, privateEndpointName string, options *armnetwork.PrivateEndpointsClientGetOptions) (resp azfake.Responder[armnetwork.PrivateEndpointsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PrivateEndpointsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armnetwork.PrivateEndpointsClientListOptions) (resp azfake.PagerResponder[armnetwork.PrivateEndpointsClientListResponse])

	// NewListBySubscriptionPager is the fake for method PrivateEndpointsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armnetwork.PrivateEndpointsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armnetwork.PrivateEndpointsClientListBySubscriptionResponse])
}

// NewPrivateEndpointsServerTransport creates a new instance of PrivateEndpointsServerTransport with the provided implementation.
// The returned PrivateEndpointsServerTransport instance is connected to an instance of armnetwork.PrivateEndpointsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateEndpointsServerTransport(srv *PrivateEndpointsServer) *PrivateEndpointsServerTransport {
	return &PrivateEndpointsServerTransport{
		srv:                        srv,
		beginCreateOrUpdate:        newTracker[azfake.PollerResponder[armnetwork.PrivateEndpointsClientCreateOrUpdateResponse]](),
		beginDelete:                newTracker[azfake.PollerResponder[armnetwork.PrivateEndpointsClientDeleteResponse]](),
		newListPager:               newTracker[azfake.PagerResponder[armnetwork.PrivateEndpointsClientListResponse]](),
		newListBySubscriptionPager: newTracker[azfake.PagerResponder[armnetwork.PrivateEndpointsClientListBySubscriptionResponse]](),
	}
}

// PrivateEndpointsServerTransport connects instances of armnetwork.PrivateEndpointsClient to instances of PrivateEndpointsServer.
// Don't use this type directly, use NewPrivateEndpointsServerTransport instead.
type PrivateEndpointsServerTransport struct {
	srv                        *PrivateEndpointsServer
	beginCreateOrUpdate        *tracker[azfake.PollerResponder[armnetwork.PrivateEndpointsClientCreateOrUpdateResponse]]
	beginDelete                *tracker[azfake.PollerResponder[armnetwork.PrivateEndpointsClientDeleteResponse]]
	newListPager               *tracker[azfake.PagerResponder[armnetwork.PrivateEndpointsClientListResponse]]
	newListBySubscriptionPager *tracker[azfake.PagerResponder[armnetwork.PrivateEndpointsClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for PrivateEndpointsServerTransport.
func (p *PrivateEndpointsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PrivateEndpointsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if privateEndpointsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = privateEndpointsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "PrivateEndpointsClient.BeginCreateOrUpdate":
				res.resp, res.err = p.dispatchBeginCreateOrUpdate(req)
			case "PrivateEndpointsClient.BeginDelete":
				res.resp, res.err = p.dispatchBeginDelete(req)
			case "PrivateEndpointsClient.Get":
				res.resp, res.err = p.dispatchGet(req)
			case "PrivateEndpointsClient.NewListPager":
				res.resp, res.err = p.dispatchNewListPager(req)
			case "PrivateEndpointsClient.NewListBySubscriptionPager":
				res.resp, res.err = p.dispatchNewListBySubscriptionPager(req)
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

func (p *PrivateEndpointsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := p.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/privateEndpoints/(?P<privateEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.PrivateEndpoint](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, privateEndpointNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		p.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		p.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		p.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (p *PrivateEndpointsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/privateEndpoints/(?P<privateEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, privateEndpointNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		p.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		p.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		p.beginDelete.remove(req)
	}

	return resp, nil
}

func (p *PrivateEndpointsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/privateEndpoints/(?P<privateEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	privateEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armnetwork.PrivateEndpointsClientGetOptions
	if expandParam != nil {
		options = &armnetwork.PrivateEndpointsClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, privateEndpointNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateEndpoint, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateEndpointsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/privateEndpoints`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListPager(resourceGroupNameParam, nil)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.PrivateEndpointsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		p.newListPager.remove(req)
	}
	return resp, nil
}

func (p *PrivateEndpointsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := p.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/privateEndpoints`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := p.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		p.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armnetwork.PrivateEndpointsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		p.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to PrivateEndpointsServerTransport
var privateEndpointsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
