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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
	"net/http"
	"net/url"
	"regexp"
)

// PlacementPoliciesServer is a fake server for instances of the armavs.PlacementPoliciesClient type.
type PlacementPoliciesServer struct {
	// BeginCreateOrUpdate is the fake for method PlacementPoliciesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, placementPolicy armavs.PlacementPolicy, options *armavs.PlacementPoliciesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armavs.PlacementPoliciesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PlacementPoliciesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, options *armavs.PlacementPoliciesClientBeginDeleteOptions) (resp azfake.PollerResponder[armavs.PlacementPoliciesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PlacementPoliciesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, options *armavs.PlacementPoliciesClientGetOptions) (resp azfake.Responder[armavs.PlacementPoliciesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PlacementPoliciesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, privateCloudName string, clusterName string, options *armavs.PlacementPoliciesClientListOptions) (resp azfake.PagerResponder[armavs.PlacementPoliciesClientListResponse])

	// BeginUpdate is the fake for method PlacementPoliciesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string, placementPolicyUpdate armavs.PlacementPolicyUpdate, options *armavs.PlacementPoliciesClientBeginUpdateOptions) (resp azfake.PollerResponder[armavs.PlacementPoliciesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewPlacementPoliciesServerTransport creates a new instance of PlacementPoliciesServerTransport with the provided implementation.
// The returned PlacementPoliciesServerTransport instance is connected to an instance of armavs.PlacementPoliciesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPlacementPoliciesServerTransport(srv *PlacementPoliciesServer) *PlacementPoliciesServerTransport {
	return &PlacementPoliciesServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armavs.PlacementPoliciesClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armavs.PlacementPoliciesClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armavs.PlacementPoliciesClientListResponse]](),
		beginUpdate:         newTracker[azfake.PollerResponder[armavs.PlacementPoliciesClientUpdateResponse]](),
	}
}

// PlacementPoliciesServerTransport connects instances of armavs.PlacementPoliciesClient to instances of PlacementPoliciesServer.
// Don't use this type directly, use NewPlacementPoliciesServerTransport instead.
type PlacementPoliciesServerTransport struct {
	srv                 *PlacementPoliciesServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armavs.PlacementPoliciesClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armavs.PlacementPoliciesClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armavs.PlacementPoliciesClientListResponse]]
	beginUpdate         *tracker[azfake.PollerResponder[armavs.PlacementPoliciesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for PlacementPoliciesServerTransport.
func (p *PlacementPoliciesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PlacementPoliciesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if placementPoliciesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = placementPoliciesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "PlacementPoliciesClient.BeginCreateOrUpdate":
				res.resp, res.err = p.dispatchBeginCreateOrUpdate(req)
			case "PlacementPoliciesClient.BeginDelete":
				res.resp, res.err = p.dispatchBeginDelete(req)
			case "PlacementPoliciesClient.Get":
				res.resp, res.err = p.dispatchGet(req)
			case "PlacementPoliciesClient.NewListPager":
				res.resp, res.err = p.dispatchNewListPager(req)
			case "PlacementPoliciesClient.BeginUpdate":
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

func (p *PlacementPoliciesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := p.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/placementPolicies/(?P<placementPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armavs.PlacementPolicy](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		placementPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("placementPolicyName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, privateCloudNameParam, clusterNameParam, placementPolicyNameParam, body, nil)
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

func (p *PlacementPoliciesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/placementPolicies/(?P<placementPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		placementPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("placementPolicyName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, privateCloudNameParam, clusterNameParam, placementPolicyNameParam, nil)
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

func (p *PlacementPoliciesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/placementPolicies/(?P<placementPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
	if err != nil {
		return nil, err
	}
	clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
	if err != nil {
		return nil, err
	}
	placementPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("placementPolicyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, privateCloudNameParam, clusterNameParam, placementPolicyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PlacementPolicy, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PlacementPoliciesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/placementPolicies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListPager(resourceGroupNameParam, privateCloudNameParam, clusterNameParam, nil)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armavs.PlacementPoliciesClientListResponse, createLink func() string) {
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

func (p *PlacementPoliciesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := p.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/placementPolicies/(?P<placementPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armavs.PlacementPolicyUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		placementPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("placementPolicyName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginUpdate(req.Context(), resourceGroupNameParam, privateCloudNameParam, clusterNameParam, placementPolicyNameParam, body, nil)
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

// set this to conditionally intercept incoming requests to PlacementPoliciesServerTransport
var placementPoliciesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
