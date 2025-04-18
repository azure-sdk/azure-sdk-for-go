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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/voiceservices/armvoiceservices"
	"net/http"
	"net/url"
	"regexp"
)

// CommunicationsGatewaysServer is a fake server for instances of the armvoiceservices.CommunicationsGatewaysClient type.
type CommunicationsGatewaysServer struct {
	// BeginCreateOrUpdate is the fake for method CommunicationsGatewaysClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, communicationsGatewayName string, resource armvoiceservices.CommunicationsGateway, options *armvoiceservices.CommunicationsGatewaysClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armvoiceservices.CommunicationsGatewaysClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method CommunicationsGatewaysClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, communicationsGatewayName string, options *armvoiceservices.CommunicationsGatewaysClientBeginDeleteOptions) (resp azfake.PollerResponder[armvoiceservices.CommunicationsGatewaysClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method CommunicationsGatewaysClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, communicationsGatewayName string, options *armvoiceservices.CommunicationsGatewaysClientGetOptions) (resp azfake.Responder[armvoiceservices.CommunicationsGatewaysClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method CommunicationsGatewaysClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armvoiceservices.CommunicationsGatewaysClientListByResourceGroupOptions) (resp azfake.PagerResponder[armvoiceservices.CommunicationsGatewaysClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method CommunicationsGatewaysClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armvoiceservices.CommunicationsGatewaysClientListBySubscriptionOptions) (resp azfake.PagerResponder[armvoiceservices.CommunicationsGatewaysClientListBySubscriptionResponse])

	// Update is the fake for method CommunicationsGatewaysClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, communicationsGatewayName string, properties armvoiceservices.CommunicationsGatewayUpdate, options *armvoiceservices.CommunicationsGatewaysClientUpdateOptions) (resp azfake.Responder[armvoiceservices.CommunicationsGatewaysClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewCommunicationsGatewaysServerTransport creates a new instance of CommunicationsGatewaysServerTransport with the provided implementation.
// The returned CommunicationsGatewaysServerTransport instance is connected to an instance of armvoiceservices.CommunicationsGatewaysClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCommunicationsGatewaysServerTransport(srv *CommunicationsGatewaysServer) *CommunicationsGatewaysServerTransport {
	return &CommunicationsGatewaysServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armvoiceservices.CommunicationsGatewaysClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armvoiceservices.CommunicationsGatewaysClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armvoiceservices.CommunicationsGatewaysClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armvoiceservices.CommunicationsGatewaysClientListBySubscriptionResponse]](),
	}
}

// CommunicationsGatewaysServerTransport connects instances of armvoiceservices.CommunicationsGatewaysClient to instances of CommunicationsGatewaysServer.
// Don't use this type directly, use NewCommunicationsGatewaysServerTransport instead.
type CommunicationsGatewaysServerTransport struct {
	srv                         *CommunicationsGatewaysServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armvoiceservices.CommunicationsGatewaysClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armvoiceservices.CommunicationsGatewaysClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armvoiceservices.CommunicationsGatewaysClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armvoiceservices.CommunicationsGatewaysClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for CommunicationsGatewaysServerTransport.
func (c *CommunicationsGatewaysServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *CommunicationsGatewaysServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if communicationsGatewaysServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = communicationsGatewaysServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "CommunicationsGatewaysClient.BeginCreateOrUpdate":
				res.resp, res.err = c.dispatchBeginCreateOrUpdate(req)
			case "CommunicationsGatewaysClient.BeginDelete":
				res.resp, res.err = c.dispatchBeginDelete(req)
			case "CommunicationsGatewaysClient.Get":
				res.resp, res.err = c.dispatchGet(req)
			case "CommunicationsGatewaysClient.NewListByResourceGroupPager":
				res.resp, res.err = c.dispatchNewListByResourceGroupPager(req)
			case "CommunicationsGatewaysClient.NewListBySubscriptionPager":
				res.resp, res.err = c.dispatchNewListBySubscriptionPager(req)
			case "CommunicationsGatewaysClient.Update":
				res.resp, res.err = c.dispatchUpdate(req)
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

func (c *CommunicationsGatewaysServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := c.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.VoiceServices/communicationsGateways/(?P<communicationsGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armvoiceservices.CommunicationsGateway](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		communicationsGatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("communicationsGatewayName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, communicationsGatewayNameParam, body, nil)
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

func (c *CommunicationsGatewaysServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.VoiceServices/communicationsGateways/(?P<communicationsGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		communicationsGatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("communicationsGatewayName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, communicationsGatewayNameParam, nil)
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

func (c *CommunicationsGatewaysServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.VoiceServices/communicationsGateways/(?P<communicationsGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	communicationsGatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("communicationsGatewayName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, communicationsGatewayNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CommunicationsGateway, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CommunicationsGatewaysServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := c.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.VoiceServices/communicationsGateways`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		c.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armvoiceservices.CommunicationsGatewaysClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		c.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (c *CommunicationsGatewaysServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := c.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.VoiceServices/communicationsGateways`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := c.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		c.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armvoiceservices.CommunicationsGatewaysClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		c.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (c *CommunicationsGatewaysServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.VoiceServices/communicationsGateways/(?P<communicationsGatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armvoiceservices.CommunicationsGatewayUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	communicationsGatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("communicationsGatewayName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Update(req.Context(), resourceGroupNameParam, communicationsGatewayNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CommunicationsGateway, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to CommunicationsGatewaysServerTransport
var communicationsGatewaysServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
