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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateLinkServicesForMIPPolicySyncServer is a fake server for instances of the armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClient type.
type PrivateLinkServicesForMIPPolicySyncServer struct {
	// BeginCreateOrUpdate is the fake for method PrivateLinkServicesForMIPPolicySyncClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, resourceName string, privateLinkServicesForMIPPolicySyncDescription armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncDescription, options *armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PrivateLinkServicesForMIPPolicySyncClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, resourceName string, options *armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientBeginDeleteOptions) (resp azfake.PollerResponder[armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PrivateLinkServicesForMIPPolicySyncClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, options *armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientGetOptions) (resp azfake.Responder[armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PrivateLinkServicesForMIPPolicySyncClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientListOptions) (resp azfake.PagerResponder[armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientListResponse])

	// NewListByResourceGroupPager is the fake for method PrivateLinkServicesForMIPPolicySyncClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientListByResourceGroupOptions) (resp azfake.PagerResponder[armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientListByResourceGroupResponse])

	// BeginUpdate is the fake for method PrivateLinkServicesForMIPPolicySyncClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, resourceName string, servicePatchDescription armm365securityandcompliance.ServicesPatchDescription, options *armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientBeginUpdateOptions) (resp azfake.PollerResponder[armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewPrivateLinkServicesForMIPPolicySyncServerTransport creates a new instance of PrivateLinkServicesForMIPPolicySyncServerTransport with the provided implementation.
// The returned PrivateLinkServicesForMIPPolicySyncServerTransport instance is connected to an instance of armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateLinkServicesForMIPPolicySyncServerTransport(srv *PrivateLinkServicesForMIPPolicySyncServer) *PrivateLinkServicesForMIPPolicySyncServerTransport {
	return &PrivateLinkServicesForMIPPolicySyncServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientDeleteResponse]](),
		newListPager:                newTracker[azfake.PagerResponder[armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientListByResourceGroupResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientUpdateResponse]](),
	}
}

// PrivateLinkServicesForMIPPolicySyncServerTransport connects instances of armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClient to instances of PrivateLinkServicesForMIPPolicySyncServer.
// Don't use this type directly, use NewPrivateLinkServicesForMIPPolicySyncServerTransport instead.
type PrivateLinkServicesForMIPPolicySyncServerTransport struct {
	srv                         *PrivateLinkServicesForMIPPolicySyncServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientDeleteResponse]]
	newListPager                *tracker[azfake.PagerResponder[armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientListByResourceGroupResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for PrivateLinkServicesForMIPPolicySyncServerTransport.
func (p *PrivateLinkServicesForMIPPolicySyncServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PrivateLinkServicesForMIPPolicySyncServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if privateLinkServicesForMipPolicySyncServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = privateLinkServicesForMipPolicySyncServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "PrivateLinkServicesForMIPPolicySyncClient.BeginCreateOrUpdate":
				res.resp, res.err = p.dispatchBeginCreateOrUpdate(req)
			case "PrivateLinkServicesForMIPPolicySyncClient.BeginDelete":
				res.resp, res.err = p.dispatchBeginDelete(req)
			case "PrivateLinkServicesForMIPPolicySyncClient.Get":
				res.resp, res.err = p.dispatchGet(req)
			case "PrivateLinkServicesForMIPPolicySyncClient.NewListPager":
				res.resp, res.err = p.dispatchNewListPager(req)
			case "PrivateLinkServicesForMIPPolicySyncClient.NewListByResourceGroupPager":
				res.resp, res.err = p.dispatchNewListByResourceGroupPager(req)
			case "PrivateLinkServicesForMIPPolicySyncClient.BeginUpdate":
				res.resp, res.err = p.dispatchBeginUpdate(req)
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

func (p *PrivateLinkServicesForMIPPolicySyncServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := p.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.M365SecurityAndCompliance/privateLinkServicesForMIPPolicySync/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncDescription](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, resourceNameParam, body, nil)
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

func (p *PrivateLinkServicesForMIPPolicySyncServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.M365SecurityAndCompliance/privateLinkServicesForMIPPolicySync/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
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

func (p *PrivateLinkServicesForMIPPolicySyncServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.M365SecurityAndCompliance/privateLinkServicesForMIPPolicySync/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateLinkServicesForMIPPolicySyncDescription, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateLinkServicesForMIPPolicySyncServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.M365SecurityAndCompliance/privateLinkServicesForMIPPolicySync`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := p.srv.NewListPager(nil)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientListResponse, createLink func() string) {
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

func (p *PrivateLinkServicesForMIPPolicySyncServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := p.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.M365SecurityAndCompliance/privateLinkServicesForMIPPolicySync`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		p.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armm365securityandcompliance.PrivateLinkServicesForMIPPolicySyncClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		p.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (p *PrivateLinkServicesForMIPPolicySyncServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := p.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.M365SecurityAndCompliance/privateLinkServicesForMIPPolicySync/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armm365securityandcompliance.ServicesPatchDescription](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginUpdate(req.Context(), resourceGroupNameParam, resourceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		p.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		p.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		p.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to PrivateLinkServicesForMIPPolicySyncServerTransport
var privateLinkServicesForMipPolicySyncServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
